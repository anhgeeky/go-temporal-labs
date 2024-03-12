package account

type CheckBalanceReq struct {
	// TODO: Check với Sơn response
}

type CheckBalanceRes struct {
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

type CreateTransactionReq struct {
	CRefNum string `json:"cRefNum"`
}

type CreateTransactionRes struct {
	// TODO: Check với Sơn response
}

type CreateOTPReq struct {
	CRefNum string `json:"cRefNum"`
}

type CreateOTPRes struct {
}
