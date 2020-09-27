package mail

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

func GetPlainAuth(host string, username string, password string) smtp.Auth {
	return smtp.PlainAuth("", username, password, host)
}

func SendMail(e *email.Email) error {
	auth := GetPlainAuth(
		os.Getenv("MAIL_HOST"),
		os.Getenv("MAIL_USERNAME"),
		os.Getenv("MAIL_PASSWORD"),
	)

	e.From = fmt.Sprintf(
		"%s <%s>",
		os.Getenv("MAIL_FROM_NAME"),
		os.Getenv("MAIL_FROM_ADDRESS"),
	)

	return e.SendWithStartTLS(
		fmt.Sprintf("%s:%s", os.Getenv("MAIL_HOST"), os.Getenv("MAIL_PORT")),
		auth,
		&tls.Config{
			InsecureSkipVerify: true,
		},
	)
}
