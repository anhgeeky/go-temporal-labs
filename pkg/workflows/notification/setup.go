package notification

import (
	"github.com/anhgeeky/go-temporal-labs/notification/activities"
	"github.com/anhgeeky/go-temporal-labs/notification/workflows"
	"go.temporal.io/sdk/worker"
)

// Notification workflow
func SetupNotificationWorkflow(w worker.Worker) {
	notificationActivity := &activities.NotificationActivity{}
	w.RegisterActivity(notificationActivity.GetDeviceToken)
	w.RegisterActivity(notificationActivity.PushSMS)
	w.RegisterActivity(notificationActivity.PushNotification)
	w.RegisterActivity(notificationActivity.PushInternalApp)
	w.RegisterWorkflow(workflows.NotificationWorkflow)
}