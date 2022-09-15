package logger

import (
	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	opts := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	}
	logger, err := zap.NewProduction(opts...)
	if err != nil {
		panic("zap.NewProduction could not initialize logger")
	}
	return logger
}
