package modules

import (
	"github.com/anhgeeky/go-temporal-labs/core/broker"
	"github.com/anhgeeky/go-temporal-labs/mcs-money-transfer/modules/transaction"
	"github.com/anhgeeky/go-temporal-labs/mcs-money-transfer/repos"
)

func SetupServices(bk broker.Broker) map[string]interface{} {
	// Init repositories
	transactionRepo := repos.TransactionRepo{}

	// Init services
	transactionService := transaction.Service{
		Repo:   transactionRepo,
		Broker: bk,
	}

	return map[string]interface{}{
		"transactionService": transactionService,
	}
}
