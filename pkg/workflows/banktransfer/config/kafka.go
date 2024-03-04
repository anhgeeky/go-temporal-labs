package config

type KafkaConfig struct {
	Brokers string `mapstructure:"KAFKA_BROKERS"`
}
