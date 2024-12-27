# 🚀 Automated Email Sending with Mailtrap using Go

This project demonstrates how to send automated emails using Mailtrap and the `gomail` library in Go. The code sets up an SMTP server and sends an email in both text and HTML formats to a recipient's address.

## 📝 **Glossary**
- [📋 Prerequisites](#prerequisites)
- [🛠️ Installation](#installation)
- [🚀 Environment Variables](#environment-variables)
- [⚙️ Go Code](#go-code)
- [📁 Code Structure](#code-structure)
- [✅ Execution](#execution)

---

## 📋 Prerequisites <a name="prerequisites"></a>

Before running the project, make sure the following are installed:

- [Go](https://golang.org/dl/) (version 1.18 or higher)
- [Mailtrap](https://mailtrap.io/) (for obtaining SMTP credentials)

## 🛠️ Installation <a name="installation"></a>

1. Clone this repository:

    ```bash
    git clone https://github.com/YOUR-USER/email-automatizado.git
    cd email-automatizado
    ```

2. Install the required dependencies:

    - Install the `gomail` library for sending emails:

    ```bash
    go get gopkg.in/gomail.v2
    ```

3. Adjust variables as needed:

- host := "The Mailtrap host"
- port := "The Mailtrap port (587 is the most usable)"
- username := "The Mailtrap username"
- password := "The password generated in Mailtra"

## 🚀 Environment Variables <a name="environment-variables"></a>

The code uses the `USERNAME` and `PASSWORD` environment variables to configure SMTP authentication with Mailtrap. These variables should be defined in the terminal or in a `.env` file at the root of the project.

- `USERNAME`: The Mailtrap username (available in your Mailtrap account).
- `PASSWORD`: The password generated in Mailtrap (available in your Mailtrap account).

If you prefer, you can set the variables directly in the terminal (as I did):

```bash
export USERNAME="your_username"
export PASSWORD="your_password"
```

## ⚙️ Go Code <a name="go-code"></a>
The Go code uses the gomail library to send an email in both text and HTML formats. The SMTP server configuration is done using the environment variables for authentication.

### 📁 Code Structure <a name="code-structure"></a>

- main.go: The main file that contains the code to send the email.
- emailtemplate/template.go: The html template for the email.
- .env: (optional) File to store the USERNAME and PASSWORD environment variables.

## ✅ Execution <a name="execution"></a>
After setting up the environment variables or the .env file, run the code:

```bash
go run main.go
```

## Example Output
If the email is sent successfully, you will see the following output in the terminal:

```bash
SMTP Host: sandbox.smtp.mailtrap.io, Port: 587, Username: your_username, From: from@example.com, To: to@example.com
Attempting to send the email...
Email sent successfully!
```
If there’s an error while sending the email, the error will be displayed, and you can investigate it based on the returned message.