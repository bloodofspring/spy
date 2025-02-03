package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/actions"
	"main/database"
	"main/filters"
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
	return handlers.ActiveHandlers{Handlers: []handlers.Handler{
		handlers.CommandHandler.Product(actions.Start{Name: "start-cmd", Client: bot}, []handlers.Filter{filters.StartCommandFilter}),
		handlers.BusinnesMessageHandler.Product(actions.SavePhoto{Name: "save-secret-photo", Client: bot}, []handlers.Filter{filters.ReplyPhotoFilter}),
		handlers.BusinnesMessageHandler.Product(actions.SaveVideoNoteCallback{Name: "save-secret-video-note", Client: bot}, []handlers.Filter{filters.ReplyVideoNoteFilter}),
		handlers.BusinnesMessageHandler.Product(actions.SaveVideoMessage{Name: "save-secret-video", Client: bot}, []handlers.Filter{filters.ReplyVideoFilter}),
		handlers.BusinnesMessageHandler.Product(actions.SaveVoiceMessage{Name: "save-secret-voice", Client: bot}, []handlers.Filter{filters.ReplyVoiceFilter}),
		handlers.EditedBusinnesMessageHandler.Product(actions.SaveEdiedMessage{Name: "resend-edited-message", Client: bot}, []handlers.Filter{filters.MessageEditedByInterlocutor}),
		handlers.DeletedBusinnesMessageHandler.Product(actions.SaveDeletedMessage{Name: "resend-deleted-message", Client: bot}, []handlers.Filter{filters.AllFilter}),
		handlers.BusinnesMessageHandler.Product(actions.RegisterMessage{Name: "reg-message", Client: bot}, []handlers.Filter{filters.TextMessageFilter}),

		handlers.CallbackQueryHandler.Product(actions.BugReport{Name: "bug-report", Client: bot}, []handlers.Filter{filters.BugReportCallDataFilter}),
		handlers.CallbackQueryHandler.Product(actions.Settings{Name: "settings", Client: bot}, []handlers.Filter{filters.SettingsCallDataFilter}),
		handlers.CallbackQueryHandler.Product(actions.SetupInstruction{Name: "setup-instruction", Client: bot}, []handlers.Filter{filters.InstructionCallDataFilter}),

		handlers.CallbackQueryHandler.Product(actions.Start{Name: "start-cmd", Client: bot}, []handlers.Filter{filters.ToMainCallDataFilter}),
	}}
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
