package google

import (
	"cloud.google.com/go/logging"
	"context"
	"search_data-crawler_crawler-api/pkg/ports/out"
)

type GoogleLoggerAdapter struct {
	logger *logging.Logger
}

func NewGoogleLoggerAdapter(ctx context.Context, projectID string, logName string) (out.Logger, error) {
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	logger := client.Logger(logName)

	return &GoogleLoggerAdapter{
		logger: logger,
	}, nil
}

func (l *GoogleLoggerAdapter) Error(message string, err error) {
	l.logger.Log(logging.Entry{Severity: logging.Error, Payload: message + ": " + err.Error()})
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
	l.logger.Log(logging.Entry{Severity: logging.Warning, Payload: payload})
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
	l.logger.Log(logging.Entry{Severity: logging.Info, Payload: payload})
}

func (l *GoogleLoggerAdapter) Close() error {
	return l.logger.Flush()
}
