package workflows

import (
	"time"

	cw "github.com/anhgeeky/go-temporal-labs/core/workflow"
	"github.com/anhgeeky/go-temporal-labs/notification/activities"
	"github.com/anhgeeky/go-temporal-labs/notification/messages"
	"go.temporal.io/sdk/workflow"
)

// ================================================
// Luồng gửi thông báo
// ================================================

func NotificationWorkflow(ctx workflow.Context, state messages.NotificationMessage) (string, error) {
	// https://docs.temporal.io/docs/concepts/workflows/#workflows-have-options
	logger := workflow.GetLogger(ctx)

	err := workflow.SetQueryHandler(ctx, "getNotification", func(input []byte) (messages.NotificationMessage, error) {
		return state, nil
	})
	if err != nil {
		logger.Info("SetQueryHandler failed.", "Error", err)
		return "", err
	}
	logger.Info("NotificationWorkflow start")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout:    10 * time.Second,
		ScheduleToCloseTimeout: 1 * time.Minute,
		RetryPolicy:            cw.WorkflowConfigs.RetryPolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var a *activities.NotificationActivity
	var results []string
	var token *messages.DeviceToken

	err = workflow.ExecuteActivity(ctx, a.GetDeviceToken).Get(ctx, &token)
	if err != nil {
		logger.Error("Failure sending response activity", "error", err)
	}

	// Lấy được Device token thì run các activities parallel
	// Start a goroutine in a workflow safe way
	workflow.Go(ctx, func(gCtx workflow.Context) {
		var result1 string
		err = workflow.ExecuteActivity(gCtx, a.PushInternalApp, token).Get(gCtx, &result1)
		if err != nil {
			return
		}

		var result2 string
		err = workflow.ExecuteActivity(gCtx, a.PushNotification, token).Get(gCtx, &result2)
		if err != nil {
			return
		}

		var result3 string
		err = workflow.ExecuteActivity(gCtx, a.PushSMS, token).Get(gCtx, &result3)
		if err != nil {
			return
		}
		results = append(results, result1, result2, result3)
	})

	logger.Info("NotificationWorkflow before result", results)

	_ = workflow.Await(ctx, func() bool {
		logger.Info("NotificationWorkflow Await result", results)
		return err != nil || len(results) == 3 // 3 task push đã chạy xong -> Completed
	})

	logger.Info("NotificationWorkflow end")

	return "DONE", err
}
