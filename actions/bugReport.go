package actions

import (
	"main/database"
	"main/database/models"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type AddBugReport struct {
	Name string
	Client tgbotapi.BotAPI
}

func (e AddBugReport) Run(update tgbotapi.Update) error {
	if strings.Trim(update.Message.Text, " ") == "/bug" {
		resp := tgbotapi.NewMessage(update.Message.Chat.ID, "Использование команды: /bug [сообщение об ошибке]")
		_, err := e.Client.Send(resp)

		return err
	}

	db := database.Connect()
	defer db.Close()

	report := &models.BugReport{
		FromUser: update.Message.From.UserName,
		Text: strings.Trim(update.Message.Text, "/bug "),
	}
	_, err := db.Model(report).Insert()
	if err != nil {
		return err
	}

	resp := tgbotapi.NewMessage(update.Message.Chat.ID, "Сообщение об ошибке сохранено, спасибо за содействие!")
	_, err = e.Client.Send(resp)

	return err
}

func (e AddBugReport) GetName() string {
	return e.Name
}
