package main

import (
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

func main() {
	// Configurações do Mailtrap
	host := "sandbox.smtp.mailtrap.io"
	port := 587
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	// Verificando as variáveis de ambiente
	if username == "" || password == "" {
		log.Fatal("As variáveis de ambiente USERNAME ou PASSWORD não estão definidas corretamente.")
	}

	fromEmail := "from@example.com" // Certifique-se de que este é um e-mail válido
	toEmail := "to@example.com"     // E-mail de destino

	// Logs de depuração
	log.Printf("SMTP Host: %s, Port: %d, Username: %s, From: %s, To: %s", host, port, username, fromEmail, toEmail)

	// Criando a mensagem de e-mail
	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "You are awesome!")
	m.SetBody("text/plain", "Congrats for sending test email with Mailtrap!")

	// Corpo alternativo em HTML
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

	// Criando o dialer SMTP
	d := gomail.NewDialer(host, port, username, password)

	// Tentando enviar o e-mail
	log.Println("Tentando enviar o e-mail...")
	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("Erro ao enviar e-mail: %v", err)
	}

	log.Println("E-mail enviado com sucesso!")
}
