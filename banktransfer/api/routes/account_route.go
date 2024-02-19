package routes

import (
	"github.com/anhgeeky/go-temporal-labs/banktransfer/api/controllers"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/modules/account"
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

func StartAccountRoute(app *fiber.App, temporal client.Client, services map[string]interface{}) {
	controller := controllers.AccountController{
		Service:        services["accountService"].(account.Service),
		TemporalClient: temporal,
	}
	group := app.Group("/accounts")
	group.Get("/", controller.GetAccountsHandler)
}