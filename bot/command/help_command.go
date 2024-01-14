package command

import (
	"github.com/41x3n/TeleUtil/bootstrap"
	"github.com/41x3n/TeleUtil/bot/controller"
	"github.com/41x3n/TeleUtil/domain"
	"github.com/41x3n/TeleUtil/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HelpCommand(update *tgbotapi.Update, user *domain.User, msg *tgbotapi.MessageConfig, app *bootstrap.Application) {
	ur := repository.NewUserRepository(app.Postgres, domain.TableUser)
	hc := controller.NewHelpController(ur)
	if err := hc.Run(user); err != nil {
		msg.Text = "Something wrong happened"
		return
	}

	msg.Text = "Hey my name is TeleUtil, I can help you to do pesky task easily"
}
