package activities_test

import (
	"fmt"
	"testing"

	gomail "gopkg.in/mail.v2"
)

func Test_SendEmail(t *testing.T) {

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
}
