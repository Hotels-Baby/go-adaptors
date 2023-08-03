package factory

import (
	"fmt"
	"github.com/hotels-baby/go-adaptors/logging/google"
	"github.com/hotels-baby/go-adaptors/logging/zap"
)

// Logger types.
type LoggerType string

const (
	LoggerTypeZap    LoggerType = "Zap"    // Represents a Zap Logger type
	LoggerTypeGoogle LoggerType = "Google" // Represents a Google Logger type
)

// Config struct
type Config struct {
	LogFilePath     string // This is required
	GoogleProjectID string // This is optional
}

// Logger interface
type Logger interface {
	Error(message string, err error)
	Info(message string, fields ...interface{})
}

// NewLogger creates a new logger based on the provided type.
func NewLogger(t LoggerType, config Config) (Logger, error) {
	switch t {
	case LoggerTypeZap:
		return zap.NewZapLoggerAdapter(config.LogFilePath)
	case LoggerTypeGoogle:
		if config.GoogleProjectID == "" {
			return nil, fmt.Errorf("google logger requires a Google Project ID")
		}
		return google.NewGoogleLoggerAdapter(config.LogFilePath, config.GoogleProjectID)
	default:
		return nil, fmt.Errorf("unsupported logger type: %s", t)
	}
}
