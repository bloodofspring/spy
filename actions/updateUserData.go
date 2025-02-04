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
	switch {
	case update.Message != nil:
		return database.UpdateAllUserData(update.Message.From.ID, "", true)
	case update.BusinnesMessage != nil:
		return database.UpdateAllUserData(update.BusinnesMessage.From.ID, update.BusinnesMessage.BusinessConnectionId, false)
	case update.EditedBusinnesMessage != nil:
		return database.UpdateAllUserData(update.EditedBusinnesMessage.From.ID, update.EditedBusinnesMessage.BusinessConnectionId, false)
	case update.DeletedBusinnesMessage != nil:
		return nil // ToDo : Implement
	case update.CallbackQuery != nil:
		return database.UpdateAllUserData(update.CallbackQuery.From.ID, "", true)
	}

	return nil
}

func (e UpdateUserData) GetName() string {
	return e.Name
}
