package activities

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/messages"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/account"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/moneytransfer"
	"github.com/anhgeeky/go-temporal-labs/core/broker"
	"go.temporal.io/sdk/activity"
)

type TransferActivity struct {
	Broker               broker.Broker
	AccountService       account.AccountService
	MoneyTransferService moneytransfer.MoneyTransferService
}

func (a *TransferActivity) CheckBalance(ctx context.Context, msg messages.Transfer) (*account.CheckBalanceRes, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: CheckBalance", msg)
	requestTopic := config.Messages.CHECK_BALANCE_REQUEST_TOPIC
	replyTopic := config.Messages.CHECK_BALANCE_REPLY_TOPIC
	action := config.Messages.CHECK_BALANCE_ACTION
	// ======================== TEST ONLY ========================
	// Call REST api
	// res, err := a.AccountService.GetBalance()
	// if err != nil {
	// 	return err
	// }
	// ======================== TEST ONLY ========================

	// ======================== REQUEST: SEND REQUEST ========================
	req := account.CheckBalanceReq{}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	fMsg := broker.Message{
		Body: body,
		Headers: map[string]string{
			"workflow_id": msg.WorkflowID,
			"activity-id": action,
		},
	}

	a.Broker.Publish(requestTopic, &fMsg)
	// ======================== REQUEST: SEND REQUEST ========================

	// ======================== REPLY: GET RESPONSE ========================
	isReceived := false
	var res account.CheckBalanceRes
	// csGroupOpt := broker.WithSubscribeGroup(config.Messages.GROUP)

	// Loop -> khi nào có message phù hợp -> Nhận + parse message -> Done activity
	// TODO: Trường hợp không tìm thấy được message phù hợp -> Timeout
	for {
		a.Broker.Subscribe(replyTopic, func(e broker.Event) error {
			headers := e.Message().Headers
			fmt.Printf("Received message from topic %v: Header: %v\n", replyTopic, headers)
			// TODO: Nhận response từ API Microservice push vào topic Reply

			// Kiểm tra theo điều kiện phù hợp
			if headers["workflow_id"] == msg.WorkflowID && headers["activity-id"] == action { // TODO: check lại với Sơn
				body := string(e.Message().Body)
				if body != "" {
					err := json.Unmarshal(e.Message().Body, &res)
					if err != nil {
						return err // Đúng message + Payload res bị sai struct -> Fail Activity
					} else {
						isReceived = true
					}
				}
			}

			return nil
		}) //, csGroupOpt)

		if isReceived {
			break
		}
	}

	if isReceived {
		logger.Info("TransferActivity: CheckBalance done", res)
	}
	// ======================== REPLY: GET RESPONSE ========================

	return &res, nil
}

// func (a *TransferActivity) CheckTargetAccount(ctx context.Context, msg messages.Transfer) error {
// 	logger := activity.GetLogger(ctx)
// 	logger.Info("TransferActivity: CheckTargetAccount", msg)
// 	return nil
// }

func (a *TransferActivity) CreateTransferTransaction(ctx context.Context, msg messages.Transfer) (*account.CreateTransactionRes, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("TransferActivity: CreateTransferTransaction", msg)

	// ======================== TEST ONLY ========================
	// res, err := a.MoneyTransferService.CreateTransferTransaction(msg.WorkflowID)
	// if err != nil {
	// 	logger.Error("TransferActivity CreateTransferTransaction failed.", "Error", err)
	// 	return err
	// }
	// ======================== TEST ONLY ========================

	requestTopic := config.Messages.CREATE_TRANSACTION_REQUEST_TOPIC
	replyTopic := config.Messages.CREATE_TRANSACTION_REPLY_TOPIC
	action := config.Messages.CREATE_TRANSACTION_ACTION

	// ======================== REQUEST: SEND REQUEST ========================
	req := account.CreateTransactionReq{}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	fMsg := broker.Message{
		Body: body,
		Headers: map[string]string{
			"workflow_id": msg.WorkflowID,
			"activity-id": action,
		},
	}

	a.Broker.Publish(requestTopic, &fMsg)
	// ======================== REQUEST: SEND REQUEST ========================

	// ======================== REPLY: GET RESPONSE ========================

	isReceived := false
	var res account.CreateTransactionRes // TODO: check lại với Sơn
	// csGroupOpt := broker.WithSubscribeGroup(config.Messages.GROUP)

	// Loop -> khi nào có message phù hợp -> Nhận + parse message -> Done activity
	// TODO: Trường hợp không tìm thấy được message phù hợp -> Timeout
	for {
		a.Broker.Subscribe(replyTopic, func(e broker.Event) error {
			headers := e.Message().Headers
			fmt.Printf("Received message from topic %v: Header: %v\n", replyTopic, headers)
			// TODO: Nhận response từ API Microservice push vào topic Reply

			// Kiểm tra theo điều kiện phù hợp
			if headers["workflow_id"] == msg.WorkflowID && headers["activity-id"] == action { // TODO: check lại với Sơn
				body := string(e.Message().Body)
				if body != "" {
					err := json.Unmarshal(e.Message().Body, &res)
					if err != nil {
						return err // Đúng message + Payload res bị sai struct -> Fail Activity
					} else {
						isReceived = true
					}
				}
			}

			return nil
		}) //, csGroupOpt)
		// }, csGroupOpt)

		if isReceived {
			break
		}
	}

	if isReceived {
		logger.Info("TransferActivity: CheckBalance done", res)
	}
	// ======================== REPLY: GET RESPONSE ========================

	logger.Info("TransferActivity: CreateTransferTransaction done", res)

	return &res, nil
}

