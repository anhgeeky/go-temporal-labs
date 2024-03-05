package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/messages"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/account"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/utils"
	"github.com/anhgeeky/go-temporal-labs/core/broker"
	"github.com/anhgeeky/go-temporal-labs/core/broker/kafka"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func apiCreateTransfer(temporalClient client.Client) (string, error) {
	workflowID := "BANK_TRANSFER-" + fmt.Sprintf("%d", time.Now().Unix())
	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: config.TaskQueues.TRANSFER_QUEUE,
	}

	now := time.Now()

	msg := messages.Transfer{
		Id:                   uuid.NewString(),
		WorkflowID:           workflowID,
		AccountOriginId:      "123", // Test Only
		AccountDestinationId: "456", // Test Only
		CreatedAt:            &now,
	}

	we, err := temporalClient.ExecuteWorkflow(context.Background(), options, "TransferWorkflow", msg)
	if err != nil {
		return "", err
	}

	return we.GetID(), nil
}

func apiVerifyOtp(temporalClient client.Client, workflowID string) error {
	item := messages.VerifyOtpReq{
		FlowId: workflowID,
		Token:  "token",
		Code:   "code",
		Trace:  "trace",
	}

	update := messages.VerifiedOtpSignal{Route: config.RouteTypes.VERIFY_OTP, Item: item}

	// Trigger Signal Transfer Flow
	err := temporalClient.SignalWorkflow(context.Background(), item.FlowId, "", config.SignalChannels.VERIFY_OTP_CHANNEL, update)
	if err != nil {
		return err
	}

	return nil
}

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

	temporalClient, err := client.NewLazyClient(client.Options{
		HostPort:  "localhost:7233",
		Namespace: "staging",
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	log.Println("Temporal client connected")

	// workflowID := "BANK_TRANSFER-1709632114"

	// 1. Tạo lệnh chuyển tiền
	workflowID, err := apiCreateTransfer(temporalClient)
	if err != nil {
		log.Fatalln("error apiCreateTransfer", err)
	}

	// 2. Xác thực OTP
	err = apiVerifyOtp(temporalClient, workflowID)
	if err != nil {
		log.Fatalln("error apiCreateTransfer", err)
	}

	// 3. Nhận message check balance từ Temporal
	go func() {
		if err := runCheckBalance(bk, workflowID); err != nil {
			errChan <- err
			cancel()
		}
	}()

	// 4. Nhận message create transaction từ Temporal
	go func() {
		if err := runCreateTransferTransaction(bk, workflowID); err != nil {
			errChan <- err
			cancel()
		}
	}()

	// 5. Done 2 activity + 1 activity notification

	select {}
}
