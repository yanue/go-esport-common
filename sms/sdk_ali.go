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
	"fmt"
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
{
	Message: "签名不合法(不存在或被拉黑)",
	RequestId: "3C506816-9F62-4F59-9A4B-1DFB0CE89219",
	Code: "isv.SMS_SIGNATURE_ILLEGAL"
}
*/
type aliErrorResponse struct {
	Code      string `json:"Code"`
	Message   string `json:"Message"`
	RequestId string `json:"RequestId"`
}

// 阿里云短信结构
type AliSmsSdk struct {
	//system parameters
	AccessKeyId      string
	AccessKeySecret  string
	Timestamp        string
	Format           string
	SignatureMethod  string
	SignatureVersion string
	SignatureNonce   string
	Signature        string

	//business parameters
	Action        string
	Version       string
	RegionId      string
	PhoneNumbers  string
	SignName      string
	TemplateCode  string
	TemplateParam string
	OutId         string
}

// 发送验证码
func (this *AliSmsSdk) sendAliSms(phoneNumbers, templateCode, templateParam string) ([]byte, int32) {
	this.init()

	this.PhoneNumbers = phoneNumbers
	this.TemplateCode = templateCode
	this.TemplateParam = templateParam

	signUrl, errno := this.buildSignUrl()
	if errno > 0 {
		return nil, errno
	}
	common.Logs.Info("signUrl:", signUrl)
	//return nil, 0

	resp, err := http.Get(signUrl)
	if err != nil {
		return nil, errcode.ErrCommonRemotecall
	}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errcode.ErrCommonRemotecall
	}

	common.Logs.Info("respData:", string(respData))

	if strings.Contains(string(respData), "code") {
		var e = new(aliErrorResponse)
		json.Unmarshal(respData, e)
		return nil, this.parseErr(e.Code)
	}

	return respData, errcode.No_Error
}

// 构建url
func (this *AliSmsSdk) buildSignUrl() (string, int32) {
	if len(this.AccessKeyId) == 0 {
		return "", errcode.ErrSmsInvalidAccesskeyid
	}

	if len(this.PhoneNumbers) == 0 {
		return "", errcode.ErrSmsInvalidAccesskeysecret
	}

	if len(this.SignName) == 0 {
		return "", errcode.ErrSmsInvalidSignname
	}

	if len(this.TemplateCode) == 0 {
		return "", errcode.ErrSmsInvalidTemplatecode
	}

	if len(this.TemplateParam) == 0 {
		return "", errcode.ErrSmsInvalidTemplateparam
	}

	// common params
	systemParams := make(map[string]string)
	systemParams["SignatureMethod"] = this.SignatureMethod
	systemParams["SignatureNonce"] = this.SignatureNonce
	systemParams["AccessKeyId"] = this.AccessKeyId
	systemParams["SignatureVersion"] = this.SignatureVersion
	systemParams["Timestamp"] = this.Timestamp
	systemParams["Format"] = this.Format

	// business params
	businessParams := make(map[string]string)
	businessParams["Action"] = this.Action
	businessParams["Version"] = this.Version
	businessParams["RegionId"] = this.RegionId
	businessParams["PhoneNumbers"] = this.PhoneNumbers
	businessParams["SignName"] = this.SignName
	businessParams["TemplateParam"] = this.TemplateParam
	businessParams["TemplateCode"] = this.TemplateCode
	businessParams["OutId"] = this.OutId

	// generate signature and sorted query
	sortedQueryString, signature := this.generateQueryStringAndSignature(businessParams, systemParams)

	fmt.Println("Signature:", signature)
	fmt.Println("sortedQueryString:", sortedQueryString)

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
	signature := this.sign(this.AccessKeySecret+"&", stringToSign)
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
	case "isv.out-of-service":
		return errcode.ErrSmsOutOfService
	case "isv.product-unsubscribe":
		return errcode.ErrSmsProductUnsubscribe
	case "isv.account-not-exists":
		return errcode.ErrSmsAccountNotExists
	case "isv.account-abnormal":
		return errcode.ErrSmsAccountAbnormal
	case "isv.sms-template-illegal":
		return errcode.ErrSmsSmsTemplateIllegal
	case "isv.sms-signature-illegal":
		return errcode.ErrSmsSmsSignatureIllegal
	case "isv.phone-number-illegal":
		return errcode.ErrSmsMobileNumberIllegal
	case "isv.phone-count-over-limit":
		return errcode.ErrSmsMobileCountOverLimit
	case "isv.template-missing-parameters":
		return errcode.ErrSmsTemplateMissingParameters
	case "isv.invalid-parameters":
		return errcode.ErrSmsInvalidParameters
	case "isv.business-limit-control":
		return errcode.ErrSmsBusinessLimitControl
	case "isv.invalid-json-param":
		return errcode.ErrSmsInvalidJsonParam
	case "isv.system-error":
		return errcode.ErrSmsSystemError
	case "isv.black-key-control-limit":
		return errcode.ErrSmsBlackKeyControlLimit
	case "isv.param-not-support-url":
		return errcode.ErrSmsParamNotSupportUrl
	case "isv.param-length-limit":
		return errcode.ErrSmsParamLengthLimit
	case "isv.amount-not-enough":
		return errcode.ErrSmsAmountNotEnough
	case "isv.appkey-not-exists":
		return errcode.ErrSmsInvalidAccesskeyid
	case "SignatureDoesNotMatch":
		return errcode.ErrSmsInvalidSignature
	}

	// 记录日志
	common.Logs.Info("subCode:", subCode)

	return errcode.ErrSmsUnknownResponse
}

// 初始化默认参数
func (this *AliSmsSdk) init() {
	local, _ := time.LoadLocation("GMT")
	this.Timestamp = time.Now().In(local).Format("2006-01-02T15:04:05Z")
	this.Format = "json"
	this.SignatureMethod = "HMAC-SHA1"
	this.SignatureVersion = "1.0"
	this.SignatureNonce = this.newSignatureNonce()
	this.Action = "SendSms"
	this.Version = "2017-05-25"
	this.RegionId = "cn-hangzhou"
	this.OutId = "abcdefg"
}
