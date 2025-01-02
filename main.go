package main

import (
	emailtemplate "email-automatizado/emailTemplate"
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
	m.SetHeader("Subject", "Boas vindas ao To Do List")

	data := emailtemplate.EmailData{
		Title:      "Seja bem-vindo ao To Do List!",
		Message:    "Esperamos que você goste da experiência com nosso gerenciador de tarefas queridinho.",
		ImageURL:   "email-automatizado/assets/image.svg",
		FooterNote: "Garanto que não viverá mais sem ele rs!",
	}

	m.AddAlternative("text/html", emailtemplate.GetHTMLBody(data))

	d := gomail.NewDialer(host, port, username, password)

	log.Println("Trying to send the email...")
	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("Error sending email: %v", err)
	}

	log.Println("Email sent successfully!")
}
