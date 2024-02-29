package account

type AccountList struct {
	Accounts []Account `json:"accounts"`
}

type BalanceRes struct {
	Balance float64 `json:"balance"`
}
