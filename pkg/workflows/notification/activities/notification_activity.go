package activities

import (
	"context"
	"fmt"
	"net/smtp"

	"github.com/anhgeeky/go-temporal-labs/notification/messages"
	"github.com/anhgeeky/go-temporal-labs/notification/outbound/notification"
	"github.com/google/uuid"
	"go.temporal.io/sdk/activity"
)

type NotificationActivity struct {
	NotificationService notification.NotificationService
}

func (a *NotificationActivity) GetDeviceToken(ctx context.Context, msg interface{}) (*messages.DeviceToken, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("NotificationActivity: GetDeviceToken", msg)

	token := messages.DeviceToken{
		FirebaseToken: uuid.New().String(),
	}

	return &token, nil
}

func (a *NotificationActivity) PushEmail(ctx context.Context, msg *messages.DeviceToken) (string, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("NotificationActivity: PushEmail", msg)

	from := "anhnguyen.sogo@gmail.com"
	password := "oesb wira pygw ncqe" // test only

	to := []string{
		"anhgeeky@gmail.com",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("This is a test email message.")
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return "Failed", err
	}
	fmt.Println("Email Sent!")

	return "OK", nil
}

func (a *NotificationActivity) PushSMS(ctx context.Context, msg *messages.DeviceToken) (string, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("NotificationActivity: PushSMS", msg)

	return "OK", nil
}

func (a *NotificationActivity) PushNotification(ctx context.Context, msg *messages.DeviceToken) (string, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("NotificationActivity: PushNotification", msg)

	return "OK", nil
}

func (a *NotificationActivity) PushInternalApp(ctx context.Context, msg *messages.DeviceToken) (string, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("NotificationActivity: PushInternalApp", msg)

	return "OK", nil
}
