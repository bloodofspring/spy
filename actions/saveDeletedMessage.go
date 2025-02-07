package actions

import (
	"fmt"
	"main/database"

	models "main/database/models"

	"github.com/go-pg/pg/v10"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SaveDeletedMessage struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e SaveDeletedMessage) fabricateAnswer(update tgbotapi.Update, mId int) (tgbotapi.Chattable, error) {
	messageDb := &models.Message{}

	db := database.Connect()
	defer db.Close()

	err := db.Model(messageDb).
		Where("tg_id = ? AND business_connection_id = ?", mId, update.DeletedBusinnesMessage.BusinessConnectionId).
		Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	if messageDb.Text == "" {
		return nil, fmt.Errorf("message text is empty")
	}

	username := update.DeletedBusinnesMessage.Chat.UserName
	if username == "" {
		username = "Пользователь"
	}

	_, err = db.Model(messageDb).WherePK().Delete()
	if err != nil {
		return nil, err
	}

	msg := tgbotapi.NewMessage(messageDb.ToUserTgId, fmt.Sprintf("@%s удалил(а) сообщение:\n<blockquote>%s</blockquote>", username, messageDb.Text))
	msg.ParseMode = "HTML"
	return msg, nil
}

func (e SaveDeletedMessage) Run(update tgbotapi.Update) error {
	for _, mId := range update.DeletedBusinnesMessage.DeletedMessagesIds {
		ans, err := e.fabricateAnswer(update, mId)
		if err != nil {
			if err == pg.ErrNoRows {
				continue
			}
			return err
		}

		_, err = e.Client.Send(ans)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e SaveDeletedMessage) GetName() string {
	return e.Name
}
