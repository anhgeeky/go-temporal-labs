package activities

import (
	"context"
	"fmt"

	"github.com/anhgeeky/go-temporal-labs/notification/config"
	"github.com/anhgeeky/go-temporal-labs/notification/messages"
	"go.temporal.io/sdk/activity"
	gomail "gopkg.in/mail.v2"
)

type NotificationActivity struct {
	Config *config.EmailConfig
}

func (a *NotificationActivity) PushEmail(ctx context.Context, msg *messages.NotificationMessage) (string, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("NotificationActivity: PushEmail", msg)

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", a.Config.EmailFrom)

	// Set E-Mail receivers
	m.SetHeader("To", a.Config.EmailTo)

	// Set E-Mail subject
	m.SetHeader("Subject", "Bank Transfer Completed")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "Transfed done")

	// Settings for SMTP server
	d := gomail.NewDialer(a.Config.SmtpHost, a.Config.SmtpPort, a.Config.SmtpAccount, a.Config.SmtpPassword)

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
