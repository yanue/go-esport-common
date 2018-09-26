/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : sms_test.go
 Time    : 2018/9/25 18:42
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package sms

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/yanue/go-esport-common/errcode"
	"testing"
)

func TestNewSmsUtil(t *testing.T) {
	appKey := "11111111"
	appSecret := "22222222222222222222222222222222"
	smsFreeSignName := "大鱼测试"
	yunpianApiKey := "11111"

	// 建立连接
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.5.201:6379",
		Password: "",
		DB:       0,
		PoolSize: 300, // 连接池大小
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	_, err := client.Ping().Result()
	if err != nil {
		panic("redis连接失败:" + err.Error())
	}

	//common.Logs.Info("redis connected.")

	sms := NewSms(appKey, appSecret, smsFreeSignName, yunpianApiKey, client)
	err1 := sms.SendCode("13800000000", SmsCodeTypeBind, "112", AreaCode_CN, SmsLanguage_CN)
	fmt.Println("err", err1, errcode.GetErrMsg(err1))
}