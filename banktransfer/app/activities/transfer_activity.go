package activities

import (
	"context"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/messages"
	"go.temporal.io/sdk/activity"
)

type TransferActivity struct {
}

func (a *TransferActivity) CreateTransfer(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)

	logger.Info("activity: create transfer", msg)

	return nil
}

func (a *TransferActivity) SendTransferNotification(_ context.Context, msg messages.Transfer) error {
	return nil
}
