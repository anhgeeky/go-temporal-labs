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

func NewReader(kafkaURL, topic string, group string, isLastOffset bool) *kk.Reader {
	cfg := kk.ReaderConfig{
		Brokers:        strings.Split(kafkaURL, ","),
		Topic:          topic,
		GroupID:        group,
		MinBytes:       1e3,  // 10KB
		MaxBytes:       10e6, // 10MB
		MaxWait:        1 * time.Second,
		CommitInterval: time.Second, // flushes commits to Kafka every second
	}

	if isLastOffset {
		cfg.StartOffset = kk.LastOffset
	}

	return kk.NewReader(cfg)
}
