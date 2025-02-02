package actions

import (
	"fmt"
	"main/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SaveEdiedMessage struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e SaveEdiedMessage) fabricateAnswer(update tgbotapi.Update) tgbotapi.Chattable {
	// ToDo: Добавить сравнение было | стало (после реализации истории событий в бд)
	msg := tgbotapi.NewMessage(1044385209, fmt.Sprintf("Пользователь %s изменил сообщение:\n<blockquote>%s</blockquote>", update.EditedBusinnesMessage.From.UserName, update.EditedBusinnesMessage.Text))
	msg.ParseMode = "HTML"
	return msg
}

func (e SaveEdiedMessage) Run(update tgbotapi.Update) error {
	if err := database.UpdateAllUserData(update.EditedBusinnesMessage.From.ID, update.EditedBusinnesMessage.BusinessConnectionId, false); err != nil {
		return err
	}

	_, err := e.Client.Send(e.fabricateAnswer(update))

	return err
}

func (e SaveEdiedMessage) GetName() string {
	return e.Name
}
