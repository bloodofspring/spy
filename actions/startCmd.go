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
	msg := tgbotapi.NewAnimation(message.Chat.ID, tgbotapi.FileID("CgACAgIAAxkBAAP2Z6T7_pfO3KIMLK-gSVFkVTRtWHsAAppoAAIIOilJmvsd946QA7k2BA"))
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
