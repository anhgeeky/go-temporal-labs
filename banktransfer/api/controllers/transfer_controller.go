package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	apiUtils "github.com/anhgeeky/go-temporal-labs/banktransfer/api/utils"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/messages"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/utils"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/workflows"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/modules/transaction"
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

type TransferController struct {
	Service        transaction.Service
	TemporalClient client.Client
}

func (r TransferController) CreateTransferHandler(c *fiber.Ctx) error {
	workflowID := "TRANSFER-" + fmt.Sprintf("%d", time.Now().Unix())

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: utils.Workflows.BANK_TRANSFER,
	}

	msg := messages.TransferState{Items: make([]messages.TransferItem, 0)}
	we, err := r.TemporalClient.ExecuteWorkflow(context.Background(), options, workflows.TransferWorkflow, msg)
	if err != nil {
		return apiUtils.WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["msg"] = msg
	res["workflowID"] = we.GetID()

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (r TransferController) GetTransferHandler(c *fiber.Ctx) error {
	workflowID := c.Params("workflowID")
	response, err := r.TemporalClient.QueryWorkflow(context.Background(), workflowID, "", "getTransfer")
	if err != nil {
		return apiUtils.WriteError(c, err)
	}
	var res interface{}
	if err := response.Get(&res); err != nil {
		return apiUtils.WriteError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (r TransferController) AddToTransferHandler(c *fiber.Ctx) error {
	workflowID := c.Params("workflowID")
	var item messages.TransferItem
	json.Unmarshal(c.Body(), &item)

	update := messages.AddToTransferSignal{Route: utils.RouteTypes.ADD_TO_TRANSFER, Item: item}

	err := r.TemporalClient.SignalWorkflow(context.Background(), workflowID, "", utils.SignalChannels.ADD_TO_TRANSFER_CHANNEL, update)
	if err != nil {
		return apiUtils.WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["ok"] = 1
	return c.Status(fiber.StatusOK).JSON(res)
}

func (r TransferController) RemoveFromTransferHandler(c *fiber.Ctx) error {
	workflowID := c.Params("workflowID")
	var item messages.TransferItem
	json.Unmarshal(c.Body(), &item)

	update := messages.RemoveFromTransferSignal{Route: utils.RouteTypes.REMOVE_FROM_TRANSFER, Item: item}

	err := r.TemporalClient.SignalWorkflow(context.Background(), workflowID, "", utils.SignalChannels.REMOVE_FROM_TRANSFER_CHANNEL, update)
	if err != nil {
		return apiUtils.WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["ok"] = 1
	return c.Status(fiber.StatusOK).JSON(res)
}

func (r TransferController) UpdateEmailHandler(c *fiber.Ctx) error {
	workflowID := c.Params("workflowID")

	var body transaction.UpdateEmailRequest
	json.Unmarshal(c.Body(), &body)
	updateEmail := messages.UpdateEmailSignal{Route: utils.RouteTypes.UPDATE_EMAIL, Email: body.Email}

	err := r.TemporalClient.SignalWorkflow(context.Background(), workflowID, "", utils.SignalChannels.UPDATE_EMAIL_CHANNEL, updateEmail)
	if err != nil {
		return apiUtils.WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["ok"] = 1
	return c.Status(fiber.StatusOK).JSON(res)
}

func (r TransferController) CheckoutHandler(c *fiber.Ctx) error {
	workflowID := c.Params("workflowID")

	var body transaction.CheckoutRequest
	json.Unmarshal(c.Body(), &body)
	checkout := messages.CheckoutSignal{Route: utils.RouteTypes.CHECKOUT, Email: body.Email}

	err := r.TemporalClient.SignalWorkflow(context.Background(), workflowID, "", utils.SignalChannels.CHECKOUT_CHANNEL, checkout)
	if err != nil {
		return apiUtils.WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["sent"] = true
	return c.Status(fiber.StatusOK).JSON(res)
}
