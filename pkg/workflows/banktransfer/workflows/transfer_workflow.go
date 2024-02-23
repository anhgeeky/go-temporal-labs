package workflows

import (
	"fmt"
	"time"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/activities"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/messages"
	coreWorkflow "github.com/anhgeeky/go-temporal-labs/core/workflow"
	notiMsg "github.com/anhgeeky/go-temporal-labs/notification/messages"
	notiWorkflows "github.com/anhgeeky/go-temporal-labs/notification/workflows"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/multierr"
)

var (
	transferTimeout = 5 * time.Second
)

func TransferWorkflow(ctx workflow.Context, state messages.Transfer) (err error) {
	// https://docs.temporal.io/docs/concepts/workflows/#workflows-have-options
	logger := workflow.GetLogger(ctx)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 2 * time.Minute,
		HeartbeatTimeout:    10 * time.Second,
		RetryPolicy:         coreWorkflow.WorkflowConfigs.RetryPolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	verifyOtpChannel := workflow.GetSignalChannel(ctx, config.SignalChannels.VERIFY_OTP_CHANNEL)
	verifiedOtp := false
	completed := false

	var a *activities.TransferActivity

	for {
		selector := workflow.NewSelector(ctx)

		selector.AddReceive(verifyOtpChannel, func(c workflow.ReceiveChannel, _ bool) {

			var signal interface{}
			c.Receive(ctx, &signal)

			var message messages.VerifiedOtpSignal
			err = mapstructure.Decode(signal, &message)
			if err != nil {
				logger.Error("Invalid signal type %v", err)
				return
			}

			err = workflow.ExecuteActivity(ctx, a.CheckBalance, state).Get(ctx, nil)
			if err != nil {
				return
			}

			err = workflow.ExecuteActivity(ctx, a.CheckTargetAccount, state).Get(ctx, nil)
			if err != nil {
				return
			}

			err = workflow.ExecuteActivity(ctx, a.CreateTransferTransaction, state).Get(ctx, nil)
			if err != nil {
				return
			}

			// Compensation
			defer func() {
				if err != nil {
					errCompensation := workflow.ExecuteActivity(ctx, a.CreateTransferTransactionCompensation, state).Get(ctx, nil)
					err = multierr.Append(err, errCompensation)
				}
			}()

			err = workflow.ExecuteActivity(ctx, a.WriteCreditAccount, state).Get(ctx, nil)
			if err != nil {
				return
			}

			// Compensation
			defer func() {
				if err != nil {
					errCompensation := workflow.ExecuteActivity(ctx, a.WriteCreditAccountCompensation, state).Get(ctx, nil)
					err = multierr.Append(err, errCompensation)
				}
			}()

			err = workflow.ExecuteActivity(ctx, a.WriteDebitAccount, state).Get(ctx, nil)
			if err != nil {
				return
			}

			// Compensation
			defer func() {
				if err != nil {
					errCompensation := workflow.ExecuteActivity(ctx, a.WriteDebitAccountCompensation, state).Get(ctx, nil)
					err = multierr.Append(err, errCompensation)
				}
			}()

			verifiedOtp = true
		})

		// Call subflow -> Gửi notification
		if !completed && verifiedOtp {
			selector.AddFuture(workflow.NewTimer(ctx, transferTimeout), func(f workflow.Future) {
				execution := workflow.GetInfo(ctx).WorkflowExecution
				childID := fmt.Sprintf("NOTIFICATION: %v", execution.RunID)
				cwo := workflow.ChildWorkflowOptions{
					WorkflowID: childID,
				}
				ctx = workflow.WithChildOptions(ctx, cwo)

				msgNotfication := notiMsg.NotificationMessage{
					// TODO: Bổ sung payload
					Token: notiMsg.DeviceToken{
						FirebaseToken: uuid.New().String(),
					},
				}

				var result string
				err = workflow.ExecuteChildWorkflow(ctx, notiWorkflows.NotificationWorkflow, msgNotfication).Get(ctx, &result)
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
	return
}