// func (a *TransferActivity) WriteCreditAccount(ctx context.Context, msg messages.Transfer) error {
// 	logger := activity.GetLogger(ctx)
// 	// time.Sleep(time.Duration(30) * time.Minute) // TODO: Test chờ 30p
// 	logger.Info("TransferActivity: WriteCreditAccount", msg)
// 	res, err := a.MoneyTransferService.WriteCreditAccount(msg.WorkflowID)
// 	if err != nil {
// 		logger.Error("TransferActivity WriteCreditAccount failed.", "Error", err)
// 		return err
// 	}
// 	logger.Info("TransferActivity: WriteCreditAccount done", res)
// 	return nil
// }

// func (a *TransferActivity) WriteDebitAccount(ctx context.Context, msg messages.Transfer) error {
// 	logger := activity.GetLogger(ctx)
// 	// time.Sleep(time.Duration(30) * time.Minute) // TODO: Test chờ 30p
// 	logger.Info("TransferActivity: WriteDebitAccount", msg)
// 	res, err := a.MoneyTransferService.WriteDebitAccount(msg.WorkflowID)
// 	if err != nil {
// 		logger.Error("TransferActivity WriteDebitAccount failed.", "Error", err)
// 		return err
// 	}
// 	logger.Info("TransferActivity: WriteDebitAccount done", res)
// 	return nil
// }

// func (a *TransferActivity) AddNewActivity(ctx context.Context, msg messages.Transfer) error {
// 	logger := activity.GetLogger(ctx)
// 	// time.Sleep(time.Duration(30) * time.Minute) // TODO: Test chờ 30p
// 	logger.Info("TransferActivity: AddNewActivity", msg)
// 	res, err := a.MoneyTransferService.AddNewActivity(msg.WorkflowID)
// 	if err != nil {
// 		logger.Error("TransferActivity AddNewActivity failed.", "Error", err)
// 		return err
// 	}
// 	logger.Info("TransferActivity: AddNewActivity done", res)
// 	return nil
// }

// ============================================
// Rollback
// ============================================

// func (a *TransferActivity) CreateTransferTransactionCompensation(ctx context.Context, msg messages.Transfer) error {
// 	logger := activity.GetLogger(ctx)
// 	logger.Info("TransferActivity: CreateTransferTransaction", msg)
// 	res, err := a.MoneyTransferService.CreateTransferTransaction(msg.WorkflowID)
// 	if err != nil {
// 		logger.Error("TransferActivity CreateTransferTransactionCompensation failed.", "Error", err)
// 		return err
// 	}
// 	logger.Info("TransferActivity: CreateTransferTransaction done", res)
// 	return nil
// }

// func (a *TransferActivity) WriteCreditAccountCompensation(ctx context.Context, msg messages.Transfer) error {
// 	logger := activity.GetLogger(ctx)
// 	logger.Info("TransferActivity: WriteCreditAccountCompensation", msg)
// 	res, err := a.MoneyTransferService.WriteCreditAccountCompensation(msg.WorkflowID)
// 	if err != nil {
// 		logger.Error("TransferActivity WriteCreditAccountCompensation failed.", "Error", err)
// 		return err
// 	}
// 	logger.Info("TransferActivity: WriteCreditAccountCompensation done", res)
// 	return nil
// }

// func (a *TransferActivity) WriteDebitAccountCompensation(ctx context.Context, msg messages.Transfer) error {
// 	logger := activity.GetLogger(ctx)
// 	logger.Info("TransferActivity: WriteDebitAccountCompensation", msg)
// 	res, err := a.MoneyTransferService.WriteDebitAccountCompensation(msg.WorkflowID)
// 	if err != nil {
// 		logger.Error("TransferActivity WriteDebitAccountCompensation failed.", "Error", err)
// 		return err
// 	}
// 	logger.Info("TransferActivity: WriteDebitAccountCompensation done", res)
// 	return nil
// }

// func (a *TransferActivity) AddNewActivityCompensation(ctx context.Context, msg messages.Transfer) error {
// 	logger := activity.GetLogger(ctx)
// 	logger.Info("TransferActivity: AddNewActivityCompensation", msg)
// 	res, err := a.MoneyTransferService.AddNewActivityCompensation(msg.WorkflowID)
// 	if err != nil {
// 		logger.Error("TransferActivity AddNewActivityCompensation failed.", "Error", err)
// 		return err
// 	}
// 	logger.Info("TransferActivity: AddNewActivityCompensation done", res)
// 	return nil
// }
