package actions

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type SetupInstruction struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e SetupInstruction) fabricateAnswer(update tgbotapi.Update) tgbotapi.AnimationConfig {
	msg := tgbotapi.NewAnimation(update.CallbackQuery.From.ID, tgbotapi.FileID("CgACAgIAAxkBAAP2Z6T7_pfO3KIMLK-gSVFkVTRtWHsAAppoAAIIOilJmvsd946QA7k2BA"))
	msg.Caption = "<b><i>==КАК УСТАНОВИТЬ==</i></b>\nНиже представлена пошаговая инструкция по установке и настройке @ChatDetectiveBot:\n\n1. Зайдите в настройки Telegram\n\n2. Пролистайте открывшееся меню вниз и перейдите в раздел 'Telegram для бизнеса'\n\n3.Выберите раздел 'чат-боты'. В строке поиска наберите имя пользователя бота (@ChatDetectiveBot) и нажмите на кнопку 'добавить'.\n\nГотово! Ниже можно выбрать, в каких чатах будет работать бот."
	msg.ParseMode = "HTML"

	return msg
}

func (e SetupInstruction) Run(update tgbotapi.Update) error {
	msg := e.fabricateAnswer(update)

	_, err := e.Client.Send(msg)

	return err
}

func (e SetupInstruction) GetName() string {
	return e.Name
}
