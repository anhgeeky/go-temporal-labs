package notification

import (
	"github.com/anhgeeky/go-temporal-labs/notification/activities"
	"github.com/anhgeeky/go-temporal-labs/notification/workflows"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

// Notification workflow
func SetupNotificationWorkflow(w worker.Worker) {
	notificationActivity := &activities.NotificationActivity{}
	w.RegisterActivity(notificationActivity.PushEmail)
	w.RegisterWorkflowWithOptions(workflows.NotificationWorkflow, workflow.RegisterOptions{Name: "NotificationWorkflow"})
}
