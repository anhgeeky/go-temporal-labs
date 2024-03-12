package config

var TaskQueues = struct {
	TRANSFER_QUEUE string
}{
	TRANSFER_QUEUE: "TransferTaskQueue",
}

var Workflows = struct {
	TransferWorkflow string
}{
	TransferWorkflow: "TransferWorkflow",
}

var Messages = struct {
	CHECK_BALANCE_ACTION             string
	CHECK_BALANCE_REQUEST_TOPIC      string
	CHECK_BALANCE_REPLY_TOPIC        string
	CREATE_TRANSACTION_ACTION        string
	CREATE_TRANSACTION_REQUEST_TOPIC string
	CREATE_TRANSACTION_REPLY_TOPIC   string
	CREATE_OTP_ACTION                string
	CREATE_OTP_REQUEST_TOPIC         string
	CREATE_OTP_REPLY_TOPIC           string
}{
	CHECK_BALANCE_ACTION:        "check-balance", // => activityID
	CHECK_BALANCE_REQUEST_TOPIC: "OCB.REQUEST.CHECK_BALANCE",
	CHECK_BALANCE_REPLY_TOPIC:   "OCB.REPLY.CHECK_BALANCE",

	CREATE_TRANSACTION_ACTION:        "create-transaction", // => activityID
	CREATE_TRANSACTION_REQUEST_TOPIC: "OCB.REQUEST.FUND_TRANSFER",
	CREATE_TRANSACTION_REPLY_TOPIC:   "OCB.REPLY.FUND_TRANSFER",

	CREATE_OTP_ACTION:        "create-otp", // => activityID
	CREATE_OTP_REQUEST_TOPIC: "OCB.REQUEST.GENERATE_OTP",
	CREATE_OTP_REPLY_TOPIC:   "OCB.REPLY.FUND_TRANSFER",
}
