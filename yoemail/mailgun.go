package yoemail

import (
	"context"
	"github.com/juandunbar/yobirthday/yodata"
	"github.com/mailgun/mailgun-go/v4"
	"log"
	"time"
)

type MailGun struct {
	client *mailgun.MailgunImpl
	ds *yodata.DataService
}

func (mg *MailGun) Init() error {
	// Get apikey from our env variables
	apikey := config.GetString("MAILGUN_APIKEY")
	// Get domain from our config.json file
	domain := config.GetString("mailgun.domain")
	mg.client = mailgun.NewMailgun(domain, apikey)
	// Get our data service object
	ds, err := yodata.NewService()
	if err != nil {
		return err
	}
	mg.ds = ds
	return nil
}

func (mg *MailGun) SendEmails() error {
	var err error
	// Grab any birthday emails we need to send
	emails, err := mg.ds.GetEmails()
	if err != nil {
		return err
	}
	// Grab any default values from our config values
	sender := config.GetString("mailgun.sender")
	subject := config.GetString("mailgun.subject")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Loop through our birthday emails and send them
	// TODO should this be its own routine?
	for _, email := range emails {
		// Get a canned response from one of our templates
		body, err := GetEmailBody(&email)
		if err != nil {
			break
		}
		message := mg.client.NewMessage(sender, subject, body, email.Email)
		message.SetDKIM(true)
		// Send the message with a 10 second timeout
		resp, id, err := mg.client.Send(ctx, message)
		if err != nil {
			break
		}
		log.Printf("ID: %s Resp: %s\n", id, resp)
	}
	return err
}