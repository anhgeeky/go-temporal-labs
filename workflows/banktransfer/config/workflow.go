package config

var (
	VERSION_1_0 = "1.0"
	VERSION_2_0 = "2.0"
	VERSION_3_0 = "3.0"
	VERSION_4_0 = "4.0"
)

var TaskQueues = struct {
	TRANSFER_QUEUE string
}{
	TRANSFER_QUEUE: "TRANSFER_QUEUE",
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
	CHECK_BALANCE_ACTION:        "check-balance",               // => activityID
	CHECK_BALANCE_REQUEST_TOPIC: "check-balance-request-topic", // TODO: Check với Sơn
	CHECK_BALANCE_REPLY_TOPIC:   "check-balance-reply-topic",   // TODO: Check với Sơn

	CREATE_TRANSACTION_ACTION:        "create-transaction",               // => activityID
	CREATE_TRANSACTION_REQUEST_TOPIC: "create-transaction-request-topic", // TODO: Check với Sơn
	CREATE_TRANSACTION_REPLY_TOPIC:   "create-transaction-reply-topic",   // TODO: Check với Sơn

	CREATE_OTP_ACTION:        "create-otp",               // => activityID
	CREATE_OTP_REQUEST_TOPIC: "create-otp-request-topic", // TODO: Check với Sơn
	CREATE_OTP_REPLY_TOPIC:   "create-otp-reply-topic",   // TODO: Check với Sơn
}
