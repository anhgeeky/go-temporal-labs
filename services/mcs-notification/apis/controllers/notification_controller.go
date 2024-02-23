package controllers

import (
	"github.com/anhgeeky/go-temporal-labs/mcs-notification/modules/notification"
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

type NotificationController struct {
	Service        notification.Service
	TemporalClient client.Client
}

func (r NotificationController) NotificationOtp(c *fiber.Ctx) error {
	res := make(map[string]interface{})
	res["ok"] = 1
	return c.Status(fiber.StatusOK).JSON(res)
}
