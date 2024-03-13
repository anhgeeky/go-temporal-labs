package notification

import (
	"github.com/anhgeeky/go-temporal-labs/notification/activities"
	"github.com/anhgeeky/go-temporal-labs/notification/config"
	"github.com/anhgeeky/go-temporal-labs/notification/workflows"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func NotificationWorkflowRegister(w worker.Registry) {
	notificationActivity := &activities.NotificationActivity{}
	w.RegisterWorkflowWithOptions(workflows.NotificationWorkflow, workflow.RegisterOptions{Name: config.Workflows.NotificationName})
	w.RegisterActivity(notificationActivity.PushEmail)
}
