package messages

import "time"

type VerifyOtpReq struct {
	FlowId string `json:"workflow_id"` // WorkflowID
	Token  string `json:"token"`
	Code   string `json:"code"`
	Trace  string `json:"trace"`
}

type VerifiedOtpSignal struct {
	WorkflowID  string     `json:"worflowId"`
	FromAccount string     `json:"fromAccount"`
	ToAccount   string     `json:"toAccount"`
	CRefNum     string     `json:"cRefNum"`
	Amount      float64    `json:"amount"`
	CreatedAt   *time.Time `json:"createdAt"`
	TransferAt  *time.Time `json:"transferAt"`
	TransNo     string     `json:"transNo"`
	Status      int        `json:"status"`
}

type CreateTransactionReq struct {
	CRefNum string `json:"cRefNum"`
}

type CreateTransactionSignal struct {
	WorkflowID  string     `json:"worflowId"`
	FromAccount string     `json:"fromAccount"`
	ToAccount   string     `json:"toAccount"`
	CRefNum     string     `json:"cRefNum"`
	Amount      float64    `json:"amount"`
	CreatedAt   *time.Time `json:"createdAt"`
	TransferAt  *time.Time `json:"transferAt"`
	TransNo     string     `json:"transNo"`
	Status      int        `json:"status"`
}
