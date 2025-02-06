package actions

import (
	"main/database"
	"main/database/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Settings struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e Settings) getKeyboard(settingsDb *models.UserSettings) *[][]tgbotapi.InlineKeyboardButton {
	type buttonConfig struct {
		text    string
		setting bool
		action  string
	}

	createButton := func(cfg buttonConfig) tgbotapi.InlineKeyboardButton {
		status := " ❌"
		action := cfg.action + "true"
		if cfg.setting {
			status = " ✅"
			action = cfg.action + "false"
		}
		callbackData := action
		return tgbotapi.InlineKeyboardButton{
			Text:         cfg.text + status,
			CallbackData: &callbackData,
		}
	}

	if !settingsDb.GetEvents {
		return &[][]tgbotapi.InlineKeyboardButton{
			{createButton(buttonConfig{"Уведомления", false, "settings-GetEvents-"})},
		}
	}

	buttons := []buttonConfig{
		{"Уведомления", settingsDb.GetEvents, "settings-GetEvents-"},
		{"Удаленные сообщения", settingsDb.SaveDeletedMessages, "settings-DeletedMessages-"},
		{"Измененные сообщения", settingsDb.SaveEditedMessages, "settings-EditedMessages-"},
		{"Секретные фото", settingsDb.SaveSelfDistructingPhotos, "settings-SecretPhotos-"},
		{"Секретные видеосообщения", settingsDb.SaveSelfDistructingVideoNotes, "settings-SecretVideoNotes-"},
		{"Секретные голосовые сообщения", settingsDb.SaveSelfDistructingVoices, "settings-SecretVoices-"},
		{"Секретные видео", settingsDb.SaveSelfDistructingVideos, "settings-SecretVideos-"},
	}

	keyboard := make([][]tgbotapi.InlineKeyboardButton, len(buttons))
	for i, btn := range buttons {
		keyboard[i] = []tgbotapi.InlineKeyboardButton{createButton(btn)}
	}

	return &keyboard
}

func (e Settings) fabricateAnswer(message tgbotapi.Message, keyboard *[][]tgbotapi.InlineKeyboardButton, createEditedMessage bool) tgbotapi.Chattable {
	const text = "<b><i>==НАСТРОЙКИ==</i></b>\nНа этой странице вы можете выбрать, о каких событиях вам будут приходить уведомления:"
	markup := tgbotapi.InlineKeyboardMarkup{InlineKeyboard: *keyboard}

	if createEditedMessage {
		msg := tgbotapi.NewEditMessageText(message.From.ID, message.MessageID, text)
		msg.ParseMode = "HTML"
		msg.ReplyMarkup = &markup
		return msg
	}

	msg := tgbotapi.NewMessage(message.From.ID, text)
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = &markup
	return msg
}

func (e Settings) Run(update tgbotapi.Update) error {
	var message tgbotapi.Message
	var createEditedMessage bool

	switch {
	case update.CallbackQuery != nil:
		message = *update.CallbackQuery.Message
		message.From = update.CallbackQuery.From
		createEditedMessage = true
	case update.Message != nil:
		message = *update.Message
		createEditedMessage = false
	}

	settings := models.UserSettings{}
	user := models.TelegramUser{}

	db := database.Connect()
	defer db.Close()

	err := db.Model(&user).
		Where("tg_id = ?", message.From.ID).
		Select()

	if err != nil {
		return err
	}

	err = db.Model(&settings).
		Where("user_tg_id = ?", user.TgId).
		Select()
	if err != nil {
		return err
	}

	msg := e.fabricateAnswer(message, e.getKeyboard(&settings), createEditedMessage)
	_, err = e.Client.Send(msg)

	return err
}

func (e Settings) GetName() string {
	return e.Name
}
