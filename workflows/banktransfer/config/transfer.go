package config

var SignalChannels = struct {
	VERIFY_OTP_CHANNEL         string
	CREATE_TRANSACTION_CHANNEL string
}{
	VERIFY_OTP_CHANNEL:         "VERIFY_OTP_CHANNEL",
	CREATE_TRANSACTION_CHANNEL: "CREATE_TRANSACTION_CHANNEL",
}
