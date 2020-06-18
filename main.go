package main

import (
	"github.com/juandunbar/yobirthday/yoconfig"
	"github.com/juandunbar/yobirthday/yodata"
	"github.com/juandunbar/yobirthday/yoemail"
	"log"
)

func main() {
	// Load our config and environment variables
	config, err := yoconfig.Load("")
	if err != nil {
		log.Fatal(err)
	}
	// Share our config state with our email and data package
	yoemail.SetConfig(config)
	yodata.SetConfig(config)

	// Create a new email client
	emailClient, err := yoemail.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	if err = emailClient.SendEmails(); err != nil {
		log.Fatal(err)
	}
}
