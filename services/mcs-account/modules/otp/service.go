package otp

import "github.com/anhgeeky/go-temporal-labs/core/broker"

type Service struct {
	Repo   Repository
	Broker broker.Broker
}
