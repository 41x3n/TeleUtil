package controller

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func HelpController(msg *tgbotapi.MessageConfig) {
	msg.Text = "In progress. We will list down everything that you need to know about this bot."
}
