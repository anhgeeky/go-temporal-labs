package activities_test

import (
	"fmt"
	"net/smtp"
	"testing"
)

func Test_LoadConfig(t *testing.T) {
	from := "anhnguyen.sogo@gmail.com"
	password := "oesb wira pygw ncqe"

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
	}
	fmt.Println("Email Sent!")
}
