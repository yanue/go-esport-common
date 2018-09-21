/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : config.go
 Time    : 2018/9/12 17:48
 Author  : yanue
 
 - 配置文件: 相关配置信息
 
------------------------------- go ---------------------------------*/

package common

import (
	pb "github.com/yanue/go-esport-common/proto"
	"go.uber.org/zap"
)

const (
	// EnvProduction
	EnvProduction = "prod" //生产环境
	EnvTest       = "test" //测试环境

	ServiceNameAccount = "account" // 账号相关微服务名

	//MicroServicePrefix 微服务前缀
	MicroServicePrefix      = "go.micro.service."
	MicroServiceNameAccount = MicroServicePrefix + ServiceNameAccount // account微服务名
)

const (
	OS_ANDROID pb.Os = pb.Os_ANDROID
	OS_IOS     pb.Os = pb.Os_IOS
	OS_WEB     pb.Os = pb.Os_WEB
)

var (
	// LogPath 默认的文本日志生成目录
	LogPath = "../logs"
	// LogLevel 暴露日志等级给外部读取,注意必须是: zapcore.Level
	LogLevel = zap.DebugLevel
	// ConfigEnv 启动参数 EnvProduction|EnvTest
	ConfigEnv = ""

	// Logs 用于外部使用日志
	Logs *zap.SugaredLogger
)
