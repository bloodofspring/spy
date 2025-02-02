package actions

import (
	"fmt"
	"main/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SaveDeletedMessage struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e SaveDeletedMessage) fabricateAnswer(update tgbotapi.Update) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(1044385209, fmt.Sprintf("Пользователь %s удалил сообщение:\n<blockquote>IDs=%v</blockquote>", update.DeletedBusinnesMessage.Chat.UserName, update.DeletedBusinnesMessage.DeletedMessagesIds))
	msg.ParseMode = "HTML"
	return msg
}

func (e SaveDeletedMessage) Run(update tgbotapi.Update) error {
	if err := database.UpdateAllUserData(update.DeletedBusinnesMessage.From.ID, update.DeletedBusinnesMessage.BusinessConnectionId, false); err != nil {
		return err
	}

	_, err := e.Client.Send(e.fabricateAnswer(update))

	return err
}

func (e SaveDeletedMessage) GetName() string {
	return e.Name
}
