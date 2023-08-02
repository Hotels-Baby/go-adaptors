package logging

import (
	"github.com/hotels-baby/go-adaptors/logging/factory"
	"github.com/hotels-baby/go-adaptors/logging/interfaces"
)

type Config struct {
	LogFilePath     string
	LogLevel        string
	GoogleProjectID string // This is optional
}

// LoggerType represents the type of logger.
type LoggerType string

// Logger types.
const (
	LoggerTypeZap    LoggerType = "Zap"
	LoggerTypeGoogle LoggerType = "Google"
)

type Client struct {
	logger interfaces.Logger
	config Config
}

func NewClient(t LoggerType, c Config) (*Client, error) {
	logger, err := factory.NewLogger(t, c)
	if err != nil {
		return nil, err
	}

	return &Client{
		logger: logger,
		config: c,
	}, nil
}
