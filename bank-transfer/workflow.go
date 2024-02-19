package banktransfer

import (
	"time"

	"github.com/mitchellh/mapstructure"
	"go.temporal.io/sdk/workflow"
)

type (
	TransferItem struct {
		ProductId int
		Quantity  int
	}

	TransferState struct {
		Items []TransferItem
		Email string
	}

	UpdateTransferMessage struct {
		Remove bool
		Item   TransferItem
	}
)

var (
	abandonedTransferTimeout = 10 * time.Second
)

func TransferWorkflow(ctx workflow.Context, state TransferState) error {
	// https://docs.temporal.io/docs/concepts/workflows/#workflows-have-options
	logger := workflow.GetLogger(ctx)

	err := workflow.SetQueryHandler(ctx, "getTransfer", func(input []byte) (TransferState, error) {
		return state, nil
	})
	if err != nil {
		logger.Info("SetQueryHandler failed.", "Error", err)
		return err
	}

	addToTransferChannel := workflow.GetSignalChannel(ctx, SignalChannels.ADD_TO_TRANSFER_CHANNEL)
	removeFromTransferChannel := workflow.GetSignalChannel(ctx, SignalChannels.REMOVE_FROM_TRANSFER_CHANNEL)
	updateEmailChannel := workflow.GetSignalChannel(ctx, SignalChannels.UPDATE_EMAIL_CHANNEL)
	checkoutChannel := workflow.GetSignalChannel(ctx, SignalChannels.CHECKOUT_CHANNEL)
	checkedOut := false
	sentAbandonedTransferEmail := false

	var a *Activities

	for {
		selector := workflow.NewSelector(ctx)

		selector.AddReceive(addToTransferChannel, func(c workflow.ReceiveChannel, _ bool) {
			var signal interface{}
			c.Receive(ctx, &signal)

			var message AddToTransferSignal
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

			var message RemoveFromTransferSignal
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

			var message UpdateEmailSignal
			err := mapstructure.Decode(signal, &message)
			if err != nil {
				logger.Error("Invalid signal type %v", err)
				return
			}

			state.Email = message.Email
			sentAbandonedTransferEmail = false
		})

		selector.AddReceive(checkoutChannel, func(c workflow.ReceiveChannel, _ bool) {
			var signal interface{}
			c.Receive(ctx, &signal)

			var message CheckoutSignal
			err := mapstructure.Decode(signal, &message)
			if err != nil {
				logger.Error("Invalid signal type %v", err)
				return
			}

			state.Email = message.Email

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

		if !sentAbandonedTransferEmail && len(state.Items) > 0 {
			selector.AddFuture(workflow.NewTimer(ctx, abandonedTransferTimeout), func(f workflow.Future) {
				sentAbandonedTransferEmail = true
				ao := workflow.ActivityOptions{
					StartToCloseTimeout: time.Minute,
				}

				ctx = workflow.WithActivityOptions(ctx, ao)

				err := workflow.ExecuteActivity(ctx, a.SendTransferNotification, state.Email).Get(ctx, nil)
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

func (state *TransferState) AddToTransfer(item TransferItem) {
	for i := range state.Items {
		if state.Items[i].ProductId != item.ProductId {
			continue
		}

		state.Items[i].Quantity += item.Quantity
		return
	}

	state.Items = append(state.Items, item)
}

func (state *TransferState) RemoveFromTransfer(item TransferItem) {
	for i := range state.Items {
		if state.Items[i].ProductId != item.ProductId {
			continue
		}

		state.Items[i].Quantity -= item.Quantity
		if state.Items[i].Quantity <= 0 {
			state.Items = append(state.Items[:i], state.Items[i+1:]...)
		}
		break
	}
}
