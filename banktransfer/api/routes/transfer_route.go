package routes

import (
	"github.com/anhgeeky/go-temporal-labs/banktransfer/api/controllers"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/modules/transaction"
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

func StartTransferRoute(app *fiber.App, temporal client.Client, services map[string]interface{}) {
	controller := controllers.TransferController{
		Service:        services["transactionService"].(transaction.Service),
		TemporalClient: temporal,
	}
	group := app.Group("/transfers")

	group.Post("/", controller.CreateTransferHandler)
	group.Get("/:workflowID", controller.GetTransferHandler)
	group.Put("/:workflowID/add", controller.AddToTransferHandler)
	group.Put("/:workflowID/remove", controller.RemoveFromTransferHandler)
	group.Put("/:workflowID/checkout", controller.CheckoutHandler)
	group.Put("/:workflowID/email", controller.UpdateEmailHandler)
}
