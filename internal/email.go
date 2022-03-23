package internal

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"

	"github.com/erdincmutlu/newsapi/types"
)

func sendEmail(request types.ShareRequest) error {
	message, err := prepareMessage(request.ID)
	if err != nil {
		return err
	}

	from := "ierdincmutlu@gmail.com"
	to := []string{request.Recipient}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, emailPass, smtpHost)

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
}

func prepareMessage(itemID string) ([]byte, error) {
	t, err := template.ParseFiles("internal/templates/email.html")
	if err != nil {
		return nil, err
	}

	newsItem, err := getItem(itemID)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, newsItem)
	if err != nil {
		log.Printf("Error in t.Execute: %s\n", err.Error())
		return nil, err
	}

	return buf.Bytes(), nil
}
