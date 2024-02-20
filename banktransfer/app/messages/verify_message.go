package messages

type VerifyOtp struct {
	Token string `json:"token"`
	Code  string `json:"code"`
	// NextFlow string `json:"next_flow"`
	Payload string `json:"payload"`
}
