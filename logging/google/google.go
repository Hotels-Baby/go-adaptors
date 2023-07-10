package google

import (
	"adaptors/logging"
	google "cloud.google.com/go/logging"
	"context"
)

type GoogleLoggerAdapter struct {
	logger *google.Logger
}

func NewGoogleLoggerAdapter(ctx context.Context, projectID string, logName string) (logging.Logger, error) {
	client, err := google.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	logger := client.Logger(logName)

	return &GoogleLoggerAdapter{
		logger: logger,
	}, nil
}

func (l *GoogleLoggerAdapter) Error(message string, err error) {
	l.logger.Log(google.Entry{Severity: google.Error, Payload: message + ": " + err.Error()})
}

func (l *GoogleLoggerAdapter) Warning(message string, fields ...interface{}) {
	payload := make(map[string]interface{})
	payload["message"] = message
	for i := 0; i < len(fields); i += 2 {
		key, ok := fields[i].(string)
		if !ok {
			continue
		}
		payload[key] = fields[i+1]
	}
	l.logger.Log(google.Entry{Severity: google.Warning, Payload: payload})
}

func (l *GoogleLoggerAdapter) Info(message string, fields ...interface{}) {
	payload := make(map[string]interface{})
	payload["message"] = message
	for i := 0; i < len(fields); i += 2 {
		key, ok := fields[i].(string)
		if !ok {
			continue
		}
		payload[key] = fields[i+1]
	}
	l.logger.Log(google.Entry{Severity: google.Info, Payload: payload})
}

func (l *GoogleLoggerAdapter) Close() error {
	return l.logger.Flush()
}
