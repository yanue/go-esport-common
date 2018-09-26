/*go**************************************************************************
 File            : Sms.go
 Subsystem       :
 Author          : mingang.he
 Date&Time       : 2016-08-03
 Description     : 短信服务
 Revision        :

 History
 -------

 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/
package sms

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/yanue/go-esport-common"
	"github.com/yanue/go-esport-common/errcode"
	"github.com/yanue/go-esport-common/validator"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	SENDSMSAPI    = "alibaba.aliqin.fc.sms.num.send"
	APIURL        = "http://gw.api.taobao.com/router/rest"
	YUNPIANAPIURL = "https://sms.yunpian.com/v1/sms/tpl_send.json"
)

type SmsLanguage int

// 语言类型
const (
	SmsLanguage_CN SmsLanguage = iota // 简体中文
	SmsLanguage_TW                    // 繁体中文
	SmsLanguage_EN                    // 英文
	SmsLanguage_KR                    // 韩文
)

const letterBytes = "0123456789"

/*
	{
		"code":0,
		"msg":"OK",
		"result":{
			"count":1,
			"sid":10846370200
		}
	}
*/
type resultErrorYunpian struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		Count int `json:"count"`
		Sid   int `json:"sid"`
	} `json:"result"`
}

/*
   {
	   "error_response":{
		   "code":50,
		   "msg":"Remote service error",
		   "sub_code":"isv.invalid-parameter",
		   "sub_msg":"非法参数"
	   }
   }
*/
type resultError struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
}

type smsSdk struct {
	AppKey          string
	AppSecret       string
	SmsFreeSignName string
	YunpianApiKey   string // 云片apikey
}

/**
 *@note 通用发送短信
 *@param phone 手机号
 *@param code 验证码
 *@param area 国家码
 *@param lang 语言
 *@return 错误信息
 */
func (this *smsSdk) SendCommonCode(phone, code string, area AreaCode, lang SmsLanguage) int32 {
	// 国内阿里短信
	if area == AreaCode_CN {
		return this.SendMessage(phone, smsCommon_CN, map[string]interface{}{"name": code})
	}

	// 国际云片短信发送
	tmp := YunpianSmsCommon_EN
	switch area {
	case validator.AreaCode_TW:
		tmp = YunpianSmsCommon_TW
	}

	areaStr := string(area)

	return this.SendOutboundMessage(areaStr+phone, tmp, map[string]interface{}{"code": code})
}

/*
 *@note 通过阿里大于发送国内短信通知
 *@param phone 手机号
 *@param tID 短信模板ID
 *@param smsParam 短信模板变量参数
 *@return 错误信息
 */
func (this *smsSdk) SendMessage(phone, tID string, smsParam map[string]interface{}) (errCode int32) {
	var req = map[string]string{
		"method":             SENDSMSAPI,
		"app_key":            this.AppKey,
		"v":                  "2.0",
		"timestamp":          time.Now().Format("2006-01-02 15:04:05"),
		"format":             "json",
		"sign_method":        "md5",
		"rec_num":            phone,
		"sms_type":           "normal",
		"sms_template_code":  tID,
		"sms_free_sign_name": this.SmsFreeSignName,
	}

	if smsParam != nil {
		bJson, _ := json.Marshal(&smsParam)
		req["sms_param"] = string(bJson)
	}

	// sign
	sign := fmt.Sprintf("%s%s%s", this.AppSecret, this.signParams(req, false), this.AppSecret)
	req["sign"] = fmt.Sprintf("%X", md5.Sum([]byte(sign)))

	data := url.Values{}
	for key, value := range req {
		data.Set(key, value)
	}

	// request
	_, errCode = this.doHttpPost(APIURL, data)

	return
}

/*
 *@note 通过云片发送国际短信通知
 *@param phone 手机号
 *@param tID 短信模板ID
 *@param smsParam 短信模板变量参数
 *@return 错误信息
 */
func (this *smsSdk) SendOutboundMessage(phone, tID string, smsParam map[string]interface{}) (errCode int32) {
	tplValue := url.Values{}
	for k, v := range smsParam {
		tplValue["#"+k+"#"] = []string{v.(string)}
	}

	dataTplSms := url.Values{"apikey": {this.YunpianApiKey}, "phone": {phone}, "tpl_id": {tID}, "tpl_value": {tplValue.Encode()}}

	common.Logs.Debug("dataTplSms = %v", dataTplSms)

	// 发送请求
	_, errCode = this.doHttpsPostYunpian(YUNPIANAPIURL, dataTplSms)

	return
}

/**
 *@note 请求发送阿里短信通知
 */
