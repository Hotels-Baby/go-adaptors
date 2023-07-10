package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/hotels-baby/go-adaptors/logging"
)

type LoggerAdapter struct {
	logger *zap.Logger
}

func NewZapLoggerAdapter(logName string, logLevel string) (logging.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// set the logger level
	var zapLevel zapcore.Level
	switch logLevel {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.InfoLevel
	}
	config.Level.SetLevel(zapLevel)

	// Customize the logger to write logs to a file
	config.OutputPaths = append(config.OutputPaths, logName+".log")

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &LoggerAdapter{
		logger: logger,
	}, nil
}

func (l *LoggerAdapter) Error(message string, err error) {
	l.logger.Error(message, zap.Error(err))
}

func (l *LoggerAdapter) Info(message string, fields ...interface{}) {
	l.logger.Info(message, convertFieldsToZapFields(fields)...)
}

func (l *LoggerAdapter) Close() error {
	return l.logger.Sync()
}

func convertFieldsToZapFields(fields []interface{}) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields)/2)
	for i := 0; i < len(fields)-1; i += 2 {
		key, ok := fields[i].(string)
		if !ok {
			continue
		}
		zapFields = append(zapFields, zap.Any(key, fields[i+1]))
	}
	return zapFields
}
