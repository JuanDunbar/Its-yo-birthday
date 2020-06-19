package yoemail_test

import (
	"github.com/juandunbar/yobirthday/yoconfig"
	"github.com/juandunbar/yobirthday/yodata"
	"github.com/juandunbar/yobirthday/yoemail"
	"reflect"
	"testing"
)

// Setup our config and share with the data and email packages
// Use an in memory dsn so our db.Ping() doesnt error
func setupConfig() error {
	config, err := yoconfig.Load("../")
	if err != nil {
		return err
	}
	config.Set("database.dsn", "file::memory:?cache=shared")
	yoemail.SetConfig(config)
	yodata.SetConfig(config)
	return nil
}

func TestNewClient(t *testing.T) {
	// Setup our config
	if err := setupConfig(); err != nil {
		t.Fatalf("failed to setup config with error: %v", err.Error())
	}
	// Create our email client
	client, err := yoemail.NewClient()
	if err != nil {
		t.Fatalf("failed to init client with error: %v", err.Error())
	}
	// Make sure our client type is MailGun
	switch client.(type) {
	case *yoemail.MailGun:
		// Success do nothing
	default:
		t.Fatalf("invalid client type returned: %v", reflect.TypeOf(client).String())
	}
}