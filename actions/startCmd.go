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

	instructionCallbackData := "instruction"
	settingsCallbackData := "settings"
	bugReportCallbackData := "bugReport"
	additionalInfoURL := "https://telegram.org"

	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.InlineKeyboardButton{Text: "Инструкция по установке", CallbackData: &instructionCallbackData}},
		{tgbotapi.InlineKeyboardButton{Text: "Настройки", CallbackData: &settingsCallbackData}, tgbotapi.InlineKeyboardButton{Text: "Информация о боте", URL: &additionalInfoURL}},
		{tgbotapi.InlineKeyboardButton{Text: "Сообщить об ошибке", CallbackData: &bugReportCallbackData}},
	}}

	return msg
}

func (e SayHi) Run(update tgbotapi.Update) error {
	if err := database.UpdateAllUserData(update.Message.Chat.ID, "", true); err != nil {
		return err
	}

	if _, err := e.Client.Send(e.fabricateAnswer(update)); err != nil {
		return err
	}

	return nil
}

func (e SayHi) GetName() string {
	return e.Name
}
