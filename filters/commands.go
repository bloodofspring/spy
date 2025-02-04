package filters

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var StartCommandFilter = func(update tgbotapi.Update) bool {
	return update.Message.Command() == "start"
}

var BugCommandFilter = func(update tgbotapi.Update) bool {
	return update.Message.Command() == "bug"
}
