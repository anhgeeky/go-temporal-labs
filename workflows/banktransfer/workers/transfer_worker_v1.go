package workers

import (
	"github.com/anhgeeky/go-temporal-labs/banktransfer/activities"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/account"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/moneytransfer"
	tranFlow "github.com/anhgeeky/go-temporal-labs/banktransfer/workflows"
	"github.com/anhgeeky/go-temporal-labs/core/broker"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

type TransferWorkerV1 struct {
	Broker broker.Broker
	Config config.ExternalConfig
}

func (r TransferWorkerV1) Register(register worker.Registry) {
	transferActivity := &activities.TransferActivity{
		Broker: r.Broker,
		AccountService: account.AccountService{
			Host: r.Config.AccountHost,
		},
		MoneyTransferService: moneytransfer.MoneyTransferService{
			Host: r.Config.MoneyTransferHost,
		},
	}

	register.RegisterWorkflowWithOptions(tranFlow.TransferWorkflow, workflow.RegisterOptions{Name: config.Workflows.TransferWorkflow})
	register.RegisterActivity(transferActivity.CheckBalance)
	register.RegisterActivity(transferActivity.CreateOTP)
	register.RegisterActivity(transferActivity.CreateTransaction)
}
