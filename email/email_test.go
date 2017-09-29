package email

import (
	"testing"
)

func TestSendEmail(t *testing.T) {
	host := "smtp.exmail.qq.com:25"
	user := "boomer_kefu@truxing.com"
	password := "NyxQGrXTpjiad18H"

	email := NewEmail(host, user, password, false)
	err := email.Send("sundyli@truxing.com", "this is an email", "just for test2", "")
	if err != nil {
		panic(err)
	}
}

func TestSendSslEmail(t *testing.T) {
	host := "smtp.exmail.qq.com:465"
	user := "boomer_kefu@truxing.com"
	password := "NyxQGrXTpjiad18H"

	email := NewEmail(host, user, password, true)
	err := email.Send("sundyli@truxing.com", "this is an ssl email", "just for test2", "")
	if err != nil {
		panic(err)
	}
}
