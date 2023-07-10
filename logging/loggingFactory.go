package logging

import (
	"fmt"
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

// NewLogger creates a new logger based on the provided type.
func NewLogger(t LoggerType, logFilePath string, logLevel string, googleProjectID string) (Logger, error) {
	switch t {
	case LoggerTypeZap:
		return zap.NewZapLoggerAdapter(logFilePath, logLevel)
	case LoggerTypeGoogle:
		return google.NewGoogleLoggerAdapter(logFilePath, logLevel, googleProjectID)
	default:
		return nil, fmt.Errorf("unsupported logger type: %s", t)
	}
}
