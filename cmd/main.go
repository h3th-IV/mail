package main

import (
	"fmt"
	"log"
	"time"

	"github.com/h3th-IV/mail/mail"
)

func main() {
	config := mail.MailerConfig{
		Host:     "live.smtp.mailtrap.io",
		Port:     587,
		Username: "api",
		Password: "36d48c8c389668abc3c28087d34793d1",
		Timeout:  5 * time.Second,
		Sender:   "samuelbonux10@gmail.com",
	}

	sender := mail.New(config)

	err := sender.Send("samuelbonu0@gmail.com", "welcome.html", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent!")
}
