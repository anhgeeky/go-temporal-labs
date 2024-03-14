package notification

import (
	"log"

	"github.com/anhgeeky/go-temporal-labs/notification/activities"
	"github.com/anhgeeky/go-temporal-labs/notification/config"
	"github.com/anhgeeky/go-temporal-labs/notification/workflows"
	"github.com/spf13/viper"
	"go.temporal.io/sdk/worker"
)

// Notification workflow
func SetupNotificationWorkflow(w worker.Worker) {
	cfg := &config.EmailConfig{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		log.Fatalln("Could not load `EmailConfig` configuration", err)
	}

	notificationActivity := &activities.NotificationActivity{
		Config: cfg,
	}
	w.RegisterActivity(notificationActivity.PushEmail)
	w.RegisterWorkflow(workflows.NotificationWorkflow)
}
