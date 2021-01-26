package main

import (
	//"gopkg.in/gomail.v2"
	"var/lib/jenkins/go/src/gopkg.in/gomail.v2"
)

func sendEmail(filename string, errorMessage string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "blockchainwarning@omnisolu.com")
	m.SetHeader("To", "folowal757@loopsnow.com", "example@gmail.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Jenkins_Demo:")
	m.SetBody("text/html", "<p>" + errorMessage + "</p>")
	m.Attach(filename)

	d := gomail.NewDialer("smtp.gmail.com", 587, "blockchainwarning@omnisolu.com", "01353751")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}