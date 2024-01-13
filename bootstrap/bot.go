package bootstrap

import (
	"log"

	"github.com/41x3n/TeleUtil/bot/command"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	API *tgbotapi.BotAPI
}

func NewBot(env *Env) *Bot {
	token := env.TelegramBotToken

	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	botAPI.Debug = true

	log.Println("Connection to Telegram established.")
	log.Println("Authorized on account ", botAPI.Self.UserName)

	bot := &Bot{
		API: botAPI,
	}

	return bot
}
func HandleUpdates(bot *Bot) {
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

		command.HandleCommand(&msg, update.Message.Command())

		if _, err := api.Send(msg); err != nil {
			log.Printf("Failed to send message: %v", err)
			continue
		}
	}
}
