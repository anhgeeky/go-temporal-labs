package notification

import (
	"github.com/anhgeeky/go-temporal-labs/notification/activities"
	"github.com/anhgeeky/go-temporal-labs/notification/config"
	"github.com/anhgeeky/go-temporal-labs/notification/outbound/notification"
	"github.com/anhgeeky/go-temporal-labs/notification/workflows"
	"go.temporal.io/sdk/worker"
)

// Notification workflow
func SetupNotificationWorkflow(w worker.Worker) {
	notificationActivity := &activities.NotificationActivity{
		NotificationService: notification.NotificationService{
			Host: config.MCS_NOTIFICATION_HOST,
		},
	}
	w.RegisterActivity(notificationActivity.GetDeviceToken)
	w.RegisterActivity(notificationActivity.PushSMS)
	w.RegisterActivity(notificationActivity.PushNotification)
	w.RegisterActivity(notificationActivity.PushInternalApp)
	w.RegisterWorkflow(workflows.NotificationWorkflow)
}
