package controllers

import (
	"context"
	"encoding/json"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/messages"
	"github.com/anhgeeky/go-temporal-labs/core/apis/responses"
	"github.com/anhgeeky/go-temporal-labs/mcs-account/modules/otp"
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

type VerifyController struct {
	Service        otp.Service
	TemporalClient client.Client
}

func (r VerifyController) VerifyOtp(c *fiber.Ctx) error {
	var item messages.VerifyOtpReq
	json.Unmarshal(c.Body(), &item)

	update := messages.VerifiedOtpSignal{Route: config.RouteTypes.VERIFY_OTP, Item: item}

	// Trigger Signal Transfer Flow
	err := r.TemporalClient.SignalWorkflow(context.Background(), item.FlowId, "", config.SignalChannels.VERIFY_OTP_CHANNEL, update)
	if err != nil {
		return responses.WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["ok"] = 1
	return c.Status(fiber.StatusOK).JSON(res)
}
