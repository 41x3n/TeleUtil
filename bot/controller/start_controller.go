package controller

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func StartController(msg *tgbotapi.MessageConfig) {
	msg.Text = "Welcome to TeleUtil. This bot is still in progress."
}
