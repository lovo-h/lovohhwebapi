package infrastructure

import (
	"net/smtp"
	"os"
)

type Gmailer struct {
	un, pw string
	name   string
	host   string
	port   string
}

func (mailer *Gmailer) Send(from, to, subject, msg string) error {
	formattedMsg := "From: " + mailer.name + "<" + mailer.un + ">\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		"From: " + from + "\n\n" + msg

	auth := smtp.PlainAuth("", mailer.un, mailer.pw, mailer.host)
	sendErr := smtp.SendMail(mailer.host+":"+mailer.port, auth, from, []string{to}, []byte(formattedMsg))

	return sendErr
}

func GetAndInitGMailer() *Gmailer {
	gmailer := new(Gmailer)

	emailUN := os.Getenv("LWA_EMAIL_UN")
	emailPW := os.Getenv("LWA_EMAIL_PW")

	gmailer.name = "LovoHH API"
	gmailer.un = emailUN
	gmailer.pw = emailPW
	gmailer.host = "smtp.gmail.com"
	gmailer.port = "587"

	return gmailer
}
