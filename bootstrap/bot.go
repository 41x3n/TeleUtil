package bootstrap

import (
	"log"

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
