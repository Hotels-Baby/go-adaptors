package factory

import (
	"fmt"
	"github.com/hotels-baby/go-adaptors/logging/interfaces"

	"github.com/hotels-baby/go-adaptors/logging/google"
	"github.com/hotels-baby/go-adaptors/logging/zap"
)

// LoggerType represents the type of logger.
type LoggerType string

// Logger types.
const (
	LoggerTypeZap    LoggerType = "dev"
	LoggerTypeGoogle LoggerType = "prod"
)

type LoggerConfig struct {
	LogFilePath     string
	LogLevel        string
	GoogleProjectID string // This is optional
}

// NewLogger creates a new logger based on the provided type.
func NewLogger(t LoggerType, config *LoggerConfig) (interfaces.Logger, error) {
	switch t {
	case LoggerTypeZap:
		return zap.NewZapLoggerAdapter(config.LogFilePath, config.LogLevel)
	case LoggerTypeGoogle:
		if config.GoogleProjectID == "" {
			return nil, fmt.Errorf("google logger requires a Google Project ID")
		}
		return google.NewGoogleLoggerAdapter(config.LogFilePath, config.LogLevel, config.GoogleProjectID)
	default:
		return nil, fmt.Errorf("unsupported logger type: %s", t)
	}
}
