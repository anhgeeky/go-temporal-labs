package messages

type VerifyOtpReq struct {
	FlowId string `json:"workflow_id"` // WorkflowID
	Token  string `json:"token"`
	Code   string `json:"code"`
	Trace  string `json:"trace"`
}

type VerifiedOtpSignal struct {
	Item VerifyOtpReq
}

type CreateTransactionReq struct {
	FlowId string `json:"workflow_id"` // WorkflowID
	// TODO: Sơn bổ sung Data Response giúp anh -> Gửi email ra
}

type CreateTransactionSignal struct {
	Item CreateTransactionReq
}
