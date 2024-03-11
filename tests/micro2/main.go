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

// Signal: Xác thực OTP thành công
func apiSignalVerifyOtp(temporalClient client.Client, workflowID string) error {
	item := messages.VerifyOtpReq{
		FlowId: workflowID,
		Token:  "token",
		Code:   "code",
		Trace:  "trace",
	}

	update := messages.VerifiedOtpSignal{Item: item}

	// Trigger Signal
	err := temporalClient.SignalWorkflow(context.Background(), item.FlowId, "", "VERIFY_OTP_CHANNEL", update)
	if err != nil {
		return err
	}

	return nil
}

// Signal: Trả về kết quả Tạo giao dịch thành công
func apiSignalCreateTransaction(temporalClient client.Client, workflowID string) error {
	item := messages.CreateTransactionReq{
		FlowId: workflowID,
		// TODO: Sơn bổ sung Data Response giúp anh -> Gửi email ra
	}

	update := messages.CreateTransactionSignal{Item: item}

	// Trigger Signal
	err := temporalClient.SignalWorkflow(context.Background(), item.FlowId, "", "CREATE_TRANSACTION_CHANNEL", update)
	if err != nil {
		return err
	}

	return nil
}

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

	temporalClient, err := client.NewLazyClient(client.Options{
		HostPort:  "localhost:7233",
		Namespace: "staging",
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	log.Println("Temporal client connected")

	// err = temporalClient.UpdateWorkerBuildIdCompatibility(ctx, &client.UpdateWorkerBuildIdCompatibilityOptions{
	// 	TaskQueue: config.TaskQueues.TRANSFER_QUEUE,
	// 	Operation: &client.BuildIDOpAddNewIDInNewDefaultSet{
	// 		BuildID: config.VERSION_1_0,
	// 	},
	// })

	taskQueue := config.TaskQueues.TRANSFER_QUEUE
	beforeVersion := config.VERSION_2_0
	latestVersion := config.VERSION_3_0

	updateLatestWorkerBuildId(temporalClient, taskQueue, beforeVersion, latestVersion)

	// 1. Tạo lệnh chuyển tiền
	workflowID, err := apiCreateTransfer(temporalClient)
	if err != nil {
		log.Fatalln("error apiCreateTransfer", err)
	}

	// Nhận message check balance từ Temporal
	go func() {
		if err := runCheckBalance(bk, workflowID); err != nil {
			errChan <- err
			cancel()
		}
	}()

	// Nhận message create otp từ Temporal
	go func() {
		if err := runCreateOTP(bk, workflowID); err != nil {
			errChan <- err
			cancel()
		}
	}()

	// Nhận message create transaction từ Temporal
	go func() {
		if err := runCreateTransaction(bk, workflowID); err != nil {
			errChan <- err
			cancel()
		}
	}()

	time.Sleep(3 * time.Second)

	// 2. Xác thực OTP
	err = apiSignalVerifyOtp(temporalClient, workflowID)
	if err != nil {
		log.Fatalln("error apiCreateTransfer", err)
	}

	time.Sleep(3 * time.Second)

	// 3. Trả kết quả đã tạo giao dịch
	err = apiSignalCreateTransaction(temporalClient, workflowID)
	if err != nil {
		log.Fatalln("error apiSignalCreateTransaction", err)
	}

	select {}
}

// FAQ: https://docs.temporal.io/dev-guide/go/versioning
func updateLatestWorkerBuildId(c client.Client, taskQueue, compatibleBuildID, latestBuildID string) {
	ctx := context.Background()
	c.UpdateWorkerBuildIdCompatibility(ctx, &client.UpdateWorkerBuildIdCompatibilityOptions{
		TaskQueue: taskQueue,
		Operation: &client.BuildIDOpAddNewCompatibleVersion{
			BuildID:                   latestBuildID,
			ExistingCompatibleBuildID: compatibleBuildID,
		},
	})
}
