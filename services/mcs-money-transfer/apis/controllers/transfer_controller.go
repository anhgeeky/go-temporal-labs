package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/messages"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/workflows"

	"github.com/anhgeeky/go-temporal-labs/core/apis/responses"
	"github.com/anhgeeky/go-temporal-labs/mcs-money-transfer/modules/transaction"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

type TransferController struct {
	Service        transaction.Service
	TemporalClient client.Client
}

// Done
func (r TransferController) CreateTransfer(c *fiber.Ctx) error {
	workflowID := "BANK_TRANSFER-" + fmt.Sprintf("%d", time.Now().Unix())
	var req messages.TransferReq
	json.Unmarshal(c.Body(), &req)

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: config.TaskQueues.TRANSFER_QUEUE,
	}

	now := time.Now()

	msg := messages.Transfer{
		Id:                   uuid.NewString(),
		WorkflowID:           workflowID,
		AccountOriginId:      req.AccountOriginId,
		AccountDestinationId: req.AccountDestinationId,
		CreatedAt:            &now,
	}

	we, err := r.TemporalClient.ExecuteWorkflow(context.Background(), options, workflows.TransferWorkflow, msg)
	if err != nil {
		return responses.WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["msg"] = msg
	res["workflowID"] = we.GetID()

	return responses.SuccessResult[interface{}](c, res)
}

func (r TransferController) GetTransfer(c *fiber.Ctx) error {
	workflowID := c.Params("workflowID")
	response, err := r.TemporalClient.QueryWorkflow(context.Background(), workflowID, "", "getTransfer")
	if err != nil {
		return responses.WriteError(c, err)
	}
	var res interface{}
	if err := response.Get(&res); err != nil {
		return responses.WriteError(c, err)
	}

	return responses.SuccessResult[interface{}](c, res)
}

func (r TransferController) CreateTransferTransaction(c *fiber.Ctx) error {
	return responses.SuccessResult(c, transaction.SampleRes{Msg: "OK"})
}

func (r TransferController) WriteCreditAccount(c *fiber.Ctx) error {
	return responses.SuccessResult(c, transaction.SampleRes{Msg: "OK"})
}

func (r TransferController) WriteDebitAccount(c *fiber.Ctx) error {
	// TODO: Đang test case lỗi error
	return responses.WriteError(c, errors.New("OOPS!!! WriteDebitAccount error"))
	// return responses.SuccessResult(c, transaction.SampleRes{Msg: "OK"})
}

func (r TransferController) AddNewActivity(c *fiber.Ctx) error {
	// TODO: Đang test case lỗi error
	return responses.WriteError(c, errors.New("OOPS!!! AddNewActivity error"))
	// return responses.SuccessResult(c, transaction.SampleRes{Msg: "OK"})
}

// ============================================
// Rollback
// ============================================

func (r TransferController) CreateTransferTransactionCompensation(c *fiber.Ctx) error {
	return responses.SuccessResult(c, transaction.SampleRes{Msg: "Rollback done"})
}

func (r TransferController) WriteCreditAccountCompensation(c *fiber.Ctx) error {
	return responses.SuccessResult(c, transaction.SampleRes{Msg: "Rollback done"})
}

func (r TransferController) WriteDebitAccountCompensation(c *fiber.Ctx) error {
	return responses.SuccessResult(c, transaction.SampleRes{Msg: "Rollback done"})
}

func (r TransferController) AddNewActivityCompensation(c *fiber.Ctx) error {
	return responses.SuccessResult(c, transaction.SampleRes{Msg: "Rollback done"})
}
