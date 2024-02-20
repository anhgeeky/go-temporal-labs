package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/configs"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/messages"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/workflows"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/modules/otp"
	"github.com/anhgeeky/go-temporal-labs/core/apis"
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

type VerifyController struct {
	Service        otp.Service
	TemporalClient client.Client
}

func (r VerifyController) VerifyOtp(c *fiber.Ctx) error {
	workflowID := "VERIFY-" + fmt.Sprintf("%d", time.Now().Unix())
	var req messages.VerifyOtpReq
	json.Unmarshal(c.Body(), &req)

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: configs.Workflows.VERIFY,
	}

	msg := messages.VerifyOtpMessage{
		Token:   req.Token,
		Code:    req.Code,
		Payload: req.Payload, // TODO: Mã hóa
	}

	we, err := r.TemporalClient.ExecuteWorkflow(context.Background(), options, workflows.VerifyOtpWorkflow, msg)
	if err != nil {
		return apis.WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["msg"] = msg
	res["workflowID"] = we.GetID()

	return c.Status(fiber.StatusCreated).JSON(res)
}
