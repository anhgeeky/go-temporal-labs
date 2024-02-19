package banktransfer

import (
	"context"
)

type Activities struct {
}

func (a *Activities) CreateTransfer(_ context.Context, cart TransferState) error {
	var amount float32 = 0
	var description string = ""
	for _, item := range cart.Items {
		var product Product
		for _, _product := range Products {
			if _product.Id == item.ProductId {
				product = _product
				break
			}
		}
		amount += float32(item.Quantity) * product.Price
		if len(description) > 0 {
			description += ", "
		}
		description += product.Name
	}

	return nil
}

func (a *Activities) SendTransferNotification(_ context.Context, email string) error {
	return nil
}
