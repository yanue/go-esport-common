/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : sms.go
 Time    : 2018/9/25 14:57
 Author  : yanue
 
 - 验证码
 
------------------------------- go ---------------------------------*/

package errcode

var errMsgSms = map[ErrNo]string{
	// Sms
	ErrSmsVerifyCodeCheck:           "短信验证码校验失败",
	ErrSmsOutOfService:              "业务停机",
	ErrsmsPermissionDeny:            "权限不足",
	ErrSmsProductUnsubscribe:        "产品服务未开通",
	ErrSmsAccountNotExists:          "账户信息不存在",
	ErrSmsInvalidAccesskeyid:        "无效的AccessKeyId",
	ErrSmsInvalidAccesskeysecret:    "无效的AccessKeySecret",
	ErrSmsInvalidSignname:           "无效的SignName",
	ErrSmsInvalidTemplatecode:       "无效的TemplateCode",
	ErrSmsInvalidTemplateparam:      "无效的TemplateParam",
	ErrSmsInvalidSignature:          "签名认证无效",
	ErrSmsInvalidParameters:         "参数异常",
	ErrSmsInvalidJsonParam:          "JSON参数不合法",
	ErrSmsAccountAbnormal:           "账户信息异常",
	ErrSmsSmsTemplateIllegal:        "模板不合法",
	ErrSmsSmsSignatureIllegal:       "签名不合法",
	ErrSmsMobileNumberIllegal:       "手机号码格式错误",
	ErrSmsMobileCountOverLimit:      "手机号码数量超过限制",
	ErrSmsTemplateMissingParameters: "短信模板变量缺少参数",
	ErrSmsBusinessLimitControl:      "触发业务流控限制",
	ErrSmsSystemError:               "系统错误",
	ErrSmsSignatureNonceUsed:        "唯一随机数重复",
	ErrSmsUnknownResponse:           "未知的响应错误",
	ErrSmsBlackKeyControlLimit:      "模板变量中存在黑名单关键字",
	ErrSmsParamNotSupportUrl:        "不支持url为变量",
	ErrSmsParamLengthLimit:          "变量长度受限",
	ErrSmsAmountNotEnough:           "余额不足",
	ErrSmsSysBusy:                   "验证码获取过于频繁，请稍后再试",
	ErrSmsLimitMinute:               "验证码获取过于频繁，请1分钟后重试",
	ErrSmsLimitHour:                 "验证码获取过于频繁，请稍后再试",
}

const (
	// Sms
	ErrSmsVerifyCodeCheck           = errOffsetSms + iota // 短信验证码校验失败
	ErrSmsOutOfService                                    //业务停机",
	ErrsmsPermissionDeny                                  //产品服务未开通",
	ErrSmsProductUnsubscribe                              //产品服务未开通",
	ErrSmsAccountNotExists                                //账户信息不存在",
	ErrSmsInvalidAccesskeyid                              //无效的appkey",
	ErrSmsInvalidAccesskeysecret                          //无效的AccessKeyId",
	ErrSmsInvalidSignname                                 //无效的AccessKeySecret",
	ErrSmsInvalidTemplatecode                             //无效的TemplateCode",
	ErrSmsInvalidTemplateparam                            //无效的TemplateParam",
	ErrSmsInvalidSignature                                //无效的SignName",
	ErrSmsInvalidParameters                               //参数异常",
	ErrSmsInvalidJsonParam                                //JSON参数不合法",
	ErrSmsAccountAbnormal                                 //账号无效
	ErrSmsSmsTemplateIllegal                              //模板不合法",
	ErrSmsSmsSignatureIllegal                             //签名不合法",
	ErrSmsMobileNumberIllegal                             //手机号码格式错误",
	ErrSmsMobileCountOverLimit                            //手机号频率受限",
	ErrSmsTemplateMissingParameters                       //"短信模板变量缺少参数",
	ErrSmsBusinessLimitControl                            //触发业务流控限制",
	ErrSmsSystemError                                     //系统错误",
	ErrSmsSignatureNonceUsed                              //唯一随机数重复",
	ErrSmsUnknownResponse                                 //未知的响应错误
	ErrSmsBlackKeyControlLimit                            //模板变量中存在黑名单关键字",
	ErrSmsParamNotSupportUrl                              //不支持url为变量",
	ErrSmsParamLengthLimit                                //变量长度受限",
	ErrSmsAmountNotEnough                                 //余额不足",
	ErrSmsSysBusy                                         //验证码获取过于频繁，请稍后再试",
	ErrSmsLimitMinute                                     //验证码获取过于频繁，请稍后再试",
	ErrSmsLimitHour                                       //验证码获取过于频繁，请稍后再试",
)
