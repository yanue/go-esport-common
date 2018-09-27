/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : ali.go
 Time    : 2018/9/26 15:05
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package sms

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"github.com/pborman/uuid"
	"github.com/yanue/go-esport-common"
	"github.com/yanue/go-esport-common/errcode"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const aliSmsGatewayUrl = "https://dysmsapi.aliyuncs.com"

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h897")

/*
失败:
{"Message":"签名不合法(不存在或被拉黑)","RequestId":"3C506816-9F62-4F59-9A4B-1DFB0CE89219","Code":"isv.SMS_SIGNATURE_ILLEGAL"}
成功:
{"Message":"OK","RequestId":"13224446-52BA-4B72-8A94-BE73D1547693","BizId":"857623737957293027^0","Code":"OK"}
*/
type aliSmsResponse struct {
	BizId     string `json:"BizId"`
	Code      string `json:"Code"`
	Message   string `json:"Message"`
	RequestId string `json:"RequestId"`
}

// 阿里云短信结构
type AliSmsSdk struct {
	//system parameters
	accessKeyId      string
	accessKeySecret  string
	timestamp        string
	format           string
	signatureMethod  string
	signatureVersion string
	signatureNonce   string
	signature        string

	//business parameters
	action        string
	version       string
	regionId      string
	phoneNumbers  string
	signName      string
	templateCode  string
	templateParam string
	outId         string
}

// 发送验证码
func (this *AliSmsSdk) sendAliSms(phoneNumbers, templateCode, templateParam string) ([]byte, int32) {
	// 初始化基础参数
	this.init()

	// 其他输入参数
	this.phoneNumbers = phoneNumbers
	this.templateCode = templateCode
	this.templateParam = templateParam

	// 构建带签名url
	signUrl, errno := this.buildSignUrl()
	if errno > 0 {
		return nil, errno
	}

	resp, err := http.Get(signUrl)
	if err != nil {
		return nil, errcode.ErrCommonRemotecall
	}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errcode.ErrCommonRemotecall
	}

	var respSms = new(aliSmsResponse)
	json.Unmarshal(respData, respSms)

	// code is "OK"
	if respSms.Code == "OK" {
		return respData, errcode.No_Error
	}

	return nil, this.parseErr(respSms.Code)
}

// 构建url
func (this *AliSmsSdk) buildSignUrl() (string, int32) {
	if len(this.accessKeyId) == 0 {
		return "", errcode.ErrSmsInvalidAccesskeyid
	}

	if len(this.phoneNumbers) == 0 {
		return "", errcode.ErrSmsInvalidAccesskeysecret
	}

	if len(this.signName) == 0 {
		return "", errcode.ErrSmsInvalidSignname
	}

	if len(this.templateCode) == 0 {
		return "", errcode.ErrSmsInvalidTemplatecode
	}

	if len(this.templateParam) == 0 {
		return "", errcode.ErrSmsInvalidTemplateparam
	}

	// common params
	systemParams := make(map[string]string)
	systemParams["SignatureMethod"] = this.signatureMethod
	systemParams["SignatureNonce"] = this.signatureNonce
	systemParams["AccessKeyId"] = this.accessKeyId
	systemParams["SignatureVersion"] = this.signatureVersion
	systemParams["Timestamp"] = this.timestamp
	systemParams["Format"] = this.format

	// business params
	businessParams := make(map[string]string)
	businessParams["Action"] = this.action
	businessParams["Version"] = this.version
	businessParams["RegionId"] = this.regionId
	businessParams["PhoneNumbers"] = this.phoneNumbers
	businessParams["SignName"] = this.signName
	businessParams["TemplateParam"] = this.templateParam
	businessParams["TemplateCode"] = this.templateCode
	businessParams["OutId"] = this.outId

	// generate signature and sorted query
	sortedQueryString, signature := this.generateQueryStringAndSignature(businessParams, systemParams)

	return aliSmsGatewayUrl + "?Signature=" + signature + "&" + sortedQueryString, 0
}

// 生成参数及签名
func (this *AliSmsSdk) generateQueryStringAndSignature(businessParams map[string]string, systemParams map[string]string) (string, string) {
	keys := make([]string, 0)
	allParams := make(map[string]string)

	for key, value := range businessParams {
		keys = append(keys, key)
		allParams[key] = value
	}

	for key, value := range systemParams {
		keys = append(keys, key)
		allParams[key] = value
	}

	// key排序
	sort.Strings(keys)

	sortedQueryString := ""
	// 处理排序参数
	for _, key := range keys {
		if key == "Signature" {
			continue
		}
		encKey := this.specialUrlEncode(key)
		encVal := this.specialUrlEncode(allParams[key])
		if len(sortedQueryString) > 0 {
			sortedQueryString += "&"
		}
		sortedQueryString += encKey + "=" + encVal
	}

	// 签名参数
	stringToSign := "GET" + "&" + this.specialUrlEncode("/") + "&" + this.specialUrlEncode(sortedQueryString)
	// 签名
	signature := this.sign(this.accessKeySecret+"&", stringToSign)
	// 签名url处理
	signature = this.specialUrlEncode(signature)

	return sortedQueryString, signature
}

