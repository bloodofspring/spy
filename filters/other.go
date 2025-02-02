package filters

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var AllFilter = func(update tgbotapi.Update) bool { return true }

var TextMessageFilter = func(update tgbotapi.Update) bool {
	return update.BusinnesMessage.Text != ""
}
