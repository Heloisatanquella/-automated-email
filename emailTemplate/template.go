package emailtemplate

import (
	"bytes"
	"html/template"
	"log"
)

type EmailData struct {
	Title      string
	Message    string
	ImageURL   string
	FooterNote string
}

// GetHTMLBody retorna o corpo HTML do email
func GetHTMLBody(data EmailData) string {
	const htmlTemplate = `
		<!doctype html>
		<html>
		  <head>
		    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		  </head>
		  <body style="font-family: sans-serif;">
		    <div style="margin: auto; max-width: 600px;">
		      <h1>{{.Title}}</h1>
		      <p>{{.Message}}</p>
		      {{if .ImageURL}}
		      <img src="{{.ImageURL}}" alt="Email Image" style="width: 100%; height: auto;">
		      {{end}}
		      <p>{{.FooterNote}}</p>
		    </div>
		  </body>
		</html>
	`

	tmpl, err := template.New("email").Parse(htmlTemplate)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	return buf.String()
}
