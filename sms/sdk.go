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
	"encoding/json"
	"github.com/yanue/go-esport-common/validator"
	"math/rand"
)

const letterBytes = "0123456789"

type smsSdk struct {
	// AliSmsSdk 阿里云短信
	*AliSmsSdk
	YunpianApiKey string
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
		return this.SendAliSms(phone, smsCommon_CN, map[string]interface{}{"code": code})
	}

	// 国际云片短信发送
	tmp := YunpianSmsCommon_EN
	switch area {
	case validator.AreaCode_TW:
		tmp = YunpianSmsCommon_TW
	}

	areaStr := string(area)

	return this.SendYunpianSms(areaStr+phone, tmp, map[string]interface{}{"code": code})
}

/*
 *@note 通过阿里大于发送国内短信通知
 *@param phone 手机号
 *@param templateCode 短信模板
 *@param smsParam 短信模板变量参数
 *@return 错误信息
 */
func (this *smsSdk) SendAliSms(phone, templateCode string, smsParam map[string]interface{}) (errCode int32) {
	// phoneNumbers,  templateCode, templateParam
	templateParam := ""

	if smsParam != nil {
		bJson, _ := json.Marshal(&smsParam)
		templateParam = string(bJson)
	}

	// request
	_, errCode = this.sendAliSms(phone, templateCode, templateParam)

	return
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
