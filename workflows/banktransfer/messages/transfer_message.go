package messages

import "time"

type TransferMessage struct {
	// Id          string     `json:"id"`
	WorkflowID  string     `json:"worflowId"`
	FromAccount string     `json:"fromAccount"`
	ToAccount   string     `json:"toAccount"`
	CRefNum     string     `json:"cRefNum"`
	Amount      float64    `json:"amount"`
	CreatedAt   *time.Time `json:"createdAt"`
}

type TransferReq struct {
	FromAccount string  `json:"fromAccount"`
	ToAccount   string  `json:"toAccount"`
	Amount      float64 `json:"amount"`
}

type TransferResult struct {
	Id                           string  `json:"id"`
	OldAccountOriginBalance      float64 `json:"old_account_origin_balance"`
	NewAccountOriginBalance      float64 `json:"new_account_origin_balance"`
	OldAccountDestinationBalance float64 `json:"old_account_destination_balance"`
	NewAccountDestinationBalance float64 `json:"new_account_destination_balance"`
}
