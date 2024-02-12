package bot

import (
	"log"

	"github.com/41x3n/TeleUtil/bootstrap"
	commandHandler "github.com/41x3n/TeleUtil/bot/command"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdates(app *bootstrap.Application) {
	bot := app.Bot
	api := bot.API

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := api.GetUpdatesChan(u)
	for update := range updates {
		command := ""
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.IsCommand() { // ignore any non-command Messages
			command = update.Message.Command()
		}

		if update.Message.Photo != nil {
			command = "photo"
		}

		if command == "" {
			continue
		}

		commandHandler.HandleCommand(command, &update, &msg, app)

		if _, err := api.Send(msg); err != nil {
			log.Printf("Failed to send message: %v", err)
			continue
		}
	}
}
