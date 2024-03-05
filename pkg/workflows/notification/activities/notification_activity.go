package activities

import (
	"context"
	"fmt"

	"github.com/anhgeeky/go-temporal-labs/notification/messages"
	"github.com/anhgeeky/go-temporal-labs/notification/outbound/notification"
	"github.com/google/uuid"
	"go.temporal.io/sdk/activity"
	gomail "gopkg.in/mail.v2"
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
	to := "anhgeeky@gmail.com"
	password := "oesb wira pygw ncqe" // test only

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", from)

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	m.SetHeader("Subject", "Bank Transfer Completed")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "Transfed done")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, from, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
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
