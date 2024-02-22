package messages

type VerifyOtpReq struct {
	Token  string `json:"token"`
	Code   string `json:"code"`
	FlowId string `json:"flow_id"`
	Trace  string `json:"trace"`
}

type VerifiedOtpSignal struct {
	Route string
	Item  VerifyOtpReq
}
