package config

import (
	"fmt"

	"go.uber.org/zap"
)

type LoggingConfig struct {
	Level zap.AtomicLevel
	IsDev bool
}

func (c LoggingConfig) CreateLogger() *zap.Logger {
	zapConfig := zap.NewProductionConfig()
	if c.IsDev {
		zapConfig = zap.NewDevelopmentConfig()
	}
	zapConfig.Level = c.Level

	logger, err := zapConfig.Build()
	if err != nil {
		panic(fmt.Sprintf("Failed to construct logger: %e", err))
	}
	return logger
}

func parseLogLevel(value string) (*zap.AtomicLevel, error) {
	level, err := zap.ParseAtomicLevel(value)
	return &level, err
}

func createLoggingConfig() LoggingConfig {
	return LoggingConfig{
		Level: getEnv("LOG_LEVEL", false, parseLogLevel),
		IsDev: getEnv("LOG_DEV", false, parseBool),
	}
}
