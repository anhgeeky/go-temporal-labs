package activities

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/messages"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/account"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/moneytransfer"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/utils"
	"github.com/anhgeeky/go-temporal-labs/core/broker"
	"go.temporal.io/sdk/activity"
)

type TransferActivity struct {
	Broker               broker.Broker
	AccountService       account.AccountService
	MoneyTransferService moneytransfer.MoneyTransferService
}

var (
	workflowIDKey = "workflow_id"
	activityIDKey = "activity-id"
)

func (a *TransferActivity) getMsgHeaders(workflowId, activityId string) map[string]string {
	return map[string]string{
		workflowIDKey: workflowId,
		activityIDKey: activityId,
	}
}

func (a *TransferActivity) checkMsgHeaders(headers map[string]string, workflowId, activityId string) bool {
	return headers[workflowIDKey] == workflowId && headers[activityIDKey] == activityId
}

func (a *TransferActivity) CheckBalance(ctx context.Context, msg messages.Transfer) (*account.CheckBalanceRes, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: CheckBalance", msg)

	requestTopic := config.Messages.CHECK_BALANCE_REQUEST_TOPIC
	replyTopic := config.Messages.CHECK_BALANCE_REPLY_TOPIC
	action := config.Messages.CHECK_BALANCE_ACTION

	req := account.CheckBalanceReq{}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	msgRes, err := a.Broker.PublishAndReceive(
		requestTopic,
		&broker.Message{
			Headers: a.getMsgHeaders(msg.WorkflowID, action),
			Body:    reqBody,
		},
		broker.WithPublishReplyToTopic(replyTopic),
		broker.WithReplyConsumerGroup(utils.GetConsumerGroup(msg.WorkflowID, action)),
	)

	if err != nil {
		return nil, err
	}

	var res *broker.Response[account.CheckBalanceRes] // TODO: check lại với Sơn

	// Kiểm tra theo điều kiện phù hợp theo Headers -> TODO: Xem lại có sài theo Headers không?
	if a.checkMsgHeaders(msgRes.Headers, msg.WorkflowID, action) && len(msgRes.Body) > 0 {
		err := json.Unmarshal(msgRes.Body, &res)
		if err != nil {
			return nil, err // Đúng message + Payload res bị sai struct -> Fail Activity
		}
		if res == nil || res.Result.Status != 200 {
			// Kết quả Status <> 200 -> Return failure activity
			return nil, errors.New("Error: Invalid data result from Kafka")
		}
	}

	logger.Info("TransferActivity: CheckBalance done", res)

	return &res.Data, nil
}

func (a *TransferActivity) CreateOTP(ctx context.Context, msg messages.Transfer) (*account.CreateOTPRes, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: CreateOTP", msg)

	requestTopic := config.Messages.CREATE_OTP_REQUEST_TOPIC
	replyTopic := config.Messages.CREATE_OTP_REPLY_TOPIC
	action := config.Messages.CREATE_OTP_ACTION

	req := account.CheckBalanceReq{}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	msgRes, err := a.Broker.PublishAndReceive(
		requestTopic,
		&broker.Message{
			Headers: a.getMsgHeaders(msg.WorkflowID, action),
			Body:    reqBody,
		},
		broker.WithPublishReplyToTopic(replyTopic),
		broker.WithReplyConsumerGroup(utils.GetConsumerGroup(msg.WorkflowID, action)),
	)

	if err != nil {
		return nil, err
	}

	var res *broker.Response[account.CreateOTPRes] // TODO: check lại với Sơn

	// Kiểm tra theo điều kiện phù hợp theo Headers -> TODO: Xem lại có sài theo Headers không?
	if a.checkMsgHeaders(msgRes.Headers, msg.WorkflowID, action) && len(msgRes.Body) > 0 {
		err := json.Unmarshal(msgRes.Body, &res)
		if err != nil {
			return nil, err // Đúng message + Payload res bị sai struct -> Fail Activity
		}
		if res == nil || res.Result.Status != 200 {
			// Kết quả Status <> 200 -> Return failure activity
			return nil, errors.New("Error: Invalid data result from Kafka")
		}
	}

	logger.Info("TransferActivity: CreateOTP done", res)

	return &res.Data, nil
}

func (a *TransferActivity) CreateTransferTransaction(ctx context.Context, msg messages.Transfer) (*account.CreateTransactionRes, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: CreateTransferTransaction", msg)

	requestTopic := config.Messages.CREATE_TRANSACTION_REQUEST_TOPIC
	replyTopic := config.Messages.CREATE_TRANSACTION_REPLY_TOPIC
	action := config.Messages.CREATE_TRANSACTION_ACTION

	req := account.CreateTransactionReq{}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	msgRes, err := a.Broker.PublishAndReceive(
		requestTopic,
		&broker.Message{
			Headers: a.getMsgHeaders(msg.WorkflowID, action),
			Body:    reqBody,
		},
		broker.WithPublishReplyToTopic(replyTopic),
		broker.WithReplyConsumerGroup(utils.GetConsumerGroup(msg.WorkflowID, action)),
	)

	if err != nil {
		return nil, err
	}

	var res *broker.Response[account.CreateTransactionRes] // TODO: check lại với Sơn

	// Kiểm tra theo điều kiện phù hợp theo Headers -> TODO: Xem lại có sài theo Headers không?
	if a.checkMsgHeaders(msgRes.Headers, msg.WorkflowID, action) && len(msgRes.Body) > 0 {
		err := json.Unmarshal(msgRes.Body, &res)
		if err != nil {
			return nil, err // Đúng message + Payload res bị sai struct -> Fail Activity
		}
		if res == nil || res.Result.Status != 200 {
			// Kết quả Status <> 200 -> Return failure activity
			return nil, errors.New("Error: Invalid data result from Kafka")
		}
	}

	logger.Info("TransferActivity: CreateTransferTransaction done", res)

	return &res.Data, nil
}
