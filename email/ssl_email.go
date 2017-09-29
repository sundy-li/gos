package email

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
)

// SSL/TLS Email Example

type SslEmail struct {
	host string
	user string
	pwd  string
}

func NewSslEmail(host, user, password string) *SslEmail {
	return &SslEmail{
		host: host,
		user: user,
		pwd:  password,
	}
}

func (e *SslEmail) Send(touser, subj, body, mailtype string) error {
	from := mail.Address{"", e.user}
	to := mail.Address{"", touser}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	host, _, _ := net.SplitHostPort(e.host)
	auth := smtp.PlainAuth("", e.user, e.pwd, host)
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	conn, err := tls.Dial("tcp", e.host, tlsconfig)
	if err != nil {
		return err
	}
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		return err
	}

	if err = c.Rcpt(to.Address); err != nil {
		return err
	}
	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}
	c.Quit()
	return nil
}
