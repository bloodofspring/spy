package actions

import (
	"fmt"
	"main/database"

	models "main/database/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SaveEdiedMessage struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e SaveEdiedMessage) fabricateAnswer(update tgbotapi.Update) (tgbotapi.Chattable, error) {
	db := database.Connect()
	defer db.Close()

	messageDb := &models.Message{}
	err := db.Model(messageDb).
		Where("tg_id = ? AND business_connection_id = ?", update.EditedBusinnesMessage.MessageID, update.EditedBusinnesMessage.BusinessConnectionId).
		Select()

	if err != nil {
		return tgbotapi.NewMessage(-1, ""), err
	}

	oldMessageText := messageDb.Text

	messageDb.Text = update.EditedBusinnesMessage.Text
	_, err = db.Model(messageDb).WherePK().Update()
	if err != nil {
		return tgbotapi.NewMessage(-1, ""), err
	}

	sendToDb := &models.TelegramUser{}
	err = db.Model(sendToDb).
		Where("business_connection_id = ?", update.EditedBusinnesMessage.BusinessConnectionId).
		Select()
	if err != nil {
		return tgbotapi.NewMessage(-1, ""), err
	}

	msg := tgbotapi.NewMessage(sendToDb.TgId, fmt.Sprintf("@%s изменил(а) сообщение:\n<blockquote>%s</blockquote>\nна\n<blockquote>%s</blockquote>", update.EditedBusinnesMessage.From.UserName, oldMessageText, update.EditedBusinnesMessage.Text))
	msg.ParseMode = "HTML"
	return msg, nil
}

func (e SaveEdiedMessage) Run(update tgbotapi.Update) error {
	ans, err := e.fabricateAnswer(update)
	if err != nil {
		return err
	}

	_, err = e.Client.Send(ans)

	return err
}

func (e SaveEdiedMessage) GetName() string {
	return e.Name
}