func (this *smsSdk) doHttpPost(targetUrl string, data url.Values) ([]byte, int32) {
	resp, err := http.PostForm(targetUrl, data)
	if err != nil {
		return nil, errcode.ErrCommonRemotecall
	}

	//req, _ := http.NewRequest("POST", targetUrl, bytes.NewBuffer(body))
	//req.Header.Add("Content-type", "application/x-www-form-urlencoded;charset=UTF-8")
	//
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	//}
	//client := &http.Client{Transport: tr}
	//
	//resp, err := client.PostForm(req)
	//if err != nil {
	//	return nil, errcode.ErrCommonRemotecall
	//}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errcode.ErrCommonRemotecall
	}

	if strings.Contains(string(respData), "error_response") {
		var e resultError
		json.Unmarshal(respData, &e)
		return nil, this.parseErr(e.ErrorResponse.SubMsg)
	}

	return respData, errcode.No_Error
}

/**
 *@note 请求发送云片短信通知
 */
func (this *smsSdk) doHttpsPostYunpian(targetUrl string, data url.Values) ([]byte, int32) {
	resp, err := http.PostForm(targetUrl, data)
	if err != nil {
		return nil, errcode.ErrCommonRemotecall
	}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errcode.ErrCommonRemotecall
	}

	var e resultErrorYunpian
	json.Unmarshal(respData, &e)
	return respData, this.parseYunpianErr(e.Code)
}

/**
 *@note 云片短信错误处理
 */
func (this *smsSdk) parseYunpianErr(code int) (errCode int32) {
	if code == 0 {
		return errcode.No_Error
	}
	common.Logs.Warn(fmt.Sprintf("Error code: %d", code))
	switch code {
	case 1, 2: // 请求参数缺失,请求参数格式错误
		return errcode.Err_Sms_INVALID_PARAMETERS
	case 3: // 账户余额不足
		return errcode.Err_Sms_AMOUNT_NOT_ENOUGH
	case 5, 6, 7: // 未找到对应id的模板,模板不可用
		return errcode.Err_Sms_SMS_TEMPLATE_ILLEGAL
	case 8, 9, 33: // 手机号频率受限
		return errcode.Err_Sms_MOBILE_COUNT_OVER_LIMIT
	}
	return errcode.Err_Sms_SYS_BUSY
}

/**
 *@note 阿里短信错误处理
 */
func (this *smsSdk) parseErr(subCode string) (errCode int32) {
	switch subCode {
	case "isv.out-of-service":
		return errcode.Err_Sms_OUT_OF_SERVICE
	case "isv.product-unsubscribe":
		return errcode.Err_sms_PRODUCT_UNSUBSCRIBE
	case "isv.account-not-exists":
		return errcode.Err_Sms_ACCOUNT_NOT_EXISTS
	case "isv.account-abnormal":
		return errcode.Err_Sms_ACCOUNT_ABNORMAL
	case "isv.sms-template-illegal":
		return errcode.Err_Sms_SMS_TEMPLATE_ILLEGAL
	case "isv.sms-signature-illegal":
		return errcode.Err_Sms_SMS_SIGNATURE_ILLEGAL
	case "isv.phone-number-illegal":
		return errcode.Err_Sms_MOBILE_NUMBER_ILLEGAL
	case "isv.phone-count-over-limit":
		return errcode.Err_Sms_MOBILE_COUNT_OVER_LIMIT
	case "isv.template-missing-parameters":
		return errcode.Err_Sms_TEMPLATE_MISSING_PARAMETERS
	case "isv.invalid-parameters":
		return errcode.Err_Sms_INVALID_PARAMETERS
	case "isv.business-limit-control":
		return errcode.Err_Sms_BUSINESS_LIMIT_CONTROL
	case "isv.invalid-json-param":
		return errcode.Err_Sms_INVALID_JSON_PARAM
	case "isv.system-error":
		return errcode.Err_Sms_SYSTEM_ERROR
	case "isv.black-key-control-limit":
		return errcode.Err_Sms_BLACK_KEY_CONTROL_LIMIT
	case "isv.param-not-support-url":
		return errcode.Err_Sms_PARAM_NOT_SUPPORT_URL
	case "isv.param-length-limit":
		return errcode.Err_Sms_PARAM_LENGTH_LIMIT
	case "isv.amount-not-enough":
		return errcode.Err_Sms_AMOUNT_NOT_ENOUGH
	}

	// 记录日志
	common.Logs.Info(subCode)

	return errcode.Err_Sms_SYS_BUSY
}

/**
@note 签名参数处理
 */
func (this smsSdk) signParams(params map[string]string, escape bool) string {
	if params == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := params[k]

		if escape {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(vs))
		} else {
			buf.WriteString(k)
			buf.WriteString(vs)
		}
	}
	return buf.Strng()
}

/**
@note 生成随机验证码
 */
func (this smsSdk) randNumber(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
