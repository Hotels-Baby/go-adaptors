package factory

import (
	"fmt"
	"github.com/hotels-baby/go-adaptors/logging"
	"github.com/hotels-baby/go-adaptors/logging/interfaces"

	"github.com/hotels-baby/go-adaptors/logging/google"
	"github.com/hotels-baby/go-adaptors/logging/zap"
)

// NewLogger creates a new logger based on the provided type.
func NewLogger(t logging.LoggerType, config logging.Config) (interfaces.Logger, error) {
	switch t {
	case logging.LoggerTypeZap:
		return zap.NewZapLoggerAdapter(config.LogFilePath, config.LogLevel)
	case logging.LoggerTypeGoogle:
		if config.GoogleProjectID == "" {
			return nil, fmt.Errorf("google logger requires a Google Project ID")
		}
		return google.NewGoogleLoggerAdapter(config.LogFilePath, config.LogLevel, config.GoogleProjectID)
	default:
		return nil, fmt.Errorf("unsupported logger type: %s", t)
	}
}
