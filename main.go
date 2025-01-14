package main

import (
	"bytes"
	emailtemplate "email-automatizado/emailTemplate"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"gopkg.in/gomail.v2"
)

func main() {
	host := "sandbox.smtp.mailtrap.io"
	port := 587
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	openAIKey := os.Getenv("OPENAI_API_KEY")

	if username == "" || password == "" || openAIKey == "" {
		log.Fatal("USERNAME, PASSWORD ou OPENAI_API_KEY não configuradas corretamente.")
	}

	fromEmail := "from@example.com"
	toEmail := "to@example.com"

	log.Printf("SMTP Host: %s, Port: %d, Username: %s, From: %s, To: %s", host, port, username, fromEmail, toEmail)

	insights, err := generateInsights(openAIKey)
	if err != nil {
		log.Fatalf("Erro ao gerar insights: %v", err)
	}

	data := emailtemplate.EmailData{
		Title:      "Seja bem-vindo ao To Do List!",
		Message:    fmt.Sprintf("Esperamos que você goste da experiência com nosso gerenciador de tarefas queridinho. Aqui estão algumas dicas para aproveitar melhor: %s", insights),
		ImageURL:   "email-automatizado/assets/image.svg",
		FooterNote: "Garanto que não viverá mais sem ele rs!",
	}

	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Boas vindas ao To Do List")
	m.AddAlternative("text/html", emailtemplate.GetHTMLBody(data))

	d := gomail.NewDialer(host, port, username, password)

	err = sendEmailWithRetry(d, m, 3)
	if err != nil {
		log.Fatalf("Erro ao enviar o e-mail após várias tentativas: %v", err)
	}

	log.Println("E-mail enviado com sucesso!")
}

func sendEmailWithRetry(d *gomail.Dialer, m *gomail.Message, maxRetries int) error {
	var err error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		log.Printf("Tentando enviar e-mail, tentativa %d de %d...", attempt, maxRetries)
		err = d.DialAndSend(m)
		if err == nil {
			return nil
		}

		if strings.Contains(err.Error(), "quota") {
			log.Printf("Erro de quota detectado. Aguardando antes de tentar novamente...")
			time.Sleep(2 * time.Minute)
		} else {
			log.Printf("Erro ao enviar e-mail: %v", err)
			break
		}
	}

	return fmt.Errorf("erro ao enviar e-mail após %d tentativas: %w", maxRetries, err)
}

func generateInsights(apiKey string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	requestBody := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "system", "content": "Você é um assistente especializado em dar dicas sobre produtividade."},
			{"role": "user", "content": "Quais são as melhores práticas para gerenciar tarefas em um aplicativo de lista de tarefas? Forneça dicas úteis e inspiradoras."},
		},
		"max_tokens":  100,
		"temperature": 0.7,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("erro ao serializar o corpo da requisição: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("erro ao criar a requisição: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erro ao enviar a requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("erro na resposta da API OpenAI: %s", string(body))
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("erro ao decodificar a resposta: %w", err)
	}

	choices, ok := response["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("nenhuma resposta retornada pela API OpenAI")
	}

	choice, ok := choices[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("erro ao interpretar a resposta")
	}

	content, ok := choice["message"].(map[string]interface{})["content"].(string)
	if !ok {
		return "", fmt.Errorf("erro ao extrair o conteúdo da resposta")
	}

	return strings.TrimSpace(content), nil
}
