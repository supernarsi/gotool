package email

import (
	"context"
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type MailGoMail struct{}

func (m *MailGoMail) SendMail(ctx context.Context, address []string, subject string, body string) bool {
	gm := gomail.NewMessage()

	// Set E-Mail defSender
	gm.SetHeader("From", defSender)

	// Set E-Mail receivers
	gm.SetHeader("To", address...)

	// Set E-Mail subject
	gm.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	gm.SetBody("text/html", body)

	// Settings for SMTP server
	d := gomail.NewDialer(mailHost, mailPort, defSender, mailPass)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(gm); err != nil {
		return false
	}

	return true
}
