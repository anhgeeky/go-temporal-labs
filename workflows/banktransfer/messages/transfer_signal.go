package messages

type VerifyOtpReq struct {
	FlowId string `json:"workflow_id"` // WorkflowID
	Token  string `json:"token"`
	Code   string `json:"code"`
	Trace  string `json:"trace"`
}

type VerifiedOtpSignal struct {
	WorkflowID  string  `json:"workflowId"`
	FromAccount string  `json:"fromAccount"`
	ToAccount   string  `json:"toAccount"`
	CRefNum     string  `json:"cRefNum"`
	Amount      float64 `json:"amount"`
	CreatedAt   string  `json:"createdAt"`
	TransferAt  string  `json:"transferAt"`
	TransNo     string  `json:"transNo"`
	Status      int     `json:"status"`
}

type CreateTransactionReq struct {
	CRefNum string `json:"cRefNum"`
}

type FundTransferStatus int

const (
	TransactionStarted FundTransferStatus = iota
	TransactionVerified
	TransactionProcessing
	TransactionSucceeded
)

type CreateTransactionSignal struct {
	WorkflowID  string             `json:"workflowId"`
	FromAccount string             `json:"fromAccount"`
	ToAccount   string             `json:"toAccount"`
	CRefNum     string             `json:"cRefNum"`
	Amount      float64            `json:"amount"`
	CreatedAt   string             `json:"createdAt"`
	TransferAt  string             `json:"transferAt"`
	TransNo     string             `json:"transNo"`
	Status      FundTransferStatus `json:"status"`
}
