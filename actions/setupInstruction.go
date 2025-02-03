package actions

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type SetupInstruction struct {
	Name string
	Client tgbotapi.BotAPI
}

func (e SetupInstruction) fabricateAnswer(update tgbotapi.Update) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "")
	msg.Text = "Инструкция по установке бота..."


	toMainCallbackData := "toMain"

	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.InlineKeyboardButton{Text: "На главную", CallbackData: &toMainCallbackData}},
	}}


	return msg
}

func (e SetupInstruction) Run (update tgbotapi.Update) error {
	msg := e.fabricateAnswer(update)

	_, err := e.Client.Send(msg)

	return err
}

func (e SetupInstruction) GetName() string {
	return e.Name
}
