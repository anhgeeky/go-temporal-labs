package main

import (
	"github.com/anhgeeky/go-temporal-labs/core/broker/kafka"
)

func main() {
	// ======================= BROKER =======================
	kafka.ConnectBrokerKafka("localhost:9092,localhost:9093,localhost:9094")
	// ======================= BROKER =======================
}
