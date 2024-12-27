package emailtemplate

// GetHTMLBody retorna o corpo HTML do email
func GetHTMLBody() string {
	return `
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
	`
}
