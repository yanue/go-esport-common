/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : config.go
 Time    : 2018/9/12 17:48
 Author  : yanue
 
 - 配置文件: 相关配置信息
 
------------------------------- go ---------------------------------*/

package common

import (
	"go.uber.org/zap"
	"sync"
)

const (
	EnvProduction = "prod"
	EnvTest       = "test"
)

var (
	AppPath        = ""                  // 程序启动目录
	LogPath        = "../logs"           // 默认的文本日志生成目录
	LogLevel       = zap.DebugLevel      // 暴露日志等级给外部读取,注意必须是: zapcore.Level
	ConfigEnv      = ""                  // 启动参数 EnvProduction|EnvTest
	ElasticIndex   = ""                  // 用于创建Elasticsearch 索引名称
	CurrentService = ""                  // 当前服务
	CurrentIP      = ""                  // 当前服务IP地址
	bootstrap      = make(chan struct{}) // 初始化阻塞Chan
	bootstrapOnce  = sync.Once{}
	bootstrapGroup = new(sync.WaitGroup) // 等待初始化的队列
)
