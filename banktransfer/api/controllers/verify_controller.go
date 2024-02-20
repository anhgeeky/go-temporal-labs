package controllers

import (
	"github.com/anhgeeky/go-temporal-labs/banktransfer/modules/otp"
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

type VerifyController struct {
	Service        otp.Service
	TemporalClient client.Client
}

// TODO: Bổ sung
func (r VerifyController) VerifyOtp(c *fiber.Ctx) error {
	return nil
}
