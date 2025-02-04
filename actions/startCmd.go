package actions

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Start struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e Start) fabricateEditedAnswer(message *tgbotapi.Message) tgbotapi.Chattable {
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")
	msg.Text = "Привет! Я помогу тебе сохранить самоуничтожаюиеся фото, голосовые и видео сообщения, оповещу тебя об удаленных или измененных сообщениях. Согласен?"

	instructionCallbackData := "instruction"
	settingsCallbackData := "settings"
	webAppURL := "https://bloodofspring.github.io/spy/webApp/index.html"

	msg.ReplyMarkup = &tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.InlineKeyboardButton{Text: "Инструкция по установке", CallbackData: &instructionCallbackData}},
		{tgbotapi.InlineKeyboardButton{Text: "Настройки", CallbackData: &settingsCallbackData}, tgbotapi.InlineKeyboardButton{Text: "Информация о боте", WebApp: &tgbotapi.WebApp{URL: &webAppURL}}},
	}}

	return msg
}

func (e Start) fabricateSendAnswer(message *tgbotapi.Message) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	msg.Text = "Привет! Я помогу тебе сохранить самоуничтожаюиеся фото, голосовые и видео сообщения, оповещу тебя об удаленных или измененных сообщениях. Согласен?"

	instructionCallbackData := "instruction"
	settingsCallbackData := "settings"
	webAppURL := "https://bloodofspring.github.io/spy/webApp/index.html"

	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.InlineKeyboardButton{Text: "Инструкция по установке", CallbackData: &instructionCallbackData}},
		{tgbotapi.InlineKeyboardButton{Text: "Настройки", CallbackData: &settingsCallbackData}, tgbotapi.InlineKeyboardButton{Text: "Информация о боте", WebApp: &tgbotapi.WebApp{URL: &webAppURL}}},
	}}

	return msg
}

func (e Start) Run(update tgbotapi.Update) error {
	var msg *tgbotapi.Message
	var edit bool

	if update.Message != nil {
		msg = update.Message
		edit = false
	} else if update.CallbackQuery != nil {
		msg = update.CallbackQuery.Message
		msg.From = update.CallbackQuery.From
		if update.CallbackQuery.Message.Text == "" {
			edit = false
		} else {
			edit = true
		}
	} else {
		return nil
	}

	var err error

	if edit {
		_, err = e.Client.Send(e.fabricateEditedAnswer(msg))
	} else {
		_, err = e.Client.Send(e.fabricateSendAnswer(msg))
	}

	return err
}

func (e Start) GetName() string {
	return e.Name
}
