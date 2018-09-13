/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : config.go
 Time    : 2018/9/12 17:48
 Author  : yanue
 
 - 配置文件: 相关配置信息
 
------------------------------- go ---------------------------------*/

package common

import (
	. "go.uber.org/zap"
)

const (
	// 运行环境
	EnvProduction = "prod"
	EnvTest       = "test"

	// 微服务名
	ServiceNameAccount = "account" // 账号相关微服务名

	MicroServicePrefix      = "go.micro.service." // 微服务前缀
	MicroServiceNameAccount = MicroServicePrefix + ServiceNameAccount
)

var (
	LogPath   = "../logs"      // 默认的文本日志生成目录
	LogLevel  = DebugLevel // 暴露日志等级给外部读取,注意必须是: zapcore.Level
	ConfigEnv = ""             // 启动参数 EnvProduction|EnvTest

	Logs *SugaredLogger
)
