package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/messages"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/workflows"

	"github.com/anhgeeky/go-temporal-labs/core/apis/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

type TransferController struct {
	TemporalClient client.Client
}

// 1. Tạo lệnh chuyển tiền
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

// 6. Trả về kết quả Tạo giao dịch thành công
func (r TransferController) CreateTransaction(c *fiber.Ctx) error {
	var item messages.CreateTransactionReq
	json.Unmarshal(c.Body(), &item)

	update := messages.CreateTransactionSignal{Item: item}

	// Trigger Signal Transfer Flow
	err := r.TemporalClient.SignalWorkflow(context.Background(), item.FlowId, "", "CREATE_TRANSACTION_CHANNEL", update)
	if err != nil {
		return responses.WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["ok"] = 1
	return c.Status(fiber.StatusOK).JSON(res)
}
