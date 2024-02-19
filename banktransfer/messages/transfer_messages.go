package messages

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
