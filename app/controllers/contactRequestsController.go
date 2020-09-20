package controllers

import (
	"goDemoApi/mail"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jordan-wright/email"
)

type ContactRequestsStoreResponse struct {
	Success bool
}

func ContactRequestsStore(c *fiber.Ctx) error {
	e := &email.Email{
		To:      []string{os.Getenv("MAIL_INTERNAL_NOTIFICATION_ADDRESS")},
		Subject: os.Getenv("MAIL_INTERNAL_SUBJECT"),
		Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte("<h1>Text Body is, of course, supported!</h1>"),
	}

	err := mail.SendMail(e)
	if err != nil {
		panic(err)
	}

	return c.JSON(&ContactRequestsStoreResponse{
		Success: true,
	})
}
