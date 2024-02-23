package routes

import (
	"github.com/anhgeeky/go-temporal-labs/mcs-money-transfer/apis/controllers"
	"github.com/anhgeeky/go-temporal-labs/mcs-money-transfer/modules/transaction"
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

func StartTransferRoute(app *fiber.App, temporal client.Client, services map[string]interface{}) {
	controller := controllers.TransferController{
		Service:        services["transactionService"].(transaction.Service),
		TemporalClient: temporal,
	}
	group := app.Group("/transfers")

	group.Post("/", controller.CreateTransfer)
	group.Get("/:workflowID", controller.GetTransfer)
}
