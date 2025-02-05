package filters

import (
	"log"
	"main/database"
	"main/database/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var AllFilter = func(update tgbotapi.Update) bool { return true }

var MessageEditedByInterlocutor = func(update tgbotapi.Update) bool {
	return update.EditedBusinnesMessage.From.ID == update.EditedBusinnesMessage.Chat.ID
}

var ReceiveEditedMessagesFilter = func(update tgbotapi.Update) bool {
	var user models.TelegramUser

	db := database.Connect()
	defer db.Close()

	err := db.Model(&user).
		Where("business_connection_id = ?", update.EditedBusinnesMessage.BusinessConnectionId).
		Select()
	
	if err != nil {
		log.Println(err)
		return false
	}

	var settings models.UserSettings
	if err := db.Model(&settings).Where("user_tg_id = ?", user.TgId).Select(); err != nil {
		log.Println(err)
		return false
	}

	return err == nil && settings.GetEvents && settings.SaveEditedMessages
}

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