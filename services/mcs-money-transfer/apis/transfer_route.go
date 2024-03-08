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
	// group.Get("/:workflowID", controller.GetTransfer)
	// Actions
	group.Post("/:workflowID/transactions", controller.CreateTransferTransaction)
	// group.Post("/:workflowID/credit-accounts", controller.WriteCreditAccount)
	// group.Post("/:workflowID/debit-accounts", controller.WriteDebitAccount)
	// group.Post("/:workflowID/new-activity", controller.AddNewActivity)
	// Rollback
	// group.Post("/:workflowID/transactions/rollback", controller.CreateTransferTransactionCompensation)
	// group.Post("/:workflowID/credit-accounts/rollback", controller.WriteCreditAccountCompensation)
	// group.Post("/:workflowID/debit-accounts/rollback", controller.WriteDebitAccountCompensation)
	// group.Post("/:workflowID/new-activity/rollback", controller.AddNewActivityCompensation)
}
