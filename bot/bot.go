package bot

import (
	"log"

	"github.com/41x3n/TeleUtil/bootstrap"
	"github.com/41x3n/TeleUtil/bot/command"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdates(app *bootstrap.Application) {
	bot := app.Bot
	api := bot.API

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := api.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		command.HandleCommand(&update, &msg, app)

		if _, err := api.Send(msg); err != nil {
			log.Printf("Failed to send message: %v", err)
			continue
		}
	}
}
