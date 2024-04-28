package email

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
)

type MailStd struct{}

func (m *MailStd) SendMail(ctx context.Context, address []string, subject string, body string) bool {
	host := mailHost
	port := mailPort
	password := mailPass

	header := make(map[string]string)
	header["From"] = defSender
	header["Subject"] = subject
	header["Content-Type"] = "text/html; charset=UTF-8"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	auth := smtp.PlainAuth("", defSender, password, host)
	if err := sendMailUsingTLS(fmt.Sprintf("%s:%d", host, port), auth, defSender, address, []byte(message)); err != nil {
		return false
	}
	return true
}

func dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		//log.Println("Dialing Error:", err)
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func sendMailUsingTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
	c, err := dial(addr)
	if err != nil {
		//log.Println("Create smpt client error:", err)
		return err
	}
	defer func(c *smtp.Client) {
		if err = c.Close(); err != nil {
			//log.Print("close err:", err)
		}
	}(c)

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				//log.Println("Error during AUTH", err)
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr = range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}
