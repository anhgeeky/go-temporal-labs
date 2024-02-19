package messages

type RouteSignal struct {
	Route string
}

type AddToTransferSignal struct {
	Route string
	Item  TransferItem
}

type RemoveFromTransferSignal struct {
	Route string
	Item  TransferItem
}

type UpdateEmailSignal struct {
	Route string
	Email string
}

type CheckoutSignal struct {
	Route string
	Email string
}
