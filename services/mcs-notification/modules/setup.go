package modules

import (
	"github.com/anhgeeky/go-temporal-labs/mcs-notification/modules/notification"
	"github.com/anhgeeky/go-temporal-labs/mcs-notification/repos"
)

func SetupServices() map[string]interface{} {
	// Init repositories
	notificationRepo := repos.NotificationRepo{}

	// Init services
	notificationService := notification.Service{
		Repo: notificationRepo,
	}

	return map[string]interface{}{
		"notificationService": notificationService,
	}
}
