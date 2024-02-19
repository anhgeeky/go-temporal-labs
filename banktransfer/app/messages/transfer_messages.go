package messages

type TransferItem struct {
	ProductId int
	Quantity  int
}

type UpdateTransferMessage struct {
	Remove bool
	Item   TransferItem
}

type TransferState struct {
	Items []TransferItem
	Email string
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
