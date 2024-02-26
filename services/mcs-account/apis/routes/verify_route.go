package routes

import (
	"github.com/anhgeeky/go-temporal-labs/mcs-account/apis/controllers"
	"github.com/anhgeeky/go-temporal-labs/mcs-account/modules/otp"
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

func StartVerifyRoute(app *fiber.App, temporal client.Client, services map[string]interface{}) {
	controller := controllers.VerifyController{
		Service:        services["otpService"].(otp.Service),
		TemporalClient: temporal,
	}
	group := app.Group("/verifications")
	group.Post("/otp", controller.VerifyOtp)
}