package kafka

import (
	"context"
	"log"
	"strings"

	"github.com/anhgeeky/go-temporal-labs/core/broker"
)

func ConnectBrokerKafka(brokers string) broker.Broker {
	log.Println("Broker Kafka starting...", brokers)
	// ======================= BROKER =======================
	var config = &KafkaBrokerConfig{
		Addresses: strings.Split(brokers, ","),
	}

	// cLogger := logrus.NewLogrusLogger(
	// 	pkgLogger.WithLevel(pkgLogger.InfoLevel),
	// )

	br, err := GetKafkaBroker(
		config,
		// broker.WithLogger(cLogger),
	)

	if err != nil {
		log.Fatal(context.TODO(), "Failted to create kafka broker")
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
