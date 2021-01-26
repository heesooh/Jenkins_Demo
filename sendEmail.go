package main

import (
	"gopkg.in/gomail.v2"
)

func sendEmail(filename string, errorMessage string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "example@gmail.com")
	m.SetHeader("To", "example@gmail.com", "example@gmail.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Jenkins_Demo:")
	//message body need to contain failed messages
	m.SetBody("text/html", "<p>" + errorMessage + "</p>")
	m.Attach(filename)

	d := gomail.NewDialer("smtp.gmail.com", 587, "example@gmail.com", "password")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}