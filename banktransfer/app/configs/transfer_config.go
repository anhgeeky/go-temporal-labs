package configs

var Workflows = struct {
	BANK_TRANSFER string
}{
	BANK_TRANSFER: "BANK_TRANSFER",
}

var SignalChannels = struct {
	ADD_TO_TRANSFER_CHANNEL      string
	REMOVE_FROM_TRANSFER_CHANNEL string
	UPDATE_EMAIL_CHANNEL         string
	CHECKOUT_CHANNEL             string
}{
	ADD_TO_TRANSFER_CHANNEL:      "ADD_TO_TRANSFER_CHANNEL",
	REMOVE_FROM_TRANSFER_CHANNEL: "REMOVE_FROM_TRANSFER_CHANNEL",
	UPDATE_EMAIL_CHANNEL:         "UPDATE_EMAIL_CHANNEL",
	CHECKOUT_CHANNEL:             "CHECKOUT_CHANNEL",
}

var RouteTypes = struct {
	ADD_TO_TRANSFER      string
	REMOVE_FROM_TRANSFER string
	UPDATE_EMAIL         string
	CHECKOUT             string
}{
	ADD_TO_TRANSFER:      "add_to_msg",
	REMOVE_FROM_TRANSFER: "remove_from_msg",
	UPDATE_EMAIL:         "update_email",
	CHECKOUT:             "checkout",
}
