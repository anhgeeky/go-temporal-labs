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

func (r TransferController) CreateTransfer(c *fiber.Ctx) error {
	workflowID := "TRANSFER-" + fmt.Sprintf("%d", time.Now().Unix())
	var req messages.TransferReq
	json.Unmarshal(c.Body(), &req)

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: configs.Workflows.BANK_TRANSFER,
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

func (r TransferController) AddToTransfer(c *fiber.Ctx) error {
	workflowID := c.Params("workflowID")
	var item messages.TransferItem
	json.Unmarshal(c.Body(), &item)

	update := messages.AddToTransferSignal{Route: configs.RouteTypes.ADD_TO_TRANSFER, Item: item}

	err := r.TemporalClient.SignalWorkflow(context.Background(), workflowID, "", configs.SignalChannels.ADD_TO_TRANSFER_CHANNEL, update)
	if err != nil {
		return apis.WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["ok"] = 1
	return c.Status(fiber.StatusOK).JSON(res)
}

func (r TransferController) RemoveFromTransfer(c *fiber.Ctx) error {
	workflowID := c.Params("workflowID")
	var item messages.TransferItem
	json.Unmarshal(c.Body(), &item)

	update := messages.RemoveFromTransferSignal{Route: configs.RouteTypes.REMOVE_FROM_TRANSFER, Item: item}

	err := r.TemporalClient.SignalWorkflow(context.Background(), workflowID, "", configs.SignalChannels.REMOVE_FROM_TRANSFER_CHANNEL, update)
	if err != nil {
		return apis.WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["ok"] = 1
	return c.Status(fiber.StatusOK).JSON(res)
}

func (r TransferController) UpdateEmail(c *fiber.Ctx) error {
	workflowID := c.Params("workflowID")

	var body transaction.UpdateEmailRequest
	json.Unmarshal(c.Body(), &body)
	updateEmail := messages.UpdateEmailSignal{Route: configs.RouteTypes.UPDATE_EMAIL, Email: body.Email}

	err := r.TemporalClient.SignalWorkflow(context.Background(), workflowID, "", configs.SignalChannels.UPDATE_EMAIL_CHANNEL, updateEmail)
	if err != nil {
		return apis.WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["ok"] = 1
	return c.Status(fiber.StatusOK).JSON(res)
}

func (r TransferController) Checkout(c *fiber.Ctx) error {
	workflowID := c.Params("workflowID")

	var body transaction.CheckoutRequest
	json.Unmarshal(c.Body(), &body)
	checkout := messages.CheckoutSignal{Route: configs.RouteTypes.CHECKOUT, Email: body.Email}

	err := r.TemporalClient.SignalWorkflow(context.Background(), workflowID, "", configs.SignalChannels.CHECKOUT_CHANNEL, checkout)
	if err != nil {
		return apis.WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["sent"] = true
	return c.Status(fiber.StatusOK).JSON(res)
}
