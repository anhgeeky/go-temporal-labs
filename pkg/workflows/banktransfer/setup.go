package banktransfer

import (
	"github.com/anhgeeky/go-temporal-labs/banktransfer/activities"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/workflows"
	"go.temporal.io/sdk/worker"
)

// Transfer workflow
func SetupBankTransferWorkflow(w worker.Worker) {
	transferActivity := &activities.TransferActivity{}
	w.RegisterActivity(transferActivity.CreateTransfer)
	w.RegisterActivity(transferActivity.CheckBalance)
	w.RegisterActivity(transferActivity.CheckTargetAccount)
	w.RegisterActivity(transferActivity.CreateTransferTransaction)
	w.RegisterActivity(transferActivity.WriteCreditAccount)
	w.RegisterActivity(transferActivity.WriteDebitAccount)
	// w.RegisterActivity(transferActivity.AddNewActivity)
	w.RegisterWorkflow(workflows.TransferWorkflow)
}
