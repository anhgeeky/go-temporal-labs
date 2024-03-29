package kafka

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/anhgeeky/go-temporal-labs/core/broker"
	"github.com/anhgeeky/go-temporal-labs/core/logger"
)

// consumerGroupHandler is the implementation of sarama.ConsumerGroupHandler
type consumerGroupHandler struct {
	logger  logger.Logger
	handler broker.Handler
	subopts broker.SubscribeOptions
	kopts   broker.BrokerOptions
	cg      sarama.ConsumerGroup
	sess    sarama.ConsumerGroupSession
	ready   chan bool
}

func (h *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	close(h.ready)
	return nil
}

func (*consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg, ok := <-claim.Messages():
			ctx := context.Background()

			if !ok {
				h.logger.Info(ctx, "message channel was closed")
				return nil
			}

			if msg == nil || len(msg.Value) == 0 {
				continue
			}

			var m = broker.Message{
				Headers: make(map[string]string),
			}

			for _, header := range msg.Headers {
				m.Headers[string(header.Key)] = string(header.Value)
			}

			m.Body = []byte(msg.Value)
			p := &publication{m: &m, t: msg.Topic, km: msg, cg: h.cg, sess: session}

			err := h.handler(p)
			if err == nil && h.subopts.AutoAck {
				session.MarkMessage(msg, "")
			} else if err != nil {
				p.err = err
				errHandler := h.kopts.ErrorHandler
				if errHandler != nil {
					errHandler(p)
				} else {
					h.logger.Errorf(ctx, "[kafka] subscriber error: %v", err)
				}
			}
		case <-session.Context().Done():
			return nil
		}
	}
}
