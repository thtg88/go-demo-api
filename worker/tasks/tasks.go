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

var (
	ContactEmailTask = taskq.RegisterTask(&taskq.TaskOptions{
		Name: "contactRequestsEmailer",
		Handler: func(e *email.Email) error {
			SendContatEmail(e)
			time.Sleep(time.Millisecond)
			return nil
		},
	})
)

var counter int32

func SendContatEmail(e *email.Email) {
	mail.SendMail(e)
}

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
