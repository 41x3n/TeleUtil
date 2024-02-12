package bootstrap

import (
	"gorm.io/gorm"
)

type Application struct {
	Env      *Env
	Postgres *gorm.DB
	Bot      *Bot
	RabbitMQ *RabbitMQ
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Postgres = NewPostgresDatabase(app.Env)
	app.Bot = NewBot(app.Env)
	app.RabbitMQ = NewRabbitMQ(app.Env)
	return *app
}

func (app *Application) AutoMigrate() {
	AutoMigrate(app.Postgres)
}

func (app *Application) CloseDBConnection() {
	ClosePostgresDBConnection(app.Postgres)
}

func (app *Application) CloseRabbitMQ() {
	CloseRabbitMQChannel(app.RabbitMQ.Ch)
	CloseRabbitMQConnection(app.RabbitMQ.Conn)
}
