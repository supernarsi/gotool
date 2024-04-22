package email

import "context"

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
