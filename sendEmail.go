package main

import (
	"gopkg.in/gomail.v2"
)

func sendEmail(filename string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "example@example.com")
	m.SetHeader("To", "example@example.com", "example@example.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Jenkins_Demo:")
	m.SetBody("text/html", "This is a message for Jenkins Demo written in html format.")
	m.Attach(filename)

	d := gomail.NewDialer("smtp.example.com", 587, "example@example.com", "example")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}