package yoemail

import (
	"context"
	"fmt"
	"github.com/juandunbar/yobirthday/yoconfig"
	"github.com/mailgun/mailgun-go/v4"
	"time"
)

type MailGun struct {
	client *mailgun.MailgunImpl
}

func (mg *MailGun) Init(config *yoconfig.Config) EmailClient {
	// Get apikey from our env variables
	apikey := config.GetString("MAILGUN_APIKEY")
	// Get domain from our config.json file
	domain := config.GetString("mailgun.domain")
	mg.client = mailgun.NewMailgun(domain, apikey)
	return mg
}

func (mg *MailGun) SendEmail() error {
	sender := config.GetString("mailgun.sender")
	subject := "test subject"
	body := "happy birthday"
	recipient := config.GetString("mailgun.recipient")

	message := mg.client.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Send the message with a 10 second timeout
	resp, id, err := mg.client.Send(ctx, message)
	if err != nil {
		return err
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}