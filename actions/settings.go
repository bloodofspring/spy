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

	// Если уведомления выключены, показываем только основную кнопку
	if !settingsDb.GetEvents {
		toMain := "toMain"
		return &[][]tgbotapi.InlineKeyboardButton{
			{createButton(buttonConfig{"Уведомления", false, "settings-GetEvents-"})},
			{tgbotapi.InlineKeyboardButton{Text: "На главную", CallbackData: &toMain}},
		}
	}

	buttons := []buttonConfig{
		{"Уведомления", settingsDb.GetEvents, "settings-GetEvents-"},
		// ToDo: uncomment after fix
		// {"Удаленные сообщения", settingsDb.SaveDeletedMessages, "settings-DeletedMessages-"},
		// {"Измененные сообщения", settingsDb.SaveEditedMessages, "settings-EditedMessages-"},
		{"Секретные фото", settingsDb.SaveSelfDistructingPhotos, "settings-SecretPhotos-"},
		{"Секретные водеосообщения", settingsDb.SaveSelfDistructingVideoNotes, "settings-SecretVideoNotes-"},
		{"Секретные голосовые сообщения", settingsDb.SaveSelfDistructingVoices, "settings-SecretVoices-"},
		{"Секретные видео", settingsDb.SaveSelfDistructingVideos, "settings-SecretVideos-"},
	}

	keyboard := make([][]tgbotapi.InlineKeyboardButton, len(buttons)+1)
	for i, btn := range buttons {
		keyboard[i] = []tgbotapi.InlineKeyboardButton{createButton(btn)}
	}

	toMain := "toMain"
	keyboard[len(buttons)] = []tgbotapi.InlineKeyboardButton{
		{Text: "На главную", CallbackData: &toMain},
	}

	return &keyboard
}

func (e Settings) fabricateAnswer(update tgbotapi.Update, keyboard *[][]tgbotapi.InlineKeyboardButton) tgbotapi.Chattable {
	msg := tgbotapi.NewEditMessageText(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID, "")
	msg.Text = "<b><i>==НАСТРОЙКИ==</i></b>\nНа этой странице вы можете выбрать, о каких событиях вам будут приходить уведомления:"
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = &tgbotapi.InlineKeyboardMarkup{InlineKeyboard: *keyboard}

	return msg
}

func (e Settings) Run(update tgbotapi.Update) error {
	if err := database.UpdateAllUserData(update.CallbackQuery.From.ID, "", true); err != nil {
		return err
	}

	settings := models.UserSettings{}
	user := models.TelegramUser{}

	db := database.Connect()
	defer db.Close()

	err := db.Model(&user).
		Where("tg_id = ?", update.CallbackQuery.From.ID).
		Select()

	if err != nil {
		return err
	}

	err = db.Model(&settings).
		Where("user_id = ?", user.Id).
		Select()
	if err != nil {
		return err
	}

	msg := e.fabricateAnswer(update, e.getKeyboard(&settings))
	_, err = e.Client.Send(msg)

	return err
}

func (e Settings) GetName() string {
	return e.Name
}
