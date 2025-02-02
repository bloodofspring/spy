package actions

import (
	"main/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SayHi struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e SayHi) fabricateAnswer(update tgbotapi.Update) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "Привет! Я помогу тебе сохранить самоуничтожаюиеся фото, голосовые и видео сообщения, оповещу тебя об удаленных или измененных сообщениях. Согласен?"

	return msg
}

func (e SayHi) Run(update tgbotapi.Update) error {
	_, err := database.GetOrCreateUser(update.Message.Chat.ID, "")
	if err != nil {
		return err
	}

	if _, err = e.Client.Send(e.fabricateAnswer(update)); err != nil {
		return err
	}

	return nil
}

func (e SayHi) GetName() string {
	return e.Name
}
