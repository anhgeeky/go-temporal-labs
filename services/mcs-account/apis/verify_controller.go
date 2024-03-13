package apis

import (
	"context"
	"encoding/json"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/messages"
	"github.com/anhgeeky/go-temporal-labs/core/apis/responses"
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

type VerifyController struct {
	TemporalClient client.Client
}

// 2. Xác thực OTP
func (r VerifyController) VerifyOtp(c *fiber.Ctx) error {
	var item messages.VerifyOtpReq
	json.Unmarshal(c.Body(), &item)

	update := messages.VerifiedOtpSignal{Item: item}

	// Trigger Signal Transfer Flow
	err := r.TemporalClient.SignalWorkflow(context.Background(), item.FlowId, "", config.SignalChannels.VERIFY_OTP_CHANNEL, update)
	if err != nil {
		return responses.WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["ok"] = 1
	return c.Status(fiber.StatusOK).JSON(res)
}
