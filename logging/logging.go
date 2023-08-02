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

// LoggerType represents the type of logger that can be created by NewClient.
// The two possible LoggerTypes are LoggerTypeZap and LoggerTypeGoogle.
type LoggerType string

// Logger types.
const (
	LoggerTypeZap    LoggerType = "Zap"    // Represents a Zap Logger type
	LoggerTypeGoogle LoggerType = "Google" // Represents a Google Logger type
)

type Client struct {
	logger interfaces.Logger
	config Config
}

// NewClient creates a new logging client of the specified type. The type parameter t must be one of the defined LoggerType values: Zap or Google.
// If the logger type is unknown, the function returns an error.
// The provided config c is used to initialise the logger, GoogleProjectID is only required when using Google LoggerType.
// LogFilePath is the path to the log file, note that .log is appended automatically, It is used for the log name when using google. LogLevel is the level of logging
// Current log levels that can be used are INFO and ERROR
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
