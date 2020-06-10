package main

import (
	"github.com/juandunbar/yobirthday/yoconfig"
	"github.com/juandunbar/yobirthday/yoemail"
	"log"
)

func main() {
	// Load our config and environment variables
	config, err := yoconfig.Load("")
	if err != nil {
		log.Fatal(err)
	}
	// Share our config state with other packages
	yoemail.SetConfig(config)

	// Create a new email client
	email, err := yoemail.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	if err = email.SendEmail(); err != nil {
		log.Fatal(err)
	}
}
