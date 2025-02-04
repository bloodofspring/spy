package actions

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type SetupInstruction struct {
	Name string
	Client tgbotapi.BotAPI
}

func (e SetupInstruction) fabricateAnswer(update tgbotapi.Update) tgbotapi.Chattable {
	animation := tgbotapi.NewAnimation(update.CallbackQuery.From.ID, tgbotapi.FilePath("static/addBotExample.gif"))
	animation.Caption = "<b><i>==КАК УСТАНОВИТЬ==</i></b>\nНиже представлена пошаговая инструкция по установке и настройке @ChatDetectiveBot:\n\n1. Зайдите в настройки Telegram\n\n2. Пролистайте открывшееся меню вниз и перейдите в раздел 'Telegram для бизнеса'\n\n3.Выберите раздел 'чат-боты'. В строке поиска наберите имя пользователя бота (@ChatDetectiveBot) и нажмите на кнопку 'добавить'.\n\nГотово! Ниже можно выбрать, в каких чатах будет работать бот."
	animation.ParseMode = "HTML"
	toMainCallbackData := "toMain"
	animation.ReplyMarkup = &tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.InlineKeyboardButton{Text: "На главную", CallbackData: &toMainCallbackData}},
	}}

	return animation
}

func (e SetupInstruction) Run (update tgbotapi.Update) error {
	msg := e.fabricateAnswer(update)

	_, err := e.Client.Send(msg)

	return err
}

func (e SetupInstruction) GetName() string {
	return e.Name
}
