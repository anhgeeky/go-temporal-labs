package config

import (
	"log"

	"github.com/anhgeeky/go-temporal-labs/core/kafka"
	kk "github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

var KafkaConsumer *kk.Reader
var KafkaProducer *kk.Writer

func NewKafkaConsumer() {
	kkServers := viper.GetString("KAFKA_BROKERS")
	kkTopic := viper.GetString("KAFKA_TOPIC")
	kkGroup := viper.GetString("KAFKA_GROUP")
	KafkaConsumer = kafka.NewReader(kkServers, kkTopic, kkGroup, false)
	log.Println("Kafka consumer has been created")
	defer KafkaConsumer.Close()
}

func NewKafkaProducer() {
	kkServers := viper.GetString("KAFKA_BROKERS")
	kkTopic := viper.GetString("KAFKA_TOPIC")
	KafkaProducer = kafka.NewWriter(kkServers, kkTopic)
	log.Println("Kafka producer has been created")
	defer KafkaProducer.Close()
}
