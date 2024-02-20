package workflows

import (
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/messages"
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

	// TODO: Bổ sung

	return "DONE", err
}
