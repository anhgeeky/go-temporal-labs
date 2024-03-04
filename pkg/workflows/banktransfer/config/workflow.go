package config

var TaskQueues = struct {
	TRANSFER_QUEUE string
}{
	TRANSFER_QUEUE: "TRANSFER_QUEUE",
}

var Workflows = struct {
	TRANSFER string
}{
	TRANSFER: "TRANSFER",
}

var Messages = struct {
	GROUP                            string
	CHECK_BALANCE_REQUEST_TOPIC      string
	CHECK_BALANCE_REPLY_TOPIC        string
	CREATE_TRANSACTION_REQUEST_TOPIC string
	CREATE_TRANSACTION_REPLY_TOPIC   string
}{
	GROUP:                            "go_clean",
	CHECK_BALANCE_REQUEST_TOPIC:      "check-balance-request-topic",      // TODO: Check với Sơn
	CHECK_BALANCE_REPLY_TOPIC:        "check-balance-reply-topic",        // TODO: Check với Sơn
	CREATE_TRANSACTION_REQUEST_TOPIC: "create-transaction-request-topic", // TODO: Check với Sơn
	CREATE_TRANSACTION_REPLY_TOPIC:   "create-transaction-reply-topic",   // TODO: Check với Sơn
}
