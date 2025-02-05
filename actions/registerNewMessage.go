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
	if update.BusinnesMessage.From.ID != update.BusinnesMessage.Chat.ID {
		return nil  // message sent by owner
	}

	db := database.Connect()
	defer db.Close()

	var user models.TelegramUser
	err := db.Model(&user).
		Where("business_connection_id = ?", update.BusinnesMessage.BusinessConnectionId).
		Select()

	if err != nil {
		return err
	}

	messageDb := &models.Message{
		TgId:                 update.BusinnesMessage.MessageID,
		BusinessConnectionId: update.BusinnesMessage.BusinessConnectionId,
		FromUserTgId: update.BusinnesMessage.From.ID,
		ToUserTgId: 		  user.TgId,
		Text:                 update.BusinnesMessage.Text,
	}
	_, err = db.Model(messageDb).Insert()

	return err
}

func (e RegisterMessage) GetName() string {
	return e.Name
}
