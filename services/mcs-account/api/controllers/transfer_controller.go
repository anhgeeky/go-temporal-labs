package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/configs"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/messages"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/workflows"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/modules/transaction"
	"github.com/anhgeeky/go-temporal-labs/core/apis"
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
		TaskQueue: configs.TaskQueues.BANK_TRANSFER,
	}

	now := time.Now()

	msg := messages.Transfer{
		Id:                   uuid.NewString(),
		AccountOriginId:      req.AccountOriginId,
		AccountDestinationId: req.AccountDestinationId,
		CreatedAt:            &now,
	}

	we, err := r.TemporalClient.ExecuteWorkflow(context.Background(), options, workflows.TransferWorkflow, msg)
	if err != nil {
		return apis.WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["msg"] = msg
	res["workflowID"] = we.GetID()

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (r TransferController) GetTransfer(c *fiber.Ctx) error {
	workflowID := c.Params("workflowID")
	response, err := r.TemporalClient.QueryWorkflow(context.Background(), workflowID, "", "getTransfer")
	if err != nil {
		return apis.WriteError(c, err)
	}
	var res interface{}
	if err := response.Get(&res); err != nil {
		return apis.WriteError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
