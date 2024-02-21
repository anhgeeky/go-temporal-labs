package workflows

import (
	"fmt"
	"time"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/activities"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/configs"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/messages"
	"github.com/mitchellh/mapstructure"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

var (
	abandonedTransferTimeout = 10 * time.Second
)

func TransferWorkflow(ctx workflow.Context, state messages.Transfer) error {
	// https://docs.temporal.io/docs/concepts/workflows/#workflows-have-options
	logger := workflow.GetLogger(ctx)

	err := workflow.SetQueryHandler(ctx, "getTransfer", func(input []byte) (messages.Transfer, error) {
		return state, nil
	})
	if err != nil {
		logger.Info("SetQueryHandler failed.", "Error", err)
		return err
	}

	ao := workflow.ActivityOptions{
		// ScheduleToStartTimeout: 5 * time.Second,
		// ScheduleToCloseTimeout: 1 * time.Minute,
		RetryPolicy: &temporal.RetryPolicy{MaximumAttempts: 2},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	verifyOtpChannel := workflow.GetSignalChannel(ctx, configs.SignalChannels.VERIFY_OTP_CHANNEL)
	verifiedOtp := false
	completed := false

	var a *activities.TransferActivity

	for {
		// childCtx, cancelHandler := workflow.WithCancel(ctx)
		selector := workflow.NewSelector(ctx)

		selector.AddReceive(verifyOtpChannel, func(c workflow.ReceiveChannel, _ bool) {
			var signal interface{}
			c.Receive(ctx, &signal)

			var message messages.VerifiedOtpSignal
			err := mapstructure.Decode(signal, &message)
			if err != nil {
				logger.Error("Invalid signal type %v", err)
				return
			}

			err = workflow.ExecuteActivity(ctx, a.CheckBalance, state).Get(ctx, nil)
			if err != nil {
				logger.Error("Failure sending response activity", "error", err)
				return
			}

			err = workflow.ExecuteActivity(ctx, a.CheckTargetAccount, state).Get(ctx, nil)
			if err != nil {
				logger.Error("Failure sending response activity", "error", err)
				return
			}

			err = workflow.ExecuteActivity(ctx, a.CreateTransferTransaction, state).Get(ctx, nil)
			if err != nil {
				logger.Error("Failure sending response activity", "error", err)
				return
			}

			err = workflow.ExecuteActivity(ctx, a.WriteCreditAccount, state).Get(ctx, nil)
			if err != nil {
				logger.Error("Failure sending response activity", "error", err)
				return
			}

			err = workflow.ExecuteActivity(ctx, a.WriteDebitAccount, state).Get(ctx, nil)
			if err != nil {
				logger.Error("Failure sending response activity", "error", err)
				return
			}

			verifiedOtp = true
		})

		// Call subflow -> Gửi notification
		if !completed && verifiedOtp {
			selector.AddFuture(workflow.NewTimer(ctx, abandonedTransferTimeout), func(f workflow.Future) {
				execution := workflow.GetInfo(ctx).WorkflowExecution
				childID := fmt.Sprintf("TRANSFER:%v", execution.RunID)
				cwo := workflow.ChildWorkflowOptions{
					WorkflowID: childID,
				}
				ctx = workflow.WithChildOptions(ctx, cwo)

				msgNotfication := messages.NotificationMessage{
					// TODO: Bổ sung payload
				}

				var result string
				err = workflow.ExecuteChildWorkflow(ctx, NotificationWorkflow, msgNotfication).Get(ctx, &result)
				if err != nil {
					logger.Error("Parent execution received child execution failure.", "Error", err)
					return
				}
				// ===============================================================================
				logger.Info("Parent execution completed.", "Result", result)

				completed = true
			})
		}

		selector.Select(ctx)

		// Xử lý transfer hoàn tất
		if completed {
			break
		}
	}

	logger.Info("Workflow completed.")
	return nil
}
