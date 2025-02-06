package actions

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ExamplesOfUsage struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e ExamplesOfUsage) fabricateAnswer(update tgbotapi.Update) tgbotapi.MediaGroupConfig {
	firstVideo := tgbotapi.NewInputMediaVideo(tgbotapi.FileID("BAACAgIAAxkBAAPhZ6T0KVlxh8l38rCefIqyzGZrPYUAAjhoAAIIOilJ6QKtUFKWxjc2BA"))
	secondVideo := tgbotapi.NewInputMediaVideo(tgbotapi.FileID("BAACAgIAAxkBAAPiZ6T0NjHITKAJ0jCUj24WoiaFX7kAAjloAAIIOilJeISmhDVU1hM2BA"))
	secondVideo.Caption = "<b>Демонстрация работы бота</b>\n\nВидео 1: Скачивание фото с таймером\n\nВидео 2: Скачивание кружочка с таймером\n\n<b>Бот работает даже когда вы оффлайн!</b>"
	secondVideo.ParseMode = "HTML"

	mediaGroup := tgbotapi.NewMediaGroup(update.CallbackQuery.From.ID, []interface{}{firstVideo, secondVideo})

	return mediaGroup
}

func (e ExamplesOfUsage) Run(update tgbotapi.Update) error {
	_, err := e.Client.SendMediaGroup(e.fabricateAnswer(update))
	
	return err
}

func (e ExamplesOfUsage) GetName() string {
	return e.Name
}
