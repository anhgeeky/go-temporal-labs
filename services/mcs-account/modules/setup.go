package modules

import (
	"github.com/anhgeeky/go-temporal-labs/mcs-account/modules/account"
	"github.com/anhgeeky/go-temporal-labs/mcs-account/repos"
)

func SetupServices() map[string]interface{} {
	// Init repositories
	accountRepo := repos.AccountRepo{}

	// Init services
	accountService := account.Service{
		Repo: accountRepo,
	}

	return map[string]interface{}{
		"accountService": accountService,
	}
}
