/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : sms.go
 Time    : 2018/9/25 14:57
 Author  : yanue
 
 - 验证码
 
------------------------------- go ---------------------------------*/

package errcode

var errMsgSms = map[ErrNo]string{
	// Sms
	Err_Sms_OUT_OF_SERVICE:              "业务停机",
	Err_sms_PRODUCT_UNSUBSCRIBE:         "产品服务未开通",
	Err_Sms_ACCOUNT_NOT_EXISTS:          "账户信息不存在",
	Err_Sms_ACCOUNT_ABNORMAL:            "账户信息异常",
	Err_Sms_SMS_TEMPLATE_ILLEGAL:        "模板不合法",
	Err_Sms_SMS_SIGNATURE_ILLEGAL:       "签名不合法",
	Err_Sms_MOBILE_NUMBER_ILLEGAL:       "手机号码格式错误",
	Err_Sms_MOBILE_COUNT_OVER_LIMIT:     "手机号码数量超过限制",
	Err_Sms_TEMPLATE_MISSING_PARAMETERS: "短信模板变量缺少参数",
	Err_Sms_INVALID_PARAMETERS:          "参数异常",
	Err_Sms_BUSINESS_LIMIT_CONTROL:      "触发业务流控限制",
	Err_Sms_INVALID_JSON_PARAM:          "JSON参数不合法",
	Err_Sms_SYSTEM_ERROR:                "系统错误",
	Err_Sms_BLACK_KEY_CONTROL_LIMIT:     "模板变量中存在黑名单关键字",
	Err_Sms_PARAM_NOT_SUPPORT_URL:       "不支持url为变量",
	Err_Sms_PARAM_LENGTH_LIMIT:          "变量长度受限",
	Err_Sms_AMOUNT_NOT_ENOUGH:           "余额不足",
	Err_Sms_SYS_BUSY:                    "验证码获取过于频繁，请稍后再试",
}

const (
	// Sms
	Err_Sms_OUT_OF_SERVICE              = errOffsetSms + iota //业务停机",
	Err_sms_PRODUCT_UNSUBSCRIBE                               //产品服务未开通",
	Err_Sms_ACCOUNT_NOT_EXISTS                                //账户信息不存在",
	Err_Sms_ACCOUNT_ABNORMAL                                  //账户信息异常",
	Err_Sms_SMS_TEMPLATE_ILLEGAL                              //模板不合法",
	Err_Sms_SMS_SIGNATURE_ILLEGAL                             //签名不合法",
	Err_Sms_MOBILE_NUMBER_ILLEGAL                             //手机号码格式错误",
	Err_Sms_MOBILE_COUNT_OVER_LIMIT                           //手机号码数量超过限制",
	Err_Sms_TEMPLATE_MISSING_PARAMETERS                       //"短信模板变量缺少参数",
	Err_Sms_INVALID_PARAMETERS                                //参数异常",
	Err_Sms_BUSINESS_LIMIT_CONTROL                            //触发业务流控限制",
	Err_Sms_INVALID_JSON_PARAM                                //JSON参数不合法",
	Err_Sms_SYSTEM_ERROR                                      //系统错误",
	Err_Sms_BLACK_KEY_CONTROL_LIMIT                           //模板变量中存在黑名单关键字",
	Err_Sms_PARAM_NOT_SUPPORT_URL                             //不支持url为变量",
	Err_Sms_PARAM_LENGTH_LIMIT                                //变量长度受限",
	Err_Sms_AMOUNT_NOT_ENOUGH                                 //余额不足",
	Err_Sms_SYS_BUSY                                          //验证码获取过于频繁，请稍后再试",
)
