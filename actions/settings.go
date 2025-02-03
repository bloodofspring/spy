package actions

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Settings struct {
	Name string
	Client tgbotapi.BotAPI
}

func (e Settings) fabricateAnswer(update tgbotapi.Update) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "")
	msg.Text = "Настройки бота..."


	toMainCallbackData := "toMain"

	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.InlineKeyboardButton{Text: "На главную", CallbackData: &toMainCallbackData}},
	}}


	return msg
}

func (e Settings) Run (update tgbotapi.Update) error {
	msg := e.fabricateAnswer(update)

	_, err := e.Client.Send(msg)

	return err
}

func (e Settings) GetName() string {
	return e.Name
}
