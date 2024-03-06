package messages

type NotificationMessage struct {
	Token DeviceToken
}

type DeviceToken struct {
	FirebaseToken string `json:"firebase_token"`
}
