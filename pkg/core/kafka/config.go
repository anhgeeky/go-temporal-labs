package kafka

import (
	"strings"
	"time"

	kk "github.com/segmentio/kafka-go"
)

func NewWriter(kafkaURL, topic string) *kk.Writer {
	return kk.NewWriter(kk.WriterConfig{
		Brokers:  strings.Split(kafkaURL, ","),
		Topic:    topic,
		Balancer: &kk.LeastBytes{},
	})
}

func NewReader(kafkaURL, topic string, group string) *kk.Reader {
	return kk.NewReader(kk.ReaderConfig{
		Brokers:        strings.Split(kafkaURL, ","),
		Topic:          topic,
		GroupID:        group,
		MaxBytes:       10e6,        // 10MB
		CommitInterval: time.Second, // flushes commits to Kafka every second
	})
}
