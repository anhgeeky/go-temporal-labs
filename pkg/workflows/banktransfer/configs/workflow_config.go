package configs

var TaskQueues = struct {
	BANK_TRANSFER string
}{
	BANK_TRANSFER: "BANK_TRANSFER",
}

var Workflows = struct {
	VERIFY       string
	TRANSFER     string
	NOTIFICATION string
}{
	VERIFY:       "VERIFY",
	TRANSFER:     "TRANSFER",
	NOTIFICATION: "NOTIFICATION",
}
