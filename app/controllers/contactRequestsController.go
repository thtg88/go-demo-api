package controllers

import (
	"context"
	"goDemoApi/cmd/worker/tasks"
	"goDemoApi/queue"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
)

type contactRequestsStoreResponse struct {
	Success bool
}

// ContactRequestsStore puts an email message on the main queue
// to be sent at a later date
func ContactRequestsStore(c *gin.Context) {
	e := &email.Email{
		To:      []string{os.Getenv("MAIL_INTERNAL_NOTIFICATION_ADDRESS")},
		Subject: os.Getenv("MAIL_INTERNAL_SUBJECT"),
		Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte("<h1>Text Body is, of course, supported!</h1>"),
	}
	msg := tasks.ContactEmailTask.WithArgs(context.Background(), e)
	msg.Delay = time.Minute
	err := queue.AddToMainQueue(msg)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, &contactRequestsStoreResponse{
		Success: true,
	})
}
