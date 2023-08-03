package google

import (
	google "cloud.google.com/go/logging"
	"context"
)

// Logger interface
type Logger interface {
	Error(message string, err error)
	Info(message string, fields ...interface{})
}

type LoggerAdapter struct {
	logger *google.Logger
}

func NewGoogleLoggerAdapter(logName string, projectID string) (Logger, error) {
	client, err := google.NewClient(context.Background(), projectID)
	if err != nil {
		return nil, err
	}

	logger := client.Logger(logName)

	return &LoggerAdapter{
		logger: logger,
	}, nil
}

func (l *LoggerAdapter) log(severity google.Severity, message string, fields ...interface{}) {
	var payload interface{}
	if len(fields) > 0 {
		payload = struct {
			Message string
			Fields  []interface{}
		}{Message: message, Fields: fields}
	} else {
		payload = message
	}
	entry := google.Entry{
		Severity: severity,
		Payload:  payload,
	}
	l.logger.Log(entry)
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
