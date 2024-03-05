package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/account"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/utils"
	"github.com/anhgeeky/go-temporal-labs/core/broker"
	"github.com/anhgeeky/go-temporal-labs/core/broker/kafka"
)

func runCheckBalance(bk broker.Broker, workflowID string) error {
	requestTopic := config.Messages.CHECK_BALANCE_REQUEST_TOPIC
	replyTopic := config.Messages.CHECK_BALANCE_REPLY_TOPIC
	action := config.Messages.CHECK_BALANCE_ACTION

	// Gen consumer group theo format
	csGroupOpt := broker.WithSubscribeGroup(utils.GetConsumerGroup(workflowID, action))

	bk.Subscribe(requestTopic, func(e broker.Event) error {
		headers := e.Message().Headers
		fmt.Printf("Received message from request topic %v: Header: %v\n", requestTopic, headers)
		// TODO: Nhận response từ API Microservice push vào topic Reply

		// ======================== REPLY: SEND REQUEST ========================
		req := broker.Response[account.CheckBalanceRes]{
			Result: broker.Result{
				Status: 200, // OK
			},
			Data: account.CheckBalanceRes{Balance: 8888},
		}
		body, err := json.Marshal(req)
		if err != nil {
			panic(err)
		}

		fMsg := broker.Message{
			Body: body,
			Headers: map[string]string{
				"workflow_id":   workflowID,
				"activity-id":   action,
				"correlationId": headers["correlationId"],
			},
		}

		fmt.Printf("Reply message to reply topic %v: Header: %v\n", replyTopic, headers)
		bk.Publish(replyTopic, &fMsg)
		// ======================== REPLY: SEND REQUEST ========================

		return nil
	}, csGroupOpt)

	return nil
}

func runCreateTransferTransaction(bk broker.Broker, workflowID string) error {
	requestTopic := config.Messages.CREATE_TRANSACTION_REQUEST_TOPIC
	replyTopic := config.Messages.CREATE_TRANSACTION_REPLY_TOPIC
	action := config.Messages.CREATE_TRANSACTION_ACTION

	// Gen consumer group theo format
	csGroupOpt := broker.WithSubscribeGroup(utils.GetConsumerGroup(workflowID, action))

	bk.Subscribe(requestTopic, func(e broker.Event) error {
		headers := e.Message().Headers
		fmt.Printf("Received message from request topic %v: Header: %v\n", requestTopic, headers)
		// TODO: Nhận response từ API Microservice push vào topic Reply

		// ======================== REPLY: SEND REQUEST ========================
		req := broker.Response[account.CreateTransactionRes]{
			Result: broker.Result{
				Status: 200, // OK
			},
			Data: account.CreateTransactionRes{},
		}
		body, err := json.Marshal(req)
		if err != nil {
			panic(err)
		}

		fMsg := broker.Message{
			Body: body,
			Headers: map[string]string{
				"workflow_id":   workflowID,
				"activity-id":   action,
				"correlationId": headers["correlationId"],
			},
		}

		fmt.Printf("Reply message to reply topic %v: Header: %v\n", replyTopic, headers)
		bk.Publish(replyTopic, &fMsg)
		// ======================== REPLY: SEND REQUEST ========================

		return nil
	}, csGroupOpt)

	return nil
}

// Micro: Nhận request từ Temporal -> Reply lại Temporal
func main() {
	_, cancel := context.WithCancel(context.Background())
	errChan := make(chan error)
	// ======================= BROKER =======================
	bk := kafka.ConnectBrokerKafka("127.0.0.1:9092")
	// ======================= BROKER =======================
	workflowID := "BANK_TRANSFER-1709632114"

	go func() {
		if err := runCheckBalance(bk, workflowID); err != nil {
			errChan <- err
			cancel()
		}
	}()

	go func() {
		if err := runCreateTransferTransaction(bk, workflowID); err != nil {
			errChan <- err
			cancel()
		}
	}()

	select {}
}
