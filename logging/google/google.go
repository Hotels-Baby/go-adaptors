package google

import (
	google "cloud.google.com/go/logging"
	"context"
	"github.com/hotels-baby/go-adaptors/logging"
	"strings"
)

type LoggerAdapter struct {
	logger *google.Logger
	level  string
}

func NewGoogleLoggerAdapter(logName string, logLevel string, projectID string) (logging.Logger, error) {
	client, err := google.NewClient(context.Background(), projectID)
	if err != nil {
		return nil, err
	}

	logger := client.Logger(logName)
	level := strings.ToLower(logLevel)

	return &LoggerAdapter{
		logger: logger,
		level:  level,
	}, nil
}

func (l *LoggerAdapter) log(severity google.Severity, message string, fields ...interface{}) {
	if l.level == "debug" && severity < google.Debug {
		return
	}
	if l.level == "info" && severity < google.Info {
		return
	}

	payload := make(map[string]interface{})
	payload["message"] = message
	for i := 0; i < len(fields); i += 2 {
		key, ok := fields[i].(string)
		if !ok {
			continue
		}
		payload[key] = fields[i+1]
	}
	l.logger.Log(google.Entry{Severity: severity, Payload: payload})
}

func (l *LoggerAdapter) Error(message string, err error) {
	l.log(google.Error, message+": "+err.Error())
}

func (l *LoggerAdapter) Warning(message string, fields ...interface{}) {
	l.log(google.Warning, message, fields...)
}

func (l *LoggerAdapter) Info(message string, fields ...interface{}) {
	l.log(google.Info, message, fields...)
}

func (l *LoggerAdapter) Debug(message string, fields ...interface{}) {
	l.log(google.Debug, message, fields...)
}

func (l *LoggerAdapter) Close() error {
	return l.logger.Flush()
}
