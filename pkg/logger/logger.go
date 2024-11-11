package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"sync"
	"time"
)

var (
	Log  *zap.Logger
	once sync.Once
)

func NewSingletonLogger() *zap.Logger {

	once.Do(func() {
		loggerConfig := zap.NewProductionConfig()
		loggerConfig.DisableStacktrace = true
		loggerConfig.EncoderConfig.TimeKey = "timestamp"
		loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

		var err error
		Log, err = loggerConfig.Build()
		if err != nil {
			log.Fatal(err)
		}
	})
	return Log

}
