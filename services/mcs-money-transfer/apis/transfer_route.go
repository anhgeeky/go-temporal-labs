package apis

import (
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

func StartTransferRoute(app *fiber.App, temporal client.Client) {
	controller := TransferController{
		TemporalClient: temporal,
	}
	group := app.Group("/transfers")

	group.Post("/", controller.CreateTransfer)
	// Actions
	group.Post("/:workflowID/transactions", controller.CreateTransferTransaction)

}
