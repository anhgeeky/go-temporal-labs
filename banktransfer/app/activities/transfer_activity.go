package activities

import (
	"context"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/messages"
)

type TransferActivity struct {
}

func (a *TransferActivity) CreateTransfer(_ context.Context, msg messages.TransferState) error {
	// var amount float32 = 0
	// var description string = ""
	// for _, item := range msg.Items {
	// 	var product account.Account
	// 	for _, _product := range account.Accounts {
	// 		if _product.Id == item.Id {
	// 			product = _product
	// 			break
	// 		}
	// 	}
	// 	amount += float32(item.Quantity) * product.Price
	// 	if len(description) > 0 {
	// 		description += ", "
	// 	}
	// 	description += product.Name
	// }

	return nil
}

func (a *TransferActivity) SendTransferNotification(_ context.Context, email string) error {
	return nil
}
