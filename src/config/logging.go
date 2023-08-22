package config

import "go.uber.org/zap"

// NewLogger Creates new logger instance
func NewLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()

	return slogger
}
