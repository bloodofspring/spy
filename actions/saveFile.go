package actions

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


type SaveFile struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e SaveFile) FabricateAnswer(update tgbotapi.Update) tgbotapi.Chattable {
	msg := tgbotapi.NewPhoto(update.BusinnesMessage.From.ID, tgbotapi.FileID(update.BusinnesMessage.Photo[0].FileID))
	msg.Caption = "Вот ваше фото"

	return msg
}

func (e SaveFile) Run(update tgbotapi.Update) error {
	if _, err := e.Client.Send(e.FabricateAnswer(update)); err != nil {
		return err
	}

	return nil
}

func (e SaveFile) GetName() string {
	return e.Name
}


