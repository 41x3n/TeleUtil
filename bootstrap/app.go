package bootstrap

import (
	"gorm.io/gorm"
)

type Application struct {
	Env      *Env
	Postgres *gorm.DB
	Bot      *Bot
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Postgres = NewPostgresDatabase(app.Env)
	app.Bot = NewBot(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	ClosePostgresDBConnection(app.Postgres)
}

func (app *Application) HandleBotUpdates() {
	HandleUpdates(app.Bot)
}
