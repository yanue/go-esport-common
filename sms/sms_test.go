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
	accessKeyId := "LTAIiCR1y6RAa2IC"
	accessKeySecret := "pit2WJpgdhSwOEhzr42EdlXMuTdhpn"
	signName := "智享协同"

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

	sms := NewSms(accessKeyId, accessKeySecret, signName, client)
	err1 := sms.SendCode("18503002165", SmsCodeTypeBind, "112")
	fmt.Println("err", err1, errcode.GetErrMsg(err1))
}
