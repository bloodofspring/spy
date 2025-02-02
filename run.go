package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/actions"
	"main/database"
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
	allFilter := func(update tgbotapi.Update) bool { return true }
	textMessageFilter := func(update tgbotapi.Update) bool {
		return update.BusinnesMessage.Text != ""
	}
	startFilter := func(update tgbotapi.Update) bool {
		return update.Message.Command() == "start"
	}
	replyPhotoFilter := func(update tgbotapi.Update) bool {
		return update.BusinnesMessage.ReplyToMessage != nil && update.BusinnesMessage.ReplyToMessage.Photo != nil && update.BusinnesMessage.ReplyToMessage.HasProtectedContent
	}
	replyVideoNoteFilter := func(update tgbotapi.Update) bool {
		return update.BusinnesMessage.ReplyToMessage != nil && update.BusinnesMessage.ReplyToMessage.VideoNote != nil && update.BusinnesMessage.ReplyToMessage.HasProtectedContent
	}
	replyVoiceFilter := func(update tgbotapi.Update) bool {
		return update.BusinnesMessage.ReplyToMessage != nil && update.BusinnesMessage.ReplyToMessage.Voice != nil && update.BusinnesMessage.ReplyToMessage.HasProtectedContent
	}

	act := handlers.ActiveHandlers{Handlers: []handlers.Handler{
		// Place your handlers here
		handlers.CommandHandler.Product(actions.SayHi{Name: "start-cmd", Client: bot}, []handlers.Filter{startFilter}),
		handlers.BusinnesMessageHandler.Product(actions.SaveFile{Name: "save-secret-photo", Client: bot}, []handlers.Filter{replyPhotoFilter}),
		handlers.BusinnesMessageHandler.Product(actions.SaveVideoMessageCallback{Name: "save-secret-video-note", Client: bot}, []handlers.Filter{replyVideoNoteFilter}),
		handlers.BusinnesMessageHandler.Product(actions.SaveVoiceMessageCallback{Name: "save-secret-voice", Client: bot}, []handlers.Filter{replyVoiceFilter}),
		// ToDo: replace allFilter function with username filter
		handlers.EditedBusinnesMessageHandler.Product(actions.SaveEdiedMessage{Name: "resend-edited-message", Client: bot}, []handlers.Filter{allFilter}),
		handlers.DeletedBusinnesMessageHandler.Product(actions.SaveDeletedMessage{Name: "resend-deleted-message", Client: bot}, []handlers.Filter{allFilter}),

		handlers.BusinnesMessageHandler.Product(actions.RegisterMessage{Name: "reg-message", Client: bot}, []handlers.Filter{textMessageFilter}),
	}}

	return act
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "    ")
	return string(s)
}

func main() {
	err := database.InitDb()
	if err != nil {
		panic(err)
	}

	log.Println("Database init finished without errors!")

	client := connect(true)
	act := getBotActions(*client)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := client.GetUpdatesChan(updateConfig)

	for update := range updates {
		fmt.Println(prettyPrint(update))
		_ = act.HandleAll(update)
	}
}
