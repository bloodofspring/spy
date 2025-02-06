package actions

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type ExamplesOfUsage struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e ExamplesOfUsage) fabricateAnswer(update tgbotapi.Update) tgbotapi.MediaGroupConfig {
	envFile, _ := godotenv.Read(".env")
	var workExampleFileIds []string
	json.Unmarshal([]byte(envFile["work_example_file_ids"]), &workExampleFileIds)

	firstVideo := tgbotapi.NewInputMediaVideo(tgbotapi.FileID(workExampleFileIds[0]))
	secondVideo := tgbotapi.NewInputMediaVideo(tgbotapi.FileID(workExampleFileIds[1]))
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
