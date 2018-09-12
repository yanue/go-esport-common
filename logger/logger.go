package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New zap Logger with Config and zap option
//
// @param config *Config
// @param opts ...zap.Option
//
// @return *zap.Logger
//
func New(config *Config, opts ...zap.Option) *zap.Logger {

	if config == nil {
		log, err := zap.NewDevelopment(opts...)
		if err != nil {
			panic(err)
		}
		return log
	}

	err := config.TakeEffect()

	if err != nil {
		panic(err)
	}

	if config.Rolling {
		if config.Rotation == nil {
			config.Rotation = NewLoggerRotation()
		}
	}

	var zapConfig zap.Config

	if config.Rotation != nil && config.Rolling {

		var enc zapcore.Encoder

		if config.Development {
			zapConfig = zap.NewDevelopmentConfig()
			enc = zapcore.NewConsoleEncoder(zapConfig.EncoderConfig)
		} else {
			zapConfig = zap.NewProductionConfig()
			enc = zapcore.NewJSONEncoder(zapConfig.EncoderConfig)
		}

		ws := zapcore.AddSync(config.Rotation)

		core := zapcore.NewCore(
			enc,
			ws,
			config.level,
		)

		log := zap.New(core)

		if len(opts) > 0 {
			log = log.WithOptions(opts...)
		}

		return log
	}

	if config.Development {
		zapConfig = zap.NewDevelopmentConfig()

	} else {
		zapConfig = zap.NewProductionConfig()
	}

	log, err := zapConfig.Build(opts...)

	if err != nil {
		panic(err)
	}

	return log
}
