package account

type CheckBalanceReq struct {
	WorkflowID string `json:"workflow_id"`
	Action     string `json:"action"`
}
