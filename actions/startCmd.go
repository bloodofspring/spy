package actions

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Start struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e Start) fabricateSendAnswer(update tgbotapi.Update) tgbotapi.Chattable {
	msg := tgbotapi.NewAnimation(update.Message.Chat.ID, tgbotapi.FileID("CgACAgIAAxkBAAP2Z6T7_pfO3KIMLK-gSVFkVTRtWHsAAppoAAIIOilJmvsd946QA7k2BA"))
	msg.Caption = "Добро пожаловать!\n\nЭтот бот создан, чтобы отслеживать действия ваших собеседников в переписке.\n\nЕсли ваш собеседник изменит или удалит сообщение — вы моментально об этом узнаете 🔔\n\nТакже бот умеет скачивать фото/видео/голосовые/кружки, отправленные с таймером ⏳\n\n<b><i>❓КАК ПОДКЛЮЧИТЬ БОТА</i></b>\nСмотрите на видео выше. Также ниже представлена пошаговая инструкция по установке и настройке @ChatDetectiveBot:\n\n1. Зайдите в настройки Telegram\n\n2. Пролистайте открывшееся меню вниз и перейдите в раздел 'Telegram для бизнеса'\n\n3.Выберите раздел 'чат-боты'. В строке поиска наберите имя пользователя бота (@ChatDetectiveBot) и нажмите на кнопку 'добавить'.\n\nГотово! Ниже можно выбрать, в каких чатах будет работать бот."
	msg.ParseMode = "HTML"

	exampleOfUsageCallbackData := "exampleOfUsage"
	webAppURL := "https://bloodofspring.github.io/spy/webApp/index.html"

	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.InlineKeyboardButton{Text: "Примеры использования", CallbackData: &exampleOfUsageCallbackData}},
		{tgbotapi.InlineKeyboardButton{Text: "Часто задаваемые вопросы", WebApp: &tgbotapi.WebApp{URL: &webAppURL}}},
	}}

	return msg
}

func (e Start) Run(update tgbotapi.Update) error {
	_, err := e.Client.Send(e.fabricateSendAnswer(update))

	return err
}

func (e Start) GetName() string {
	return e.Name
}
