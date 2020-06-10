package yoemail

import (
	"context"
	"errors"
	"fmt"
	"github.com/juandunbar/yobirthday/yoconfig"
	"github.com/mailgun/mailgun-go/v4"
	"time"
)

var (
	config *yoconfig.Config
	ErrConfigNotSet = errors.New("nil config pointer, call SetConfig() to initialize config")
)

type EmailClient struct {
	mg *mailgun.MailgunImpl
}

func NewClient() (*EmailClient, error) {
	if config == nil {
		return nil, ErrConfigNotSet
	}
	// Get apikey from our env variables
	apikey := config.GetString("MAILGUN_APIKEY")
	// Get domain from our config.json file
	domain := config.GetString("mailgun.domain")
	mg := mailgun.NewMailgun(domain, apikey)

	return &EmailClient{mg:mg}, nil
}

func (ec *EmailClient) SendEmail() error {
	sender := config.GetString("mailgun.sender")
	subject := "test subject"
	body := "happy birthday"
	recipient := config.GetString("mailgun.recipient")

	message := ec.mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Send the message with a 10 second timeout
	resp, id, err := ec.mg.Send(ctx, message)
	if err != nil {
		return err
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}

func SetConfig(c *yoconfig.Config) {
	config = c
}
