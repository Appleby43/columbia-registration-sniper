package email

import (
	"log"
	"net/smtp"
	"os"
)

const smtpHost = "smtp.gmail.com:587"

func Send(body string) {
	addressFrom := os.Getenv("SNIPER_ADDR_FROM")
	addressTo := os.Getenv("SNIPER_ADDR_TO")
	password := os.Getenv("SNIPER_PASSWORD")

	rawBody := []byte(body)

	auth := smtp.PlainAuth("", addressFrom, password, "smtp.gmail.com")

	err := smtp.SendMail(smtpHost, auth, addressFrom, []string{addressTo}, rawBody)

	if err != nil {
		log.Fatal(err)
	}
}