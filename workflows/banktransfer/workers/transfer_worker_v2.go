package workers

import (
	"github.com/anhgeeky/go-temporal-labs/banktransfer"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/activities"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/account"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/moneytransfer"
	"github.com/anhgeeky/go-temporal-labs/core/broker"
	"github.com/anhgeeky/go-temporal-labs/notification"
	"go.temporal.io/sdk/worker"
)

type TransferWorkerV2 struct {
	Broker broker.Broker
	Config config.ExternalConfig
}

func (r TransferWorkerV2) Register(register worker.Registry) {
	transferActivity := &activities.TransferActivity{
		Broker: r.Broker,
		AccountService: account.AccountService{
			Host: r.Config.AccountHost,
		},
		MoneyTransferService: moneytransfer.MoneyTransferService{
			Host: r.Config.MoneyTransferHost,
		},
	}

	banktransfer.RegisterTransferWorkflowV2(register, *transferActivity)
	notification.RegisterNotificationWorkflow(register)
}
