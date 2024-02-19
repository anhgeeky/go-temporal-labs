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

	group.Post("/transfers", controller.CreateTransferHandler)
	group.Get("/transfers/:workflowID", controller.GetTransferHandler)
	group.Put("/transfers/:workflowID/add", controller.AddToTransferHandler)
	group.Put("/transfers/:workflowID/remove", controller.RemoveFromTransferHandler)
	group.Put("/transfers/:workflowID/checkout", controller.CheckoutHandler)
	group.Put("/transfers/:workflowID/email", controller.UpdateEmailHandler)
}
