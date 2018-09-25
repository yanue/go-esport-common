/*go**************************************************************************
 File            : SmsRedisKey.go
 Subsystem       :
 Author          : hemingang
 Date&Time       : 2016-08-18
 Description     : 获取reids中短信相关的key值
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/
package sms

const (
	// PrefixSms 短信key前缀
	prefixSms = "sms:sms:"
	// PrefixRegVercode 短信注册验证码前缀
	prefixRegVercode = "sms:regvercode:"
	// PrefixPwdVercode 密码重置验证码前缀
	prefixPwdVercode = "sms:pwdvercode:"
	// PrefixBindVercode 修改绑定手机验证码前缀
	prefixBindVercode = "sms:bindvercode:"
	// PrefixQuickLoginVercode 快捷登陆
	prefixQuickLoginVercode = "sms:quickloginvercode:"
)

/**
 *@note 根据设备号 imei 获取 key
 *@param imei 设备号
 *@return string，返回对应的key值
 */
func GetImeiKey(imei string) string {
	return prefixSms + imei
}

/**
 *@note 根据手机号获取手机注册验证码key
 *@param  mobile 手机号码
 *@return string，返回对应的key值
 */
func GetRegVercodeKey(mobile string) string {
	return prefixRegVercode + mobile
}

/**
 *@note 根据手机号获取密码重置验证码key
 *@param  mobile 手机号码
 *@return string，返回对应的key值
 */
func GetPwdVercodeKey(mobile string) string {
	return prefixPwdVercode + mobile
}

/**
 *@note 根据手机号获取修改手机绑定验证码key
 *@param  mobile 手机号码
 *@return string，返回对应的key值
 */
func GetBindVercodeKey(mobile string) string {
	return prefixBindVercode + mobile
}

/**
 *@note 根据手机号获取快速登录验证码key
 *@param  mobile 手机号码
 *@return string，返回对应的key值
 */
func GetQuickLoginVercodeKey(mobile string) string {
	return prefixQuickLoginVercode + mobile
}
