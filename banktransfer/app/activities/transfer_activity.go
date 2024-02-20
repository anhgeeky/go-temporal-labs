package activities

import (
	"context"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/messages"
)

type TransferActivity struct {
}

func (a *TransferActivity) CreateTransfer(_ context.Context, msg messages.Transfer) error {
	return nil
}

func (a *TransferActivity) SendTransferNotification(_ context.Context, msg messages.Transfer) error {
	return nil
}
