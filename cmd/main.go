package main

import (
	"github.com/41x3n/TeleUtil/bootstrap"
)

func main() {

	app := bootstrap.App()

	// env := app.Env
	// db := app.Postgres
	// bot := app.Bot

	defer app.CloseDBConnection()

	app.HandleBotUpdates()

}
