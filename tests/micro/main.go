package main

import (
	"encoding/json"
	"fmt"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/account"
	"github.com/anhgeeky/go-temporal-labs/core/broker"
	"github.com/anhgeeky/go-temporal-labs/core/broker/kafka"
)

// Micro: Nhận request từ Temporal -> Reply lại Temporal
func main() {
	// ======================= BROKER =======================
	bk := kafka.ConnectBrokerKafka("127.0.0.1:9092")
	// ======================= BROKER =======================

	workflowID := "BANK_TRANSFER-1709525114"
	requestTopic := config.Messages.CHECK_BALANCE_REQUEST_TOPIC
	replyTopic := config.Messages.CHECK_BALANCE_REPLY_TOPIC

	csGroupOpt := broker.WithSubscribeGroup(config.Messages.GROUP)

	bk.Subscribe(requestTopic, func(e broker.Event) error {
		headers := e.Message().Headers
		fmt.Printf("Received message from request topic %v: Header: %v\n", requestTopic, headers)
		// TODO: Nhận response từ API Microservice push vào topic Reply

		// ======================== REPLY: SEND REQUEST ========================
		req := account.CheckBalanceRes{Balance: 8888}
		body, err := json.Marshal(req)
		if err != nil {
			panic(err)
		}

		fMsg := broker.Message{
			Body: body,
			Headers: map[string]string{
				"workflow_id":   workflowID,
				"activity-id":   config.Messages.CHECK_BALANCE_ACTION,
				"correlationId": headers["correlationId"],
			},
		}

		fmt.Printf("Reply message to reply topic %v: Header: %v\n", replyTopic, headers)
		bk.Publish(replyTopic, &fMsg)
		// ======================== REPLY: SEND REQUEST ========================

		return nil
	}, csGroupOpt)

	select {}

}
