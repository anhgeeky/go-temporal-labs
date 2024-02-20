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

	group.Post("/", controller.CreateTransfer)
	group.Get("/:workflowID", controller.GetTransfer)
	group.Put("/:workflowID/add", controller.AddToTransfer)
	group.Put("/:workflowID/remove", controller.RemoveFromTransfer)
	group.Put("/:workflowID/checkout", controller.Checkout)
	group.Put("/:workflowID/email", controller.UpdateEmail)
