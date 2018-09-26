/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : send.go
 Time    : 2018/9/25 15:47
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

/*
// Example: 发送手机短信

appKey = "11111111"
appSecret = "22222222222222222222222222222222"
smsFreeSignName = "大鱼测试"
yunpianApiKey := "11111"

sms := NewSms(appKey, appSecret, smsFreeSignName, yunpianApiKey, client)
err1 := sms.SendCode("13800000000", SmsCodeTypeBind, "112", AreaCode_CN, SmsLanguage_CN)
fmt.Println("err", err1, errcode.GetErrMsg(err1))
*/

package sms

import (
	"errors"
	"github.com/go-redis/redis"
	"github.com/yanue/go-esport-common/errcode"
	"github.com/yanue/go-esport-common/validator"
	"time"
)

const (
	// 短信验证码有效时间:20分钟
	VerifyCodeExpire = 20 * 60
)

const (
	// 阿里大于短信模板
	smsCommon_CN = "SMS_43130006"
)

const (
	// 云片短信模板
	YunpianSmsCommon_TW = "1981084"
	YunpianSmsCommon_EN = "1981092"
	YunpianSmsCommon_KR = "1981082"
)

type AreaCode string

const (
	AreaCode_CN AreaCode = validator.AreaCode_CN
	AreaCode_HK AreaCode = validator.AreaCode_HK
	AreaCode_TW AreaCode = validator.AreaCode_TW
	AreaCode_MO AreaCode = validator.AreaCode_MO
	AreaCode_US AreaCode = validator.AreaCode_US
)

// 短信类型
type CodeType string

const (
	SmsCodeTypeReg        CodeType = "reg"
	SmsCodeTypeQuickLogin CodeType = "quick_login"
	SmsCodeTypeResetPass  CodeType = "reset_pass"
	SmsCodeTypeBind       CodeType = "bind"
)

const (
	smsRedisKeyImei = "sms:imei:"
)

// 获取redis中短信相关的key值
var smsRedisKey = map[CodeType]string{
	// PrefixRegVerifyCode 短信注册验证码前缀
	SmsCodeTypeReg: "sms:regvercode:",
	// PrefixPwdVerifyCode 密码重置验证码前缀
	SmsCodeTypeResetPass: "sms:pwdvercode:",
	// PrefixBindVerifyCode 修改绑定手机验证码前缀
	SmsCodeTypeBind: "sms:bindvercode:",
	// PrefixQuickLoginVerifyCode 快捷登陆
	SmsCodeTypeQuickLogin: "sms:quickloginvercode:",
}

type SmsUtil struct {
	smsSdk
	redis *redis.Client
}

/*
 *@note 短信通知(频率限制)
 *@param appKey 阿里短信appKey
 *@param appSecret 阿里短信appSecret
 *@param smsFreeSignName 阿里短信smsFreeSignName
 *@param yunpianApiKey 云片短信apikey
 *@param redisClient *redis.Client
 *@return SmsUtil
 */
func NewSms(appKey, appSecret, smsFreeSignName, yunPianApiKey string, redisClient *redis.Client) *SmsUtil {
	return &SmsUtil{
		smsSdk: smsSdk{
			AppKey:          appKey,
			AppSecret:       appSecret,
			SmsFreeSignName: smsFreeSignName,
			YunpianApiKey:   yunPianApiKey,
		},
		redis: redisClient,
	}
}

/*
 *@note 发送短信
 *@param phone 手机号
 *@param codeType 短信类型
 *@param smsFreeSignName 短信签名
 *@param redisPool
 *@return SmsUtil
 */
func (this *SmsUtil) SendCode(phone string, codeType CodeType, imei string, area AreaCode, lang SmsLanguage) int32 {
	// 检查手机号
	if errno := validator.Verify.IsPhone(phone, string(area)); errno > 0 {
		return errno
	}

	// 生成验证码
	code := this.randNumber(6)

	// 未带imei参数
	if imei != "" {
		pass := this.isPassSmsCheck(this.redis, imei)
		if pass == false {
			return errcode.Err_Sms_MOBILE_COUNT_OVER_LIMIT
		}
	}

	// 保存验证码信息
	err := this.saveVerifyCode(codeType, phone, code)
	if err != nil {
		return errcode.ErrCommonRedis
	}

	// 发送验证码
	errCode := this.smsSdk.SendCommonCode(phone, code, area, lang)
	if errCode == errcode.No_Error && imei != "" {
		// 刷新验证码发送限制
		this.renewSmsLimit(imei)
	}

	return errCode
}

/*
 *@note 验证验证码是否正确
 *@param phone 手机号
 *@param code 验证码
 *@param codeType 短信类型
 *@param isDelete 是否删除短信
 *@return SmsUtil
 */
func (this *SmsUtil) VerifyCode(phone, code string, codeType CodeType, isDelete bool) bool {
	if code == "" {
		return false
	}

	vKey, ok := smsRedisKey[codeType];
	if !ok {
		return false
	}

	verCode, _ := this.redis.Get(vKey).Result()
	if code == verCode {
		if isDelete {
			this.redis.Del(vKey)
		}
		return true
	}

	return false
}

/**
@note 保存验证码
 */
func (this *SmsUtil) saveVerifyCode(codeType CodeType, phone, code string) error {
	if key, ok := smsRedisKey[codeType]; ok {
		_, err := this.redis.Set(key, code, VerifyCodeExpire*time.Second).Result()
		return err
	}

	return errors.New("未知的验证码类型")
}

/**
@note 检查是否频繁发送验证码
 */
func (this *SmsUtil) isPassSmsCheck(conn *redis.Client, imei string) bool {
	// 一个手机设备号每分钟最多发送1条短信
	count, _ := conn.Get(smsRedisKeyImei + "min:" + imei).Int()
	if count > 0 {
		return false
	}

	// 一个手机设备号每小时最多发送6条短信
	count, _ = conn.Get(smsRedisKeyImei + ":hour" + imei).Int()
	if count > 6 {
		return false
	}

	return true
}

/**
@note 更新最后一次发送验证码
 */
func (this *SmsUtil) renewSmsLimit(imei string) {
	// 设置1分钟发送验证码信息
	this.redis.Set(smsRedisKeyImei+"min:"+imei, 1, 60*time.Second)

	// 设置1小时发送验证码信息
	hKey := smsRedisKeyImei + "hour:" + imei
	hourCnt, _ := this.redis.Get(hKey).Int()
	if hourCnt > 0 {
		origExpire, _ := this.redis.TTL(hKey).Result()
		if origExpire <= 0 {
			origExpire = 3600
		}
		this.redis.Set(hKey, hourCnt+1, origExpire*time.Second)
	} else {
		this.redis.Set(hKey, hourCnt+1, 3600*time.Second)
	}
}
