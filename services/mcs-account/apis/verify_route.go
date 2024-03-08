package apis

import (
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

func StartVerifyRoute(app *fiber.App, temporal client.Client) {
	controller := VerifyController{
		TemporalClient: temporal,
	}
	group := app.Group("/verifications")
	group.Post("/otp", controller.VerifyOtp)
}
