package main

import (
	"encoding/json"
	"log"
	"main/actions"
	"main/database"
	"main/database/models"
	"main/filters"
	"main/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	godotenv "github.com/joho/godotenv"
)

const debug = false

func connect() *tgbotapi.BotAPI {
	envFile, _ := godotenv.Read(".env")

	bot, err := tgbotapi.NewBotAPI(envFile["API_KEY"])
	if err != nil {
		panic(err)
	}

	bot.Debug = debug
	log.Printf("Successfully authorized on account @%s", bot.Self.UserName)

	return bot
}

func IsNotAdminUser(update tgbotapi.Update) bool {
	db := database.Connect()
	defer db.Close()

	var msg tgbotapi.Message
	switch {
	case update.Message != nil:
		msg = *update.Message
	case update.BusinnesMessage != nil:
		msg = *update.BusinnesMessage
	case update.EditedBusinnesMessage != nil:
		msg = *update.EditedBusinnesMessage
	case update.DeletedBusinnesMessage != nil:
		msg = *update.DeletedBusinnesMessage
	default:
		return true
	}

	var admins []models.Admin
	err := db.Model(&admins).
		Select()
	if err != nil {
		return true
	}

	for _, a := range admins {
		if a.UserTgId == msg.Chat.ID {
			return false
		}
	}

	return true
}

func getBotActions(bot tgbotapi.BotAPI) handlers.ActiveHandlers {
	return handlers.ActiveHandlers{Handlers: []handlers.Handler{
		handlers.AllHandler.Product(actions.UpdateUserData{Name: "update-user-data", Client: bot}, []handlers.Filter{filters.CanUpdate}),
		handlers.BusinnesMessageHandler.Product(actions.RegisterMessage{Name: "reg-message", Client: bot}, []handlers.Filter{filters.TextMessageFilter}),
 
		handlers.CommandHandler.Product(actions.Start{Name: "start-cmd", Client: bot}, []handlers.Filter{filters.StartCommandFilter}),
		handlers.BusinnesMessageHandler.Product(actions.SavePhoto{Name: "save-secret-photo", Client: bot}, []handlers.Filter{filters.ReplyPhotoFilter, filters.ReceivePhotosFilter, IsNotAdminUser}),
		handlers.BusinnesMessageHandler.Product(actions.SaveVideoNoteCallback{Name: "save-secret-video-note", Client: bot}, []handlers.Filter{filters.ReplyVideoNoteFilter, filters.ReceiveVideoNotesFilter, IsNotAdminUser}),
		handlers.BusinnesMessageHandler.Product(actions.SaveVideoMessage{Name: "save-secret-video", Client: bot}, []handlers.Filter{filters.ReplyVideoFilter, filters.ReceiveVideosFilter, IsNotAdminUser}),
		handlers.BusinnesMessageHandler.Product(actions.SaveVoiceMessage{Name: "save-secret-voice", Client: bot}, []handlers.Filter{filters.ReplyVoiceFilter, filters.ReceiveVoicesFilter, IsNotAdminUser}),
		handlers.EditedBusinnesMessageHandler.Product(actions.SaveEdiedMessage{Name: "resend-edited-message", Client: bot}, []handlers.Filter{filters.MessageEditedByInterlocutor, filters.ReceiveEditedMessagesFilter, IsNotAdminUser}),
		handlers.DeletedBusinnesMessageHandler.Product(actions.SaveDeletedMessage{Name: "resend-deleted-message", Client: bot}, []handlers.Filter{filters.AllFilter, filters.ReceiveDeletedMessagesFilter, IsNotAdminUser}),

		handlers.CallbackQueryHandler.Product(actions.Settings{Name: "settings-call-data", Client: bot}, []handlers.Filter{filters.SettingsCallDataFilter}),
		handlers.CommandHandler.Product(actions.Settings{Name: "settings-command", Client: bot}, []handlers.Filter{filters.SettingsCommandFilter}),

		handlers.CallbackQueryHandler.Product(actions.Start{Name: "start-cmd", Client: bot}, []handlers.Filter{filters.ToMainCallDataFilter}),
		handlers.CommandHandler.Product(actions.AddBugReport{Name: "bug-report", Client: bot}, []handlers.Filter{filters.BugCommandFilter}),
		handlers.CallbackQueryHandler.Product(actions.ExamplesOfUsage{Name: "usage-example", Client: bot}, []handlers.Filter{filters.ExampleOfUsageCallDataFilter}),
	}}
}

func printUpdate(update *tgbotapi.Update) {
	updateJSON, err := json.MarshalIndent(update, "", "    ")
	if err != nil {
		return
	}

	log.Println(string(updateJSON))
}

func main() {
	err := database.InitDb()
	if err != nil {
		panic(err)
	}

	log.Println("Database init finished without errors!")

	client := connect()
	act := getBotActions(*client)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := client.GetUpdatesChan(updateConfig)

	for update := range updates {
		if debug { printUpdate(&update) }
		_ = act.HandleAll(update)
	}
}
