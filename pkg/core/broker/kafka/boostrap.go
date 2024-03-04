package kafka

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/anhgeeky/go-temporal-labs/core/broker"
	pkgLogger "github.com/anhgeeky/go-temporal-labs/core/logger"
	"github.com/anhgeeky/go-temporal-labs/core/logger/logrus"
)

func ConnectBrokerKafka() broker.Broker {
	// ======================= BROKER =======================
	var config = &KafkaBrokerConfig{
		Addresses: strings.Split(os.Getenv("KAFKA_BROKERS"), ","),
	}

	cLogger := logrus.NewLogrusLogger(
		pkgLogger.WithLevel(pkgLogger.InfoLevel),
	)

	br, err := GetKafkaBroker(
		config,
		broker.WithLogger(cLogger),
	)

	if err != nil {
		cLogger.Error(context.TODO(), "Failted to create kafka broker")
		panic(err)
	}

	if err := br.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := br.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	// ======================= BROKER =======================

	log.Println("Broker Kafka connected")

	return br
}
