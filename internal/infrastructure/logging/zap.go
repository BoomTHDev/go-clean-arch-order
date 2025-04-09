package logging

import (
	"go.uber.org/zap"
)

type (
	ZapLogger struct {
		logger *zap.Logger
	}
)

func NewZapLogger() (*ZapLogger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"

	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	return &ZapLogger{
		logger: logger,
	}, nil
}

func (l *ZapLogger) Error(message any, fields ...any) {
	zapFields := []zap.Field{}

	for i := 0; i < len(fields); i += 2 {
		if i+1 < len(fields) {
			key := fields[i].(string)
			value := fields[i+1]
			zapFields = append(zapFields, zap.Any(key, value))
		}
	}

	switch v := message.(type) {
	case error:
		l.logger.Error(v.Error(), zapFields...)
	case string:
		l.logger.Error(v, zapFields...)
	}
}
