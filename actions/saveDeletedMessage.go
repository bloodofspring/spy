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

func (e SaveDeletedMessage) fabricateAnswer(update tgbotapi.Update, mId int, sendTo models.TelegramUser, db *pg.DB) (tgbotapi.Chattable, error) {
	messageDb := &models.Message{}
	err := db.Model(messageDb).
		Where("tg_id = ? AND business_connection_id = ?", mId, update.DeletedBusinnesMessage.BusinessConnectionId).
		Select()
	if err != nil {
		return tgbotapi.NewMessage(-1, ""), err
	}

	_, err = db.Model(messageDb).WherePK().Delete()
	if err != nil {
		return tgbotapi.NewMessage(-1, ""), err
	}

	msg := tgbotapi.NewMessage(sendTo.TgId, fmt.Sprintf("Пользователь %s удалил сообщение:\n<blockquote>%s</blockquote>", update.DeletedBusinnesMessage.Chat.UserName, messageDb.Text))
	msg.ParseMode = "HTML"
	return msg, nil
}

func (e SaveDeletedMessage) Run(update tgbotapi.Update) error {
	// if err := database.UpdateAllUserData(update.DeletedBusinnesMessage.From.ID, update.DeletedBusinnesMessage.BusinessConnectionId, false); err != nil {
	// 	return err
	// }  ToDo: Запихнуть куда-нибудь этот апдейт
	db := database.Connect()
	defer db.Close()

	sendTo := &models.TelegramUser{}
	err := db.Model(sendTo).
		Where("business_connection_id = ?", update.DeletedBusinnesMessage.BusinessConnectionId).
		Select()
	if err != nil {
		return err
	}
	if sendTo.TgId == update.DeletedBusinnesMessage.Chat.ID {
		return nil  // Сообщение удалено НЕ собеседником
	}

	fmt.Println(update.DeletedBusinnesMessage.DeletedMessagesIds)
	for _, mId := range update.DeletedBusinnesMessage.DeletedMessagesIds {
		ans, err := e.fabricateAnswer(update, mId, *sendTo, db)
		if err != nil {
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
