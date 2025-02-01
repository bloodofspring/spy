package actions

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


type SaveEdiedMessage struct {
	Name string
	Client tgbotapi.BotAPI
}


func (e SaveEdiedMessage) fabricateAnswer(update tgbotapi.Update) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(update.EditedBusinnesMessage.From.ID, fmt.Sprintf("Пользователь %s изменил сообщение:\n<q>%s</q>", update.EditedMessage.From.UserName, update.EditedBusinnesMessage.Text))
	return msg
}


func (e SaveEdiedMessage) Run(update tgbotapi.Update) error {
	_, err := e.Client.Send(e.fabricateAnswer(update))
	
	return err
}


func (e SaveEdiedMessage) GetName() string {
	return e.Name
}

