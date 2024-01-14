package command

import (
	"github.com/41x3n/TeleUtil/bootstrap"
	"github.com/41x3n/TeleUtil/bot/controller"
	"github.com/41x3n/TeleUtil/domain"
	"github.com/41x3n/TeleUtil/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartCommand(update *tgbotapi.Update, user *domain.User, msg *tgbotapi.MessageConfig, app *bootstrap.Application) {
	ur := repository.NewUserRepository(app.Postgres, domain.TableUser)
	sc := controller.NewStartController(ur)
	if err := sc.Run(user); err != nil {
		msg.Text = "Something wrong happened"
		return
	}

	msg.Text = "Welcome to TeleUtil"
}
