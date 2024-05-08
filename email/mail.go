package email

import (
	"context"
	"net/mail"
)

const (
	mailHost = ""
	mailPort = 465
	mailPass = ""
)

type Mail interface {
	SendMail(ctx context.Context, address []string, subject string, body string) bool
}

var defSender = ""

func InitMailGoMail() *MailGoMail {
	return &MailGoMail{}
}

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
