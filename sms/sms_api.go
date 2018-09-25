/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : send.go
 Time    : 2018/9/25 15:47
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package sms

import (
	"github.com/go-redis/redis"
	"github.com/yanue/go-esport-common/errcode"
	"math/rand"
	"time"
)

const (
	VercodeExpire = 300 // 短信验证码有效时间:5分钟
)

type SmsApi struct {
	CSms
	redisClent *redis.Client
}

/*
 *@note 短信通知(频率限制)
 *@param appKey
 *@param appSecret 密钥
 *@param smsFreeSignName 短信签名
 *@param redisPool
 *@return SmsApi
 */
func NewSmsApi(appKey, appSecret, smsFreeSignName string, redisClient *redis.Client) *SmsApi {
	return &SmsApi{
		CSms: CSms{
			AppKey:          appKey,
			AppSecret:       appSecret,
			SmsFreeSignName: smsFreeSignName,
		},
		redisClent: redisClient,
	}
}

/*
 *@note 发送新用户注册验证码
 *@param mobile 手机号
 *@param area 手机区域号
 *@param code 验证码
 *@param lang 语言
 *@param imei 手机设备号
 *@return 阿里返回数据, 错误信息
 */
func (this *SmsApi) SendRegisterCode(mobile, area string, lang int, imei string) int32 {
	return this.sendCode(mobile, area, lang, imei, setRegVercode, this.CSms.SendCommonCode)
}

/*
 *@note 发送重置密码验证码
 *@param mobile 手机号
 *@param area 手机区域号
 *@param code 验证码
 *@param lang 语言
 *@param imei 手机设备号
 *@return 阿里返回数据, 错误信息
 */
func (this *SmsApi) SendResetPwdCode(mobile, area string, lang int, imei string) int32 {
	return this.sendCode(mobile, area, lang, imei, setPwdVercode, this.CSms.SendCommonCode)
}

/*
 *@note 发送修改绑定手机验证码
 *@param mobile 手机号
 *@param area 手机区域号
 *@param code 验证码
 *@param lang 语言
 *@param imei 手机设备号
 *@return 阿里返回数据, 错误信息
 */
func (this *SmsApi) SendBindCode(mobile, area string, lang int, imei string) int32 {
	return this.sendCode(mobile, area, lang, imei, setBindVercode, this.CSms.SendCommonCode)
}

func (this *SmsApi) SendQuickLoginCode(mobile, area string, lang int, imei string) int32 {
	return this.sendCode(mobile, area, lang, imei, setQuickLoginVercode, this.CSms.SendCommonCode)
}

/**
发短信消息
 */
func (this *SmsApi) sendCode(mobile, area string, lang int, imei string,
	setVercodeFunc func(conn *redis.Client, mobile, code string, expire time.Duration) error,
	sendCodeFunc func(mobile, area, code string, lang int) int32) int32 {

	code := randNumber(6)
	if imei == "" {
		err := setVercodeFunc(this.redisClent, mobile, code, VercodeExpire)
		if err != nil {
			return errcode.ErrCommonRedis
		}
		return sendCodeFunc(mobile, area, code, lang)
	}

	pass := isPassSmsCheck(this.redisClent, imei)
	if pass == false {
		return errcode.Err_Sms_MOBILE_COUNT_OVER_LIMIT
	}

	err := setVercodeFunc(this.redisClent, mobile, code, VercodeExpire)
	if err != nil {
		return errcode.ErrCommonRedis
	}

	errCode := sendCodeFunc(mobile, area, code, lang)
	if errCode == errcode.No_Error {
		this.renewSmsLimit(imei)
	}

	return errCode
}

/*
 *@note 新用户注册验证码验证
 *@param mobile 手机号
 *@param code 验证码
 *@param isDelete 通过验证后是否删除数据库验证码
 *@return true 通过验证，false 验证失败
 */
func (this *SmsApi) CheckRegisterCode(mobile, code string, isDelete bool) bool {
	return this.checkCode(mobile, code, isDelete, GetRegVercodeKey)
}

/*
 *@note 密码重置验证码验证
 *@param mobile 手机号
 *@param code 验证码
 *@param isDelete 通过验证后是否删除数据库验证码
 *@return true 通过验证，false 验证失败
 */
func (this *SmsApi) CheckRetPwdCode(mobile, code string, isDelete bool) bool {
	return this.checkCode(mobile, code, isDelete, GetPwdVercodeKey)
}

/*
 *@note 修改绑定手机验证码验证
 *@param mobile 手机号
 *@param code 验证码
 *@param isDelete 通过验证后是否删除数据库验证码
 *@return true 通过验证，false 验证失败
 */
func (this *SmsApi) CheckBindCode(mobile, code string, isDelete bool) bool {
	return this.checkCode(mobile, code, isDelete, GetBindVercodeKey)
}

/*
 *@note 修改快速登录手机验证码验证
 *@param mobile 手机号
 *@param code 验证码
 *@param isDelete 通过验证后是否删除数据库验证码
 *@return true 通过验证，false 验证失败
 */
func (this *SmsApi) CheckQuickLoginCode(mobile, code string, isDelete bool) bool {
	return this.checkCode(mobile, code, isDelete, GetQuickLoginVercodeKey)
}

func (this *SmsApi) checkCode(mobile, code string, isDelete bool, getKeyFun func(mobile string) string) bool {
	if code == "" {
		return false
	}

	vKey := getKeyFun(mobile)
	verCode, _ := this.redisClent.Get(vKey).Result()
	if code == verCode {
		if isDelete {
			this.redisClent.Do("DEL", vKey)
		}
		return true
	}
	return false
}

func (this *SmsApi) renewSmsLimit(imei string) {
	this.redisClent.Set(GetImeiKey(imei)+"min", 1, 60, )
	hKey := GetImeiKey(imei) + "hour"
	hourCnt, _ := this.redisClent.Get(hKey).Int()
	if hourCnt > 0 {
		origExpire, _ := this.redisClent.TTL(hKey).Result()
		if origExpire <= 0 {
			origExpire = 3600
		}
		this.redisClent.Set(hKey, hourCnt+1, origExpire)
	} else {
		this.redisClent.Set(hKey, hourCnt+1, 3600)
	}
}

func setRegVercode(conn *redis.Client, mobile, code string, expire time.Duration) error {
	_, err := conn.Set(GetRegVercodeKey(mobile), code, expire).Result()
	return err
}

func setPwdVercode(conn *redis.Client, mobile, code string, expire time.Duration) error {
	_, err := conn.Set(GetPwdVercodeKey(mobile), code, expire).Result()
	return err
}

func setBindVercode(conn *redis.Client, mobile, code string, expire time.Duration) error {
	_, err := conn.Do(GetBindVercodeKey(mobile), code, expire).Result()
	return err
}

func setQuickLoginVercode(conn *redis.Client, mobile, code string, expire time.Duration) error {
	_, err := conn.Do(GetQuickLoginVercodeKey(mobile), code, expire).Result()
	return err
}

func isPassSmsCheck(conn *redis.Client, imei string) bool {
	return true

	// 一个手机设备号每分钟最多发送1条短信
	count, _ := conn.Get(GetImeiKey(imei) + "min").Int()
	if count > 0 {
		return false
	}

	// 一个手机设备号每小时最多发送6条短信
	count, _ = conn.Get(GetImeiKey(imei) + "hour").Int()
	if count > 6 {
		return false
	}

	return true
}

var Sms *SmsApi

const letterBytes = "0123456789"

func randNumber(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
