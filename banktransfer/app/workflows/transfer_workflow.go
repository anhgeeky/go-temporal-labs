package workflows

import (
	"time"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/activities"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/configs"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/messages"
	"github.com/mitchellh/mapstructure"
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

	verifyOtpChannel := workflow.GetSignalChannel(ctx, configs.SignalChannels.VERIFY_OTP_CHANNEL)
	verifiedOtp := false
	completed := false

	var a *activities.TransferActivity

	for {
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

			verifiedOtp = true
		})

		if !verifiedOtp {
			selector.AddFuture(workflow.NewTimer(ctx, abandonedTransferTimeout), func(f workflow.Future) {
				completed = true
				ao := workflow.ActivityOptions{
					StartToCloseTimeout: time.Minute,
				}

				ctx = workflow.WithActivityOptions(ctx, ao)

				err := workflow.ExecuteActivity(ctx, a.SendTransferNotification, state).Get(ctx, nil)
				if err != nil {
					logger.Error("Error sending email %v", err)
					return
				}
			})
		}

		selector.Select(ctx)

		// Xử lý transfer hoàn tất
		if completed {
			break
		}
	}

	return nil
}
