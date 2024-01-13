package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	CommandHelp  = "help"
	CommandStart = "start"
)

func HandleCommand(msg *tgbotapi.MessageConfig, commandProp string) {
	switch commandProp {
	case CommandHelp:
		HelpCommand(msg)
	case CommandStart:
		StartCommand(msg)
	default:
		msg.Text = "I don't know that command"
	}
}
