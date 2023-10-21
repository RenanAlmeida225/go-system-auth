package helper

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(email, token string) error {
	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")
	to := []string{email}
	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")
	address := host + ":" + port
	subject := "Subject: Confirm your email!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf("<html><body><a href=\"%s/api/v1/auth/enable/%s\">Confirm</a></body></html>", os.Getenv("URL"), token)
	message := []byte(subject + mime + body)
	auth := smtp.PlainAuth("Confirm your email", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		return err
	}

	return nil
}
