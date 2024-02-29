package bootstrap

import (
	"github.com/anhgeeky/go-temporal-labs/core/config"
	"github.com/anhgeeky/go-temporal-labs/core/logger"
	"github.com/anhgeeky/go-temporal-labs/core/logger/logrus"
	"github.com/anhgeeky/go-temporal-labs/core/trace"
)

func GetLogger(c config.Configure, tracer trace.Tracer) logger.Logger {
	levelStr := c.GetString("LOG_LEVEL")
	level, err := logger.GetLevel(levelStr)
	if err != nil {
		level = logger.InfoLevel
	}
	return logrus.NewLogrusLogger(
		logger.WithLevel(level),
		logger.WithTracer(tracer),
	)
}
