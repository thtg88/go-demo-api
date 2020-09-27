package tasks

import (
	"goDemoApi/mail"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jordan-wright/email"
	"github.com/vmihailenco/taskq/v3"
)

// ContactEmailTask is a taskq task that sends given emails
var ContactEmailTask = taskq.RegisterTask(&taskq.TaskOptions{
	Name: "contactRequestsEmailer",
	Handler: func(e *email.Email) error {
		SendContactEmail(e)
		time.Sleep(time.Millisecond)
		return nil
	},
})

// SendContactEmail sends the given email
func SendContactEmail(e *email.Email) {
	mail.SendMail(e)
}

// WaitSignal waits for a term signal to stop the worker
func WaitSignal() os.Signal {
	ch := make(chan os.Signal, 2)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
			return sig
		}
	}
}
