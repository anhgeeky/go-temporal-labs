package kafka

import "github.com/anhgeeky/go-temporal-labs/core/logger/logrus"

var (
	DefaultKafkaBroker = "127.0.0.1:9092"
	DefaultLogger      = logrus.NewLogrusLogger()
)
