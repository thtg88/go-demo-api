package controllers

import (
	"context"
	"goDemoApi/internal/worker/queue"
	"goDemoApi/internal/worker/tasks"
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
	// rules := govalidator.MapData{
	// 	"username": []string{"required", "between:3,8"},
	// 	"email":    []string{"required", "min:4", "max:20", "email"},
	// 	"web":      []string{"url"},
	// 	"phone":    []string{"digits:11"},
	// 	"agree":    []string{"bool"},
	// 	"dob":      []string{"date"},
	// }

	// opts := govalidator.Options{
	// 	Request:         c.Request, // request object
	// 	Rules:           rules,     // rules map
	// 	RequiredDefault: true,      // all the field to be pass the rules
	// }
	// v := govalidator.New(opts)
	// errors := v.Validate()
	// if len(errors) > 0 {
	// 	validationError := map[string]interface{}{"validationError": errors}
	// 	panic(validationError)
	// }

	e := &email.Email{
		To:      []string{os.Getenv("CONTACT_REQUEST_MAIL_INTERNAL_NOTIFICATION_ADDRESS")},
		Subject: os.Getenv("CONTACT_REQUEST_MAIL_INTERNAL_SUBJECT"),
		Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte("<h1>Text Body is, of course, supported!</h1>"),
	}
	err := queue.AddEmailToMainQueue(e)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, &contactRequestsStoreResponse{
		Success: true,
	})
}
