package common

import "sync"

var (
	AppPath         = ""                  // 程序启动目录
	LogPath         = "./logs"            // 默认的文本日志生成目录
	ConfigEnv       = ""                  // 启动参数
	ReadConsulInter = 10                  // 定时读取consul的间隔
	ElasticIndex    = ""                  // 用于创建Elasticsearch 索引名称
	SepLogLevel     = 7                   // 暴露日志等级给外部读取
	CurrentService  = ""                  // 当前服务
	CurrentIP       = ""                  // 当前服务IP地址
	bootstrap       = make(chan struct{}) // 初始化阻塞Chan
	bootstrapOnce   = sync.Once{}
	bootstrapGroup  = new(sync.WaitGroup) // 等待初始化的队列
)
