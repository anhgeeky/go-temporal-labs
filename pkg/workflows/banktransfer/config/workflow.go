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
	GROUP         string
	REQUEST_TOPIC string
	REPLY_TOPIC   string
}{
	GROUP:         "go_clean",
	REQUEST_TOPIC: "request-topic",
	REPLY_TOPIC:   "reply-topic",
}
