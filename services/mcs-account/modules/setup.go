package modules

import (
	"github.com/anhgeeky/go-temporal-labs/core/broker"
	"github.com/anhgeeky/go-temporal-labs/mcs-account/modules/account"
	"github.com/anhgeeky/go-temporal-labs/mcs-account/modules/otp"
	"github.com/anhgeeky/go-temporal-labs/mcs-account/repos"
)

func SetupServices(bk broker.Broker) map[string]interface{} {
	// Init repositories
	accountRepo := repos.AccountRepo{}
	otpRepo := repos.OtpRepo{}

	// Init services
	accountService := account.Service{
		Repo:   accountRepo,
		Broker: bk,
	}
	otpService := otp.Service{
		Repo:   otpRepo,
		Broker: bk,
	}

	return map[string]interface{}{
		"accountService": accountService,
		"otpService":     otpService,
	}
}
