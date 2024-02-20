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

	addToTransferChannel := workflow.GetSignalChannel(ctx, configs.SignalChannels.ADD_TO_TRANSFER_CHANNEL)
	removeFromTransferChannel := workflow.GetSignalChannel(ctx, configs.SignalChannels.REMOVE_FROM_TRANSFER_CHANNEL)
	updateEmailChannel := workflow.GetSignalChannel(ctx, configs.SignalChannels.UPDATE_EMAIL_CHANNEL)
	checkoutChannel := workflow.GetSignalChannel(ctx, configs.SignalChannels.CHECKOUT_CHANNEL)
	checkedOut := false
	sentAbandonedTransferEmail := false

	var a *activities.TransferActivity

	for {
		selector := workflow.NewSelector(ctx)

		selector.AddReceive(addToTransferChannel, func(c workflow.ReceiveChannel, _ bool) {
			var signal interface{}
			c.Receive(ctx, &signal)

			var message messages.AddToTransferSignal
			err := mapstructure.Decode(signal, &message)
			if err != nil {
				logger.Error("Invalid signal type %v", err)
				return
			}

			state.AddToTransfer(message.Item)
		})

		selector.AddReceive(removeFromTransferChannel, func(c workflow.ReceiveChannel, _ bool) {
			var signal interface{}
			c.Receive(ctx, &signal)

			var message messages.RemoveFromTransferSignal
			err := mapstructure.Decode(signal, &message)
			if err != nil {
				logger.Error("Invalid signal type %v", err)
				return
			}

			state.RemoveFromTransfer(message.Item)
		})

		selector.AddReceive(updateEmailChannel, func(c workflow.ReceiveChannel, _ bool) {
			var signal interface{}
			c.Receive(ctx, &signal)

			var message messages.UpdateEmailSignal
			err := mapstructure.Decode(signal, &message)
			if err != nil {
				logger.Error("Invalid signal type %v", err)
				return
			}

			// state.Email = message.Email
			sentAbandonedTransferEmail = false
		})

		selector.AddReceive(checkoutChannel, func(c workflow.ReceiveChannel, _ bool) {
			var signal interface{}
			c.Receive(ctx, &signal)

			var message messages.CheckoutSignal
			err := mapstructure.Decode(signal, &message)
			if err != nil {
				logger.Error("Invalid signal type %v", err)
				return
			}

			// state.Email = message.Email

			ao := workflow.ActivityOptions{
				StartToCloseTimeout: time.Minute,
			}

			ctx = workflow.WithActivityOptions(ctx, ao)

			err = workflow.ExecuteActivity(ctx, a.CreateTransfer, state).Get(ctx, nil)
			if err != nil {
				logger.Error("Error creating stripe charge: %v", err)
				return
			}

			checkedOut = true
		})

		if !sentAbandonedTransferEmail {
			selector.AddFuture(workflow.NewTimer(ctx, abandonedTransferTimeout), func(f workflow.Future) {
				sentAbandonedTransferEmail = true
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

		if checkedOut {
			break
		}
	}

	return nil
}
