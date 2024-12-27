package main

import (
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

func main() {
	host := "sandbox.smtp.mailtrap.io"
	port := 587
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	if username == "" || password == "" {
		log.Fatal("The USERNAME or PASSWORD environment variables are not set correctly.")
	}

	fromEmail := "from@example.com"
	toEmail := "to@example.com"

	log.Printf("SMTP Host: %s, Port: %d, Username: %s, From: %s, To: %s", host, port, username, fromEmail, toEmail)

	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "You are awesome!")
	m.SetBody("text/plain", "Congrats for sending test email with Mailtrap!")

	m.AddAlternative("text/html", `
		<!doctype html>
		<html>
		  <head>
		    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		  </head>
		  <body style="font-family: sans-serif;">
		    <h1>Congrats for sending test email with Mailtrap!</h1>
		    <p>If you are viewing this email in your inbox – the integration works.</p>
		    <p>Good luck! Hope it works.</p>
		  </body>
		</html>
	`)

	d := gomail.NewDialer(host, port, username, password)

	log.Println("Trying to send the email...")
	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("Error sending email: %v", err)
	}

	log.Println("Email sent successfully!")
}
