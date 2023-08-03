package logging

import "github.com/hotels-baby/go-adaptors/logging/factory"

// Logger interface
type Logger = factory.Logger

// LoggerType type
type LoggerType = factory.LoggerType

// Config struct
type Config = factory.Config

// Logger types.
const (
	LoggerTypeZap    = factory.LoggerTypeZap    // Represents a Zap Logger type
	LoggerTypeGoogle = factory.LoggerTypeGoogle // Represents a Google Logger type
)

// NewLogger creates a new logging client of the specified type. The type parameter t must be one of the defined LoggerType values: Zap or Google.
// If the logger type is unknown, the function returns an error.
// The provided config c is used to initialise the logger, GoogleProjectID is only required when using Google LoggerType.
// LogFilePath is the path to the log file, note that .log is appended automatically, It is used for the log name when using google. LogLevel is the level of logging
// Current log levels that can be used are INFO and ERROR
func NewLogger(t LoggerType, c Config) (Logger, error) {
	logger, err := factory.NewLogger(t, c)
	if err != nil {
		return nil, err
	}

	return logger, nil
}
