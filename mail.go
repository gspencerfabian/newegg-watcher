package main

import (
	"log"
	"net/smtp"
	"strconv"
)

func sendMail(title, url, price string, total, limit int) {
	from := config.Email.Sender.Address
	pass := config.Email.Sender.Password
	to := config.Email.Receiver.Address

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: NEWEGG-WATCHER | IN STOCK!\n\n" +
		"Title: " + title + "\n" +
		"Price: " + price + "\n" +
		"Limit: " + strconv.Itoa(limit) + "\n" +
		"Total: " + strconv.Itoa(total) + "\n" +
		"Url: " + url + "\n\n\n\n\n" +
		"sent using https://github.com/gspencerfabian/newegg-watcher"

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("Email smtp error: %s", err)
	} else {
		log.Println("Email sent successfully.")
	}

	return
}
