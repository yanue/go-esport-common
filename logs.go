/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : logs.go
 Time    : 2018/9/12 17:13
 Author  : yanue

 - 日志处理方法(基于zap)
 - 提供gin中间件GinZapMiddleware
 - 使用方法: common.Logs.Info

------------------------------- go ---------------------------------*/

package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// LogConfig logger config
// use lumberjack writing logs to rolling files.
// Level as AtomicLevel is an atomically changeable, dynamic logging level.
//
type LogConfig struct {
	// 日志切割配置
	Rotation *lumberjack.Logger `json:"rotation" yaml:"rotation"`
	// 日志文件目录
	LogPath string `json:"log_path" yaml:"log_path"`
	// 日志等级 zap.DebugLevel
	LogLevel zapcore.Level `json:"log_level" yaml:"log_level"`
	// 使能切割日志
	Rolling bool `json:"rolling" yaml:"rolling"`
	// 使能开发模式
	Development bool `json:"development" yaml:"development"`
}

// 默认初始化一次,避免nil错误
func init() {
	initLogs()
}

// 根据配置文件进行初始化
func initLogs() {
	cfg := new(LogConfig)
	// 生产环境
	if ConfigEnv == EnvProduction {
		// 支持日志切割
		cfg.Rolling = true
		// 日志目录
		cfg.LogPath = LogPath
		// 日志切割配置
		cfg.Rotation = cfg.NewLoggerRotation()
		// 关闭开发模式
		cfg.Development = false
	} else {
		// 开发模式
		cfg.Development = true
	}

	// 日志级别
	cfg.LogLevel = LogLevel
	// 显示文件及行号
	zapOptCaller := zap.AddCaller()
	// 初始化
	Logs = NewLogger(cfg, zapOptCaller).Sugar()
}

// New zap Logger with LogConfig and zap option
//
// @param config *LogConfig
// @param opts ...zap.Option
//
// @return *zap.Logger
//
func NewLogger(config *LogConfig, opts ...zap.Option) (log *zap.Logger) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("NewLogger err:", err)
			log, _ = zap.NewDevelopment()
		}
	}()

	if config == nil {
		log, err := zap.NewDevelopment(opts...)
		if err != nil {
			panic(err)
		}
		return log
	}

	var zapConfig zap.Config

	if config.Development {
		zapConfig = zap.NewDevelopmentConfig()
	} else {
		zapConfig = zap.NewProductionConfig()
	}

	// Rotation mode
	if config.Rolling {

		if config.Rotation == nil {
			config.Rotation = config.NewLoggerRotation()
		}

		var enc zapcore.Encoder

		if config.Development {
			enc = zapcore.NewConsoleEncoder(zapConfig.EncoderConfig)
		} else {
			enc = zapcore.NewJSONEncoder(zapConfig.EncoderConfig)
		}

		ws := zapcore.AddSync(config.Rotation)

		core := zapcore.NewCore(
			enc,
			ws,
			config.LogLevel,
		)

		log := zap.New(core)

		if len(opts) > 0 {
			log = log.WithOptions(opts...)
		}

		return log
	}

	log, err := zapConfig.Build(opts...)
	if err != nil {
		panic(err)
	}

	return log
}

// Return a default lumberjack logger config
//
// @return *lumberjack.Logger
//
func (cfg *LogConfig) NewLoggerRotation() *lumberjack.Logger {
	if len(cfg.LogPath) == 0 {
		// set path
		cfg.LogPath = "../logs"
	}
	// mkdir
	if _, err := os.Stat(cfg.LogPath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		os.MkdirAll(cfg.LogPath, 0777)
	}

	return &lumberjack.Logger{
		Filename:   cfg.LogPath + "/" + fmt.Sprintf("%s.log", os.Args[0]),
		MaxSize:    500, // MB
		MaxBackups: 3,
		MaxAge:     30, // days
		Compress:   true,
	}
}

// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
//
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
//
// It receives:
//   1. A time package format string (e.g. time.RFC3339).
//   2. A boolean stating whether to use UTC time zone or local.
func GinZapMiddleware(logger *zap.Logger, timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				logger.Error(e)
			}
		} else {
			logger.Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("time", end.Format(timeFormat)),
				zap.Duration("latency", latency),
			)
		}
	}
}
