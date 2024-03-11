package config

var TaskQueues = struct {
	NOTIFICATION_QUEUE string
}{
	NOTIFICATION_QUEUE: "NOTIFICATION_QUEUE",
}

var Workflows = struct {
	NotificationName string
}{
	NotificationName: "NotificationWorkflow",
}
