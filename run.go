package main

import (
	"fmt"
	"log"
	"main/actions"
	"main/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	godotenv "github.com/joho/godotenv"
)

func connect(debug bool) *tgbotapi.BotAPI {
	envFile, _ := godotenv.Read(".env")

	bot, err := tgbotapi.NewBotAPI(envFile["API_KEY"])
	if err != nil {
		panic(err)
	}

	bot.Debug = debug
	log.Printf("Successfully authorized on account @%s", bot.Self.UserName)

	return bot
}

func getBotActions(bot tgbotapi.BotAPI) handlers.ActiveHandlers {
	startFilter := func(update tgbotapi.Update) bool { return update.Message.Command() == "start" }
	// replyPhotoFilter := func (update tgbotapi.Update) bool {return update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.Photo != nil}
	photoFilter := func(update tgbotapi.Update) bool { return update.BusinnesMessage != nil && update.BusinnesMessage.Photo != nil }

	act := handlers.ActiveHandlers{Handlers: []handlers.Handler{
		// Place your handlers here
		handlers.CommandHandler.Product(actions.SayHi{Name: "start-cmd", Client: bot}, []handlers.Filter{startFilter}),
		handlers.BusinnesMessageHandler.Product(actions.SaveFile{Name: "save-file", Client: bot}, []handlers.Filter{photoFilter}),
	}}

	return act
}

func main() {
	client := connect(true)
	act := getBotActions(*client)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := client.GetUpdatesChan(updateConfig)

	fmt.Println(updates)

	for update := range updates {
		_ = act.HandleAll(update)
	}
}
