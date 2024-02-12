package command

import (
	"log"

	"github.com/41x3n/TeleUtil/bootstrap"
	"github.com/41x3n/TeleUtil/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	HELP  = "help"
	START = "start"
	PHOTO = "photo"
)

func extractUserInfo(update *tgbotapi.Update) *domain.User {
	user := update.Message.From
	return &domain.User{
		TelegramID:   user.ID,
		IsBot:        user.IsBot,
		FirstName:    user.FirstName,
		LastName:     &user.LastName,
		UserName:     &user.UserName,
		LanguageCode: &user.LanguageCode,
		IsActive:     true,
	}
}

func HandleCommand(command string, update *tgbotapi.Update, msg *tgbotapi.MessageConfig, app *bootstrap.Application) {
	log.Println("Handling command", command)
	user := extractUserInfo(update)

	switch command {
	case HELP:
		HelpCommand(update, user, msg, app)
	case START:
		StartCommand(update, user, msg, app)
	case PHOTO:
		PhotoCommand(update, user, msg, app)
	default:
		msg.Text = "I don't know that command"
	}

}
