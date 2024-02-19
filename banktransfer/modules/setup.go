package modules

import (
	"github.com/anhgeeky/go-temporal-labs/banktransfer/modules/account"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/modules/transaction"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/repos"
)

func SetupServices() map[string]interface{} {
	// Init repositories
	accountRepo := repos.AccountRepo{}
	transactionRepo := repos.TransactionRepo{}

	// Init services
	accountService := account.Service{
		Repo: accountRepo,
	}
	transactionService := transaction.Service{
		Repo: transactionRepo,
	}

	return map[string]interface{}{
		"accountService":     accountService,
		"transactionService": transactionService,
	}
}
