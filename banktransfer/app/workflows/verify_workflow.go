package workflows

import (
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/messages"
	"go.temporal.io/sdk/workflow"
)

// ================================================
// Xác thực trước khi chạy 1 luồng xử lý
// ================================================

func VerifyWorkflow(ctx workflow.Context, state messages.VerifyMessage) error {
	// https://docs.temporal.io/docs/concepts/workflows/#workflows-have-options
	logger := workflow.GetLogger(ctx)

	err := workflow.SetQueryHandler(ctx, "getVerify", func(input []byte) (messages.VerifyMessage, error) {
		return state, nil
	})
	if err != nil {
		logger.Info("SetQueryHandler failed.", "Error", err)
		return err
	}

	// TODO: Bổ sung

	return nil
}
