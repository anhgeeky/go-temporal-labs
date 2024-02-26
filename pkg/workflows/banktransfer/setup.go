package banktransfer

import (
	"github.com/anhgeeky/go-temporal-labs/banktransfer/activities"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/account"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/moneytransfer"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/workflows"
	"go.temporal.io/sdk/worker"
)

// Transfer workflow
func SetupBankTransferWorkflow(w worker.Worker, cfg *config.ExternalConfigs) {
	transferActivity := &activities.TransferActivity{
		AccountService: account.AccountService{
			Host: cfg.AccountHost,
		},
		MoneyTransferService: moneytransfer.MoneyTransferService{
			Host: cfg.AccountHost,
		},
	}
	w.RegisterActivity(transferActivity.CheckBalance)
	w.RegisterActivity(transferActivity.CheckTargetAccount)
	w.RegisterActivity(transferActivity.CreateTransferTransaction)
	w.RegisterActivity(transferActivity.WriteCreditAccount)
	w.RegisterActivity(transferActivity.WriteDebitAccount)
	w.RegisterActivity(transferActivity.AddNewActivity)
	// Rollback
	w.RegisterActivity(transferActivity.CreateTransferTransactionCompensation)
	w.RegisterActivity(transferActivity.WriteCreditAccountCompensation)
	w.RegisterActivity(transferActivity.WriteDebitAccountCompensation)
	w.RegisterActivity(transferActivity.AddNewActivityCompensation)
	w.RegisterWorkflow(workflows.TransferWorkflow)
}
