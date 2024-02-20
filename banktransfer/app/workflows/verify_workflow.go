package workflows

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/activities"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/messages"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// ================================================
// Xác thực trước khi chạy 1 luồng xử lý
// ================================================

func VerifyOtpWorkflow(ctx workflow.Context, state messages.VerifyOtp) error {
	// https://docs.temporal.io/docs/concepts/workflows/#workflows-have-options
	logger := workflow.GetLogger(ctx)
	err := workflow.SetQueryHandler(ctx, "getVerifyOtp", func(input []byte) (messages.VerifyOtp, error) {
		return state, nil
	})
	if err != nil {
		logger.Info("SetQueryHandler failed.", "Error", err)
		return err
	}

	var a *activities.VerifyActivity
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 2 * time.Minute,
		HeartbeatTimeout:    10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute,
			MaximumAttempts:    5,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	err = workflow.ExecuteActivity(ctx, a.VerifyOtp, state).Get(ctx, nil)
	if err != nil {
		logger.Info("Workflow completed with error.", "Error", err)
		return err
	}

	// Go to next flow
	execution := workflow.GetInfo(ctx).WorkflowExecution
	childID := fmt.Sprintf("child_workflow:%v", execution.RunID)
	cwo := workflow.ChildWorkflowOptions{
		WorkflowID: childID,
	}
	ctx = workflow.WithChildOptions(ctx, cwo)

	var payload messages.Transfer
	json.Unmarshal([]byte(state.Payload), &payload)
	var result string
	err = workflow.ExecuteChildWorkflow(ctx, TransferWorkflow, payload).Get(ctx, &result)
	if err != nil {
		logger.Error("Parent execution received child execution failure.", "Error", err)
		return err
	}

	logger.Info("Workflow completed.")

	return nil
}
