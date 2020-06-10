package yoemail

import (
	"errors"
	"fmt"
	"github.com/juandunbar/yobirthday/yoconfig"
	"github.com/juandunbar/yobirthday/yodata"
	"github.com/juandunbar/yobirthday/yotemplate"
)

var (
	config *yoconfig.Config
	ErrConfigNotSet = errors.New("config not set, call email.Config() to set")
)

// Main email client interface, any email driver must implement
type EmailClient interface {
	Init() error
	SendEmails() error
}

// Custom error type used by NewClient function
type ErrClientNotFound struct{
	Client string
}
func (e *ErrClientNotFound) Error() string {
	return fmt.Sprintf("email client not implemented: %v", e.Client)
}

func NewClient() (EmailClient, error) {
	if config == nil {
		return nil, ErrConfigNotSet
	}
	// Initialize an email client using our email driver value
	// Currently only Mailgun is supported
	driver := config.GetString("email_driver")
	switch driver {
	case "mailgun":
		client := &MailGun{}
		err := client.Init()
		if err != nil {
			return nil, err
		}
		return client, nil
	default:
		return nil, &ErrClientNotFound{driver}
	}
}

// This function will grab a canned response email body using the yodata.Email.Type
func GetEmailBody(email *yodata.Email) (string, error) {
	response, err := yotemplate.NewResponse()
	if err != nil {
		return "", err
	}
	body, err := response.Render(email)
	if err != nil {
		return "", err
	}
	return body, nil
}

func SetConfig(c *yoconfig.Config) {
	config = c
}