// 处理url参数
func (this *AliSmsSdk) specialUrlEncode(value string) string {
	rstValue := url.QueryEscape(value)
	rstValue = strings.Replace(rstValue, "+", "%20", -1)
	rstValue = strings.Replace(rstValue, "*", "%2A", -1)
	rstValue = strings.Replace(rstValue, "%7E", "~", -1)
	return rstValue
}

// 计算 HMAC 值。
//     按照 RFC2104 的定义，使用得到的签名字符串计算签名 HMAC 值。
//     注意：计算签名时使用的 Key 就是您持有的 Access Key Secret 并加上一个 “&” 字符（ASCII:38），使用的哈希算法是 SHA1。
// 计算签名值。
//     按照 Base64 编码规则 把步骤 3 中的 HMAC 值编码成字符串，即得到签名值（Signature）。
func (this *AliSmsSdk) sign(key, stringToSign string) string {
	// The signature method is supposed to be HmacSHA1
	// A switch case is required if there is other methods available
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// 签名nonce
func (this *AliSmsSdk) newSignatureNonce() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(26)
	//return b.String()
	return uuid.New()
}

/**
 *@note 阿里短信错误处理
 */
func (this *AliSmsSdk) parseErr(subCode string) (errCode int32) {
	switch subCode {
	case "isv.RAM_PERMISSION_DENY":
		return errcode.ErrsmsPermissionDeny
	case "isv.OUT_OF_SERVICE":
		return errcode.ErrSmsOutOfService
	case "isv.PRODUCT_UN_SUBSCRIPT":
		return errcode.ErrSmsOutOfService
	case "isv.PRODUCT_UNSUBSCRIBE":
		return errcode.ErrSmsProductUnsubscribe
	case "isv.ACCOUNT_NOT_EXISTS":
		return errcode.ErrSmsAccountNotExists
	case "isv.ACCOUNT_ABNORMAL":
		return errcode.ErrSmsAccountAbnormal
	case "isv.TEMPLATE_PARAMS_ILLEGAL":
		return errcode.ErrSmsSmsTemplateIllegal
	case "isv.SMS_SIGNATURE_ILLEGAL":
		return errcode.ErrSmsSmsSignatureIllegal
	case "isv.INVALID_PARAMETERS":
		return errcode.ErrSmsInvalidParameters
	case "isv.SYSTEM_ERROR":
		return errcode.ErrSmsSystemError
	case "isv.MOBILE_NUMBER_ILLEGAL":
		return errcode.ErrSmsMobileNumberIllegal
	case "isv.MOBILE_COUNT_OVER_LIMIT":
		return errcode.ErrSmsMobileCountOverLimit
	case "isv.TEMPLATE_MISSING_PARAMETERS":
		return errcode.ErrSmsTemplateMissingParameters
	case "isv.BUSINESS_LIMIT_CONTROL":
		return errcode.ErrSmsBusinessLimitControl
	case "isv.INVALID_JSON_PARAM":
		return errcode.ErrSmsInvalidJsonParam
	case "isv.BLACK_KEY_CONTROL_LIMIT":
		return errcode.ErrSmsBlackKeyControlLimit
	case "isv.PARAM_LENGTH_LIMIT":
		return errcode.ErrSmsParamLengthLimit
	case "isv.PARAM_NOT_SUPPORT_URL":
		return errcode.ErrSmsParamNotSupportUrl
	case "isv.AMOUNT_NOT_ENOUGH":
		return errcode.ErrSmsAmountNotEnough
	case "SignatureDoesNotMatch":
		return errcode.ErrSmsInvalidSignature
	case "SignatureNonceUsed":
		return errcode.ErrSmsSignatureNonceUsed
	}

	// 记录日志
	common.Logs.Info("subCode:", subCode)

	return errcode.ErrSmsUnknownResponse
}

// 初始化默认参数
func (this *AliSmsSdk) init() {
	local, _ := time.LoadLocation("GMT")
	this.timestamp = time.Now().In(local).Format("2006-01-02T15:04:05Z")
	this.format = "json"
	this.signatureMethod = "HMAC-SHA1"
	this.signatureVersion = "1.0"
	this.signatureNonce = this.newSignatureNonce()
	this.action = "SendSms"
	this.version = "2017-05-25"
	this.regionId = "cn-hangzhou"
	this.outId = "abcdefg"
}
