package main

import (
	"bufio"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/sundy-li/gos/email"
)

var (
	host     string
	user     string
	password string

	ssl     bool
	subject string
	touser  string
)

// read msgs from stdin
func main() {
	var mailcmd = &cobra.Command{
		Use:   "gomail",
		Short: "gomail",
		Long:  "gomail",
		Run: func(cmd *cobra.Command, args []string) {
			em := email.NewEmail(host, user, password, ssl)
			reader := bufio.NewReader(os.Stdin)
			var text string
			for {
				t, err := reader.ReadString('\n')
				if err == io.EOF {
					break
				}
				text += t
			}
			println("sending => ", text)
			em.Send(touser, subject, text, "")
		},
	}
	mailcmd.Flags().StringVarP(&host, "host", "q", "smtp.exmail.qq.com:25", "stmp hosts")
	mailcmd.Flags().StringVarP(&user, "user", "u", "", "sender user username")
	mailcmd.Flags().StringVarP(&password, "password", "p", "", "sender user pwd")
	mailcmd.Flags().StringVarP(&subject, "subject", "j", "subject", "send email subject")
	mailcmd.Flags().StringVarP(&touser, "touser", "t", "", "to email address")
	mailcmd.Flags().BoolVarP(&ssl, "ssl", "s", false, "ssl email server")

	mailcmd.Execute()
}

func stdinAvailable() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
