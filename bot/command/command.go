package command

import (
	"github.com/41x3n/TeleUtil/bootstrap"
	"github.com/41x3n/TeleUtil/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	CommandHelp  = "help"
	CommandStart = "start"
)

func extractUserInfo(update *tgbotapi.Update) *domain.User {
	user := update.Message.From
	return &domain.User{
		UserID:       user.ID,
		IsBot:        user.IsBot,
		FirstName:    user.FirstName,
		LastName:     &user.LastName,
		UserName:     &user.UserName,
		LanguageCode: &user.LanguageCode,
		IsActive:     true,
	}
}

func HandleCommand(update *tgbotapi.Update, msg *tgbotapi.MessageConfig, app *bootstrap.Application) {
	command := update.Message.Command()
	user := extractUserInfo(update)
	switch command {
	case CommandHelp:
		HelpCommand(update, user, msg, app)
	case CommandStart:
		StartCommand(update, user, msg, app)
	default:
		msg.Text = "I don't know that command"
	}

}
