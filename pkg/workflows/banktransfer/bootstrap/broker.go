package bootstrap

import (
	"context"
	"strings"

	"github.com/anhgeeky/go-temporal-labs/core/broker"
	"github.com/anhgeeky/go-temporal-labs/core/broker/kafka"
	"github.com/anhgeeky/go-temporal-labs/core/config"
	"github.com/anhgeeky/go-temporal-labs/core/logger"
)

func GetKafkaBroker(cfg config.Configure, logger logger.Logger) broker.Broker {
	var addrs = cfg.GetString("KAFKA_BROKERS")
	var config = &kafka.KafkaBrokerConfig{
		Addresses: strings.Split(addrs, ","),
	}

	br, err := kafka.GetKafkaBroker(
		config,
		broker.WithLogger(logger),
	)

	if err != nil {
		logger.Error(context.TODO(), "Failted to create kafka broker")
		panic(err)
	}

	return br
}
