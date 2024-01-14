package main

import (
	"github.com/41x3n/TeleUtil/bootstrap"
	"github.com/41x3n/TeleUtil/bot"
)

func main() {

	app := bootstrap.App()

	defer app.CloseDBConnection()

	app.AutoMigrate()

	bot.HandleUpdates(&app)

}
