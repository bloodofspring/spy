package filters

import (
	"main/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var AllFilter = func(update tgbotapi.Update) bool { return true }

var MessageEditedByInterlocutor = func(update tgbotapi.Update) bool {
	return update.EditedBusinnesMessage.From.ID == update.EditedBusinnesMessage.Chat.ID
}

var ReceiveEditedMessagesFilter = func(update tgbotapi.Update) bool {
	settings, err := database.GetUserSettings(update)

	return err == nil && settings.GetEvents && settings.SaveEditedMessages
}

var TextMessageFilter = func(update tgbotapi.Update) bool {
	return update.BusinnesMessage.Text != ""
}

var ReceiveDeletedMessagesFilter = func(update tgbotapi.Update) bool {
	settings, err := database.GetUserSettings(update)

	return err == nil && settings.GetEvents && settings.SaveDeletedMessages
}