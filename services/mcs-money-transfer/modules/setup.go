package modules

import (
	"github.com/anhgeeky/go-temporal-labs/mcs-money-transfer/modules/account"
	"github.com/anhgeeky/go-temporal-labs/mcs-money-transfer/modules/notification"
	"github.com/anhgeeky/go-temporal-labs/mcs-money-transfer/modules/otp"
	"github.com/anhgeeky/go-temporal-labs/mcs-money-transfer/modules/transaction"
	"github.com/anhgeeky/go-temporal-labs/mcs-money-transfer/repos"
)

func SetupServices() map[string]interface{} {
	// Init repositories
	accountRepo := repos.AccountRepo{}
	notificationRepo := repos.NotificationRepo{}
	transactionRepo := repos.TransactionRepo{}
	otpRepo := repos.OtpRepo{}

	// Init services
	accountService := account.Service{
		Repo: accountRepo,
	}
	notificationService := notification.Service{
		Repo: notificationRepo,
	}
	transactionService := transaction.Service{
		Repo: transactionRepo,
	}
	otpService := otp.Service{
		Repo: otpRepo,
	}

	return map[string]interface{}{
		"accountService":      accountService,
		"transactionService":  transactionService,
		"notificationService": notificationService,
		"otpService":          otpService,
	}
}
