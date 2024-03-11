package notification

import (
	"github.com/anhgeeky/go-temporal-labs/notification/activities"
	"github.com/anhgeeky/go-temporal-labs/notification/config"
	"github.com/anhgeeky/go-temporal-labs/notification/messages"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

// Notification workflow
func SetupNotificationWorkflow(w worker.Worker, workflowFunc func(ctx workflow.Context, state messages.NotificationMessage) error) {
	notificationActivity := &activities.NotificationActivity{}
	w.RegisterWorkflowWithOptions(workflowFunc, workflow.RegisterOptions{Name: config.Workflows.NotificationName}) //workflowcheck:ignore
	w.RegisterActivity(notificationActivity.PushEmail)
	// w.RegisterWorkflow(workflows.NotificationWorkflow)
}
