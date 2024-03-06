package messages

import "time"

type Transfer struct {
	Id                   string     `json:"id"`
	WorkflowID           string     `json:"workflow_id"`
	AccountOriginId      string     `json:"account_origin_id"`
	AccountDestinationId string     `json:"account_destination_id"`
	Amount               float64    `json:"amount"`
	CreatedAt            *time.Time `json:"created_at"`
}

type TransferReq struct {
	AccountOriginId      string  `json:"account_origin_id"`
	AccountDestinationId string  `json:"account_destination_id"`
	Amount               float64 `json:"amount"`
}

type TransferResult struct {
	Id                           string  `json:"id"`
	OldAccountOriginBalance      float64 `json:"old_account_origin_balance"`
	NewAccountOriginBalance      float64 `json:"new_account_origin_balance"`
	OldAccountDestinationBalance float64 `json:"old_account_destination_balance"`
	NewAccountDestinationBalance float64 `json:"new_account_destination_balance"`
}
