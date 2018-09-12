package logger

import (
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
)

// Config logger config
// use lumberjack writing logs to rolling files.
// Level as AtomicLevel is an atomically changeable, dynamic logging level.
//
type Config struct {
	// 日志切割配置
	Rotation *lumberjack.Logger `json:"rotation" yaml:"rotation"`
	// 日志等级
	Level string `json:"level" yaml:"level"`
	// 使能开发模式
	Development bool `json:"development" yaml:"development"`
	// 使能切割日志
	Rolling bool `json:"rolling" yaml:"rolling"`
	// 日志等级
	level zap.AtomicLevel
}

// Take effect config
//
// @return error
//
func (this *Config) TakeEffect() error {
	return this.level.UnmarshalText([]byte(this.Level))
}

// Return a default lumberjack logger config
//
// @return *lumberjack.Logger
//
func NewLoggerRotation() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s.log", os.Args[0]),
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   true,
	}
}
