package actions

import (
	"main/database"
	models "main/database/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type RegisterMessage struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e RegisterMessage) Run(update tgbotapi.Update) error {
	db := database.Connect()
	defer db.Close()

	messageDb := &models.Message{
		TgId:       update.BusinnesMessage.MessageID,
		BusinessConnectionId: update.BusinnesMessage.BusinessConnectionId,
		Text:       update.BusinnesMessage.Text,
	}
	_, err := db.Model(messageDb).Insert()

	return err
}

func (e RegisterMessage) GetName() string {
	return e.Name
}
