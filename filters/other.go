package filters

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var AllFilter = func(update tgbotapi.Update) bool { return true }

var MessageEditedByInterlocutor = func(update tgbotapi.Update) bool {
	return false
	// ToDo: uncomment
	// return update.EditedBusinnesMessage.From.ID == update.EditedBusinnesMessage.Chat.ID
}

// var ReceiveEditedMessagesFilter = func(update tgbotapi.Update) bool {
// 	a, err := json.MarshalIndent(update, " ", " ")
// 	if err != nil { panic(err) }
// 	fmt.Println(string(a))
// 	settings, err := database.GetUserSettings(*update.EditedBusinnesMessage)
// 	fmt.Println("db: ", settings)

// 	return err == nil && settings.GetEvents && settings.SaveEditedMessages
// }

var TextMessageFilter = func(update tgbotapi.Update) bool {
	return update.BusinnesMessage.Text != ""
}

// var ReceiveDeletedMessagesFilter = func(update tgbotapi.Update) bool {
// 	settings, err := database.GetUserSettings(*update.DeletedBusinnesMessage)

// 	return err == nil && settings.GetEvents && settings.SaveDeletedMessages
// }

var CanUpdate = func(update tgbotapi.Update) bool {
	// ToDo: add filter body
	return true
}