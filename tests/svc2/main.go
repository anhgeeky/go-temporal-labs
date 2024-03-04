package main

import (
	"github.com/anhgeeky/go-temporal-labs/core/broker/kafka"
)

func main() {
	// ======================= BROKER =======================
	kafka.ConnectBrokerKafka("127.0.0.1:9092,127.0.0.1:9093,127.0.0.1:9094")
	// ======================= BROKER =======================
}
