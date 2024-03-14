package messages

type FundTransferStatus int

const (
	TransactionStarted FundTransferStatus = iota
	TransactionVerified
	TransactionProcessing
	TransactionSucceeded
)

type NotificationMessage struct {
	WorflowId   string             `json:"worflowId"`
	FromAccount string             `json:"fromAccount"`
	ToAccount   string             `json:"toAccount"`
	Amount      float64            `json:"amount"`
	CRefNum     string             `json:"cRefNum"`
	CreatedAt   string             `json:"createdAt"`
	TransferAt  string             `json:"transferAt"`
	Status      FundTransferStatus `json:"status"`
	TransNo     string             `json:"transNo"`
}
