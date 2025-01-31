package actions

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type SayHi struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e SayHi) FabricateAnswer(update tgbotapi.Update) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "Hi :)"

	return msg
}

func (e SayHi) Run(update tgbotapi.Update) error {
	if _, err := e.Client.Send(e.FabricateAnswer(update)); err != nil {
		return err
	}

	return nil
}

func (e SayHi) GetName() string {
	return e.Name
}
