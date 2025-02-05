package actions

import (
	"main/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UpdateUserData struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e UpdateUserData) Run(update tgbotapi.Update) error {
	var msg tgbotapi.Message
	switch {
	case update.Message != nil:
		return database.UpdateAllUserData(update.Message.From.ID, "", true)
	case update.CallbackQuery != nil:
		return database.UpdateAllUserData(update.CallbackQuery.From.ID, "", true)
	case update.BusinnesMessage != nil:
		msg = *update.BusinnesMessage
	case update.EditedBusinnesMessage != nil:
		msg = *update.EditedBusinnesMessage
	case update.DeletedBusinnesMessage != nil:
		return nil // ToDo : Implement
	default:
		return nil
	}

	if msg.Chat.ID == msg.From.ID {
		return nil // Сообщение от собеседника, Не обновляем собеседника!!!!!
	}

	return database.UpdateAllUserData(msg.From.ID, msg.BusinessConnectionId, true)
}

func (e UpdateUserData) GetName() string {
	return e.Name
}
