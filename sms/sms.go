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
signName = "大鱼测试"
yunpianApiKey := "11111"

sms := NewSms(accessKeyId, accessKeySecret, signName, client)
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

type SmsLanguage int

// 语言类型
const (
	SmsLanguage_CN SmsLanguage = iota // 简体中文
	SmsLanguage_TW                    // 繁体中文
	SmsLanguage_EN                    // 英文
	SmsLanguage_KR                    // 韩文
)

const (
	// 阿里大于短信模板
	smsCommon_CN = "SMS_149405013"
)

const (
	// 云片短信模板
	YunpianSmsCommon_TW = "1981084"
	YunpianSmsCommon_EN = "1981092"
	YunpianSmsCommon_KR = "1981082"
)

// 国家码
type AreaCode string

const (
	// AreaCode_CN +86
	AreaCode_CN AreaCode = validator.AreaCode_CN
	AreaCode_HK AreaCode = validator.AreaCode_HK
	AreaCode_TW AreaCode = validator.AreaCode_TW
	AreaCode_MO AreaCode = validator.AreaCode_MO
	AreaCode_US AreaCode = validator.AreaCode_US
)

// 短信类型
type CodeType string

const (
	// SmsCodeTypeReg 注册
	SmsCodeTypeReg CodeType = "reg"
	// SmsCodeTypeReg 快捷登陆
	SmsCodeTypeQuickLogin CodeType = "quick_login"
	// SmsCodeTypeResetPass 重置密码
	SmsCodeTypeResetPass CodeType = "reset_pass"
	// SmsCodeTypeBind 手机号绑定
	SmsCodeTypeBind CodeType = "bind"
)

const (
	// 用于控制短信发送频度key 没法
	smsRedisKeyImeiMinute = "sms:imei:minute:" // 每分钟1条
	smsRedisKeyImeiHour   = "sms:imei:hour:"   // 每小时6条
)

// 获取redis中短信相关的key值
var smsRedisKey = map[CodeType]string{
	// SmsCodeTypeReg 短信注册码前缀
	SmsCodeTypeReg: "sms:reg:",
	// SmsCodeTypeResetPass 密码重置验证码前缀
	SmsCodeTypeResetPass: "sms:pass:",
	// SmsCodeTypeBind 修改绑定手机验证码前缀
	SmsCodeTypeBind: "sms:bind:",
	// SmsCodeTypeQuickLogin 快捷登陆
	SmsCodeTypeQuickLogin: "sms:login:",
}

type SmsUtil struct {
	*smsSdk
	redis *redis.Client
}

/*
 *@note 短信通知(频率限制)
 *@param appKey 阿里短信appKey
 *@param appSecret 阿里短信appSecret
 *@param signName 阿里短信signName
 *@param yunpianApiKey 云片短信apikey
 *@param redisClient *redis.Client
 *@return SmsUtil
 */
func NewSms(accessKeyId, accessKeySecret, signName string, redisClient *redis.Client) *SmsUtil {
	return &SmsUtil{
		smsSdk: &smsSdk{
			aliSdk: &AliSmsSdk{
				accessKeyId:     accessKeyId,
				accessKeySecret: accessKeySecret,
				signName:        signName,
			},
			yunpianApiKey: "", // 云片短信的支持 todo
		},
		redis: redisClient,
	}
}

/*
 *@note 发送短信
 *@param phone 手机号
 *@param codeType 短信类型
 *@param signName 短信签名
 *@param imei 设备号 - 设备号用于控制短信发送频度(app触发的短信必须带上imei)
 *@param area 国家码
 *@return SmsUtil
 */
func (this *SmsUtil) SendCode(phone string, codeType CodeType, imei string) int32 {
	// 检查手机号
	if errno := validator.Verify.IsPhone(phone); errno > 0 {
		return errno
	}

	// 未带imei参数
	if imei != "" {
		if errno := this.checkSmsLimit(imei); errno > 0 {
			return errno
		}
	}

	// 生成验证码
	code := this.randNumber(6)

	// 保存验证码信息
	err := this.saveVerifyCode(codeType, phone, code)
	if err != nil {
		return errcode.ErrCommonRedis
	}

	// 发送验证码
	errCode := this.smsSdk.sendCommonCode(phone, code, AreaCode_CN, SmsLanguage_CN)
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

	vKey, ok := smsRedisKey[codeType]
	if !ok {
		return false
	}

	// 前缀+phone
	pKey := vKey + phone

	// 获取数据
	verCode, _ := this.redis.Get(pKey).Result()

	if code == verCode {
		// 输入正确才删除
		if isDelete {
			this.redis.Del(pKey)
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
		// 前缀+phone
		_, err := this.redis.Set(key+phone, code, VerifyCodeExpire*time.Second).Result()
		return err
	}

	return errors.New("未知的验证码类型")
}

/**
@note 检查是否频繁发送验证码
 */
func (this *SmsUtil) checkSmsLimit(imei string) int32 {
	// 一个手机设备号每分钟最多发送1条短信
	count, _ := this.redis.Get(smsRedisKeyImeiMinute + imei).Int()
	if count > 0 {
		return errcode.ErrSmsLimitMinute
	}

	// 一个手机设备号每小时最多发送6条短信
	count, _ = this.redis.Get(smsRedisKeyImeiHour + imei).Int()
	if count > 6 {
		return errcode.ErrSmsLimitHour
	}

	return errcode.No_Error
}

/**
@note 更新最后一次发送验证码
 */
func (this *SmsUtil) renewSmsLimit(imei string) {
	// 设置1分钟发送验证码信息
	this.redis.Set(smsRedisKeyImeiMinute+imei, 1, 60*time.Second)

	// 设置1小时发送验证码信息
	hKey := smsRedisKeyImeiHour + imei

	hourCnt, _ := this.redis.Get(hKey).Int()
	if hourCnt > 0 {
		origExpire, _ := this.redis.TTL(hKey).Result()
		if origExpire <= 0 {
			origExpire = 3600
		}
		// 剩下时间 origExpire
		this.redis.Set(hKey, hourCnt+1, origExpire)
	} else {
		// 一小时
		this.redis.Set(hKey, 1, 3600*time.Second)
	}
}
