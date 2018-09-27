/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : sdk_yunpian.go
 Time    : 2018/9/26 15:51
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package sms

import (
	"encoding/json"
	"github.com/yanue/go-esport-common"
	"github.com/yanue/go-esport-common/errcode"
	"io/ioutil"
	"net/http"
	"net/url"
)

const yunpianApiUrl = "https://sms.yunpian.com/v1/sms/tpl_send.json"

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
 *@note 增加云片短信支持
 *@param apiKey 云片apikey
 */
func (this *smsSdk) AddYunpianApiKey(apiKey string) {
	this.YunpianApiKey = apiKey
}

/*
 *@note 通过云片发送国际短信通知
 *@param phone 手机号
 *@param templateCode 短信模板ID
 *@param smsParam 短信模板变量参数
 *@return 错误信息
 */
func (this *smsSdk) SendYunpianSms(phone, templateCode string, smsParam map[string]interface{}) (errCode int32) {
	tplValue := url.Values{}
	for k, v := range smsParam {
		tplValue["#"+k+"#"] = []string{v.(string)}
	}

	dataTplSms := url.Values{"apikey": {this.YunpianApiKey}, "phone": {phone}, "tpl_id": {templateCode}, "tpl_value": {tplValue.Encode()}}

	common.Logs.Debug("dataTplSms = %v", dataTplSms)

	// 发送请求
	_, errCode = this.doHttpsPostYunpian(yunpianApiUrl, dataTplSms)

	return
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
	//common.Logs.Warn(fmt.Sprintf("Error code: %d", code))
	switch code {
	case 1, 2: // 请求参数缺失,请求参数格式错误
		return errcode.ErrSmsInvalidParameters
	case 3: // 账户余额不足
		return errcode.ErrSmsAmountNotEnough
	case 5, 6, 7: // 未找到对应id的模板,模板不可用
		return errcode.ErrSmsSmsTemplateIllegal
	case 8, 9, 33: // 手机号频率受限
		return errcode.ErrSmsMobileCountOverLimit
	}
	return errcode.ErrSmsSysBusy
}
