package main

import (
	"log"
	"os"

	"github.com/41x3n/TeleUtil/bootstrap"
	"github.com/41x3n/TeleUtil/bot"
	"github.com/41x3n/TeleUtil/rabbit"
)

func main() {
	// Customize the log output format
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Llongfile)

	app := bootstrap.App()

	defer app.CloseDBConnection()
	defer app.CloseRabbitMQ()

	app.AutoMigrate()

	go rabbit.ConsumeMessages(&app)
	go bot.HandleUpdates(&app)

	// Wait indefinitely
	select {}

}
