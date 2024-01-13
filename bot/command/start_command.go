package command

import (
	"github.com/41x3n/TeleUtil/bot/controller"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartCommand(msg *tgbotapi.MessageConfig) {
	controller.StartController(msg)
}
