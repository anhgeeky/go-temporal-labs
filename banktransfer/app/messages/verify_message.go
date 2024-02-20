package messages

type VerifyOtpMessage struct {
	Token   string `json:"token"`
	Code    string `json:"code"`
	Payload string `json:"payload"`
}

type VerifyOtpReq struct {
	Token   string `json:"token"`
	Code    string `json:"code"`
	Payload string `json:"payload"`
	Trace   string `json:"trace"`
}
