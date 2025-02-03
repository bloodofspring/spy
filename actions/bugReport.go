package actions

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type BugReport struct {
	Name string
	Client tgbotapi.BotAPI
}

func (e BugReport) fabricateAnswer(update tgbotapi.Update) tgbotapi.Chattable {
	msg := tgbotapi.NewEditMessageText(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, "")
	msg.Text = "Отчет об ошибке..."


	toMainCallbackData := "toMain"

	msg.ReplyMarkup = &tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.InlineKeyboardButton{Text: "На главную", CallbackData: &toMainCallbackData}},
	}}


	return msg
}

func (e BugReport) Run (update tgbotapi.Update) error {
	msg := e.fabricateAnswer(update)

	_, err := e.Client.Send(msg)

	return err
}

func (e BugReport) GetName() string {
	return e.Name
}
