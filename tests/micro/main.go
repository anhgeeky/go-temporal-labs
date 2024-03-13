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

	csGroupOpt := broker.WithSubscribeGroup(utils.GetConsumerGroup(workflowID, action))

	bk.Subscribe(requestTopic, func(e broker.Event) error {
		headers := e.Message().Headers
		fmt.Printf("Received message from request topic %v: Header: %v\n", requestTopic, headers)

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

func runCreateTransaction(bk broker.Broker, workflowID string) error {
	requestTopic := config.Messages.CREATE_TRANSACTION_REQUEST_TOPIC
	replyTopic := config.Messages.CREATE_TRANSACTION_REPLY_TOPIC
	action := config.Messages.CREATE_TRANSACTION_ACTION

	csGroupOpt := broker.WithSubscribeGroup(utils.GetConsumerGroup(workflowID, action))

	bk.Subscribe(requestTopic, func(e broker.Event) error {
		headers := e.Message().Headers
		fmt.Printf("Received message from request topic %v: Header: %v\n", requestTopic, headers)

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

func runCreateOTP(bk broker.Broker, workflowID string) error {
	requestTopic := config.Messages.CREATE_OTP_REQUEST_TOPIC
	replyTopic := config.Messages.CREATE_OTP_REPLY_TOPIC
	action := config.Messages.CREATE_OTP_ACTION

	csGroupOpt := broker.WithSubscribeGroup(utils.GetConsumerGroup(workflowID, action))

	bk.Subscribe(requestTopic, func(e broker.Event) error {
		headers := e.Message().Headers
		fmt.Printf("Received message from request topic %v: Header: %v\n", requestTopic, headers)

		// ======================== REPLY: SEND REQUEST ========================
		req := broker.Response[account.CreateOTPRes]{
			Result: broker.Result{
				Status: 200, // OK
			},
			Data: account.CreateOTPRes{},
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
	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)
	errChan := make(chan error)
	bk := kafka.ConnectBrokerKafka("127.0.0.1:9092")

	workflowID := "BANK_TRANSFER-1710141239"

	// 1. Nhận message check balance từ Temporal
	go func() {
		if err := runCheckBalance(bk, workflowID); err != nil {
			errChan <- err
			cancel()
		}
	}()

	// 2. Nhận message create otp từ Temporal
	go func() {
		if err := runCreateOTP(bk, workflowID); err != nil {
			errChan <- err
			cancel()
		}
	}()

	// 3. Nhận message create transaction từ Temporal
	go func() {
		if err := runCreateTransaction(bk, workflowID); err != nil {
			errChan <- err
			cancel()
		}
	}()

	select {}
}
