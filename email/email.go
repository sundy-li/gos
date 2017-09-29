package email

import (
	"net/smtp"
	"strings"
)

type EmailService interface {
	Send(to, subject, msg, mailtype string) error
}

type Email struct {
	host     string
	user     string
	password string

	auth smtp.Auth
}

func NewEmail(host, user, password string, ssl bool) EmailService {
	if !ssl {
		email := &Email{host: host, user: user, password: password}
		hp := strings.Split(host, ":")
		email.auth = smtp.PlainAuth("", user, password, hp[0])
		return email
	} else {
		return NewSslEmail(host, user, password)
	}
}

func (email *Email) Send(to, subject, msg, mailtype string) error {
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	content := []byte("To: " + to + "\r\nFrom: " + email.user + "<" + email.user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + msg)
	send_to := strings.Split(to, ";")
	return smtp.SendMail(email.host, email.auth, email.user, send_to, content)
}
