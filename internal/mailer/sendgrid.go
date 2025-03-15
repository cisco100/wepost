package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridMailer struct {
	From   string
	ApiKey string
	Client *sendgrid.Client
}

func NewSendGridMailer(from, apiKey string) *SendGridMailer {
	client := sendgrid.NewSendClient(apiKey)
	return &SendGridMailer{
		From:   from,
		ApiKey: apiKey,
		Client: client,
	}
}

func (msgd *SendGridMailer) Send(templateFile, username, email string, data any, isSandbox bool) error {

	temp, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return err
	}

	from := mail.NewEmail(FromName, msgd.From)
	to := mail.NewEmail(username, email)

	subject := new(bytes.Buffer)
	err = temp.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}
	body := new(bytes.Buffer)

	err = temp.ExecuteTemplate(body, "body", data)
	if err != nil {
		return err
	}
	message := mail.NewSingleEmail(from, subject.String(), to, "", body.String())
	message.SetMailSettings(
		&mail.MailSettings{
			SandboxMode: &mail.Setting{
				Enable: &isSandbox,
			},
		},
	)

	for i := 0; i <= MaxRetriesLimit; i++ {
		response, err := msgd.Client.Send(message)
		if err != nil {
			log.Printf("Failed to send email to %s, attempt %d out of %d/n", email, i+1, MaxRetriesLimit)
			log.Println(err)
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}
		log.Printf("Email sent successfully with status code %d", response.StatusCode)
		return nil

	}
	return fmt.Errorf("error sending message after %d tries", MaxRetriesLimit)

}
