package yoemail

import (
	"errors"
	"fmt"
	"github.com/juandunbar/yobirthday/yoconfig"
)

var (
	config *yoconfig.Config
	ErrConfigNotSet = errors.New("config not set, call email.Config() to set")
)

type EmailClient interface {
	Init(config *yoconfig.Config) EmailClient
	SendEmail() error
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

	driver := config.GetString("email_driver")
	switch driver {
	case "mailgun":
		client := &MailGun{}
		return client.Init(config), nil
	default:
		return nil, &ErrClientNotFound{driver}
	}
}

func SetConfig(c *yoconfig.Config) {
	config = c
}
