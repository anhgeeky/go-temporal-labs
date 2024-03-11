package banktransfer

import (
	"github.com/anhgeeky/go-temporal-labs/banktransfer/activities"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/messages"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/account"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/moneytransfer"
	"github.com/anhgeeky/go-temporal-labs/core/broker"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

// Transfer workflow
func SetupBankTransferWorkflow(w worker.Worker, workflowFunc func(ctx workflow.Context, state messages.Transfer) error, cfg *config.ExternalConfig, bk broker.Broker) {
	transferActivity := &activities.TransferActivity{
		Broker: bk,
		AccountService: account.AccountService{
			Host: cfg.AccountHost,
		},
		MoneyTransferService: moneytransfer.MoneyTransferService{
			Host: cfg.MoneyTransferHost,
		},
	}
	w.RegisterWorkflowWithOptions(workflowFunc, workflow.RegisterOptions{Name: config.Workflows.TransferName}) //workflowcheck:ignore
	w.RegisterActivity(transferActivity.CheckBalance)
	w.RegisterActivity(transferActivity.CreateOTP)
	w.RegisterActivity(transferActivity.CreateTransaction)
	// w.RegisterWorkflow(workflows.TransferWorkflow)
}

// Transfer workflow V2
func SetupBankTransferWorkflowV2(w worker.Worker, workflowFunc func(ctx workflow.Context, state messages.Transfer) error, cfg *config.ExternalConfig, bk broker.Broker) {
	transferActivity := &activities.TransferActivity{
		Broker: bk,
		AccountService: account.AccountService{
			Host: cfg.AccountHost,
		},
		MoneyTransferService: moneytransfer.MoneyTransferService{
			Host: cfg.MoneyTransferHost,
		},
	}
	w.RegisterWorkflowWithOptions(workflowFunc, workflow.RegisterOptions{Name: config.Workflows.TransferName}) //workflowcheck:ignore
	w.RegisterActivity(transferActivity.CheckBalance)
	w.RegisterActivity(transferActivity.CreateOTP)
	w.RegisterActivity(transferActivity.CreateTransaction)
	// w.RegisterWorkflow(workflows.TransferWorkflowV2)
}

// Transfer workflow V3
func SetupBankTransferWorkflowV3(w worker.Worker, workflowFunc func(ctx workflow.Context, state messages.Transfer) error, cfg *config.ExternalConfig, bk broker.Broker) {
	transferActivity := &activities.TransferActivity{
		Broker: bk,
		AccountService: account.AccountService{
			Host: cfg.AccountHost,
		},
		MoneyTransferService: moneytransfer.MoneyTransferService{
			Host: cfg.MoneyTransferHost,
		},
	}
	w.RegisterWorkflowWithOptions(workflowFunc, workflow.RegisterOptions{Name: config.Workflows.TransferName}) //workflowcheck:ignore
	w.RegisterActivity(transferActivity.CheckBalance)
	w.RegisterActivity(transferActivity.CreateOTP)
	w.RegisterActivity(transferActivity.CreateTransaction)
	w.RegisterActivity(transferActivity.NewActivityForTest1)
	// w.RegisterWorkflow(workflows.TransferWorkflowV3)
}

// Transfer workflow V4
func SetupBankTransferWorkflowV4(w worker.Worker, workflowFunc func(ctx workflow.Context, state messages.Transfer) error, cfg *config.ExternalConfig, bk broker.Broker) {
	transferActivity := &activities.TransferActivity{
		Broker: bk,
		AccountService: account.AccountService{
			Host: cfg.AccountHost,
		},
		MoneyTransferService: moneytransfer.MoneyTransferService{
			Host: cfg.MoneyTransferHost,
		},
	}
	w.RegisterWorkflowWithOptions(workflowFunc, workflow.RegisterOptions{Name: config.Workflows.TransferName}) //workflowcheck:ignore
	w.RegisterActivity(transferActivity.CheckBalance)
	w.RegisterActivity(transferActivity.CreateOTP)
	w.RegisterActivity(transferActivity.CreateTransaction)
	w.RegisterActivity(transferActivity.NewActivityForTest22)
	// w.RegisterWorkflow(workflows.TransferWorkflowV4)
}
