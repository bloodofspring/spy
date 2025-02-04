package filters

import (
	"main/database"
	"main/database/models"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func equalCallbackData(update tgbotapi.Update, callData string) bool {
	return update.CallbackQuery.Data == callData
}

var SettingsCallDataFilter = func(update tgbotapi.Update) bool {
	data := update.CallbackQuery.Data
	if !strings.HasPrefix(data, "settings") {
		return false
	}

	db := database.Connect()
	defer db.Close()

	var user models.TelegramUser
	if err := db.Model(&user).Where("tg_id = ?", update.CallbackQuery.From.ID).Select(); err != nil {
		return false
	}

	var settings models.UserSettings
	if err := db.Model(&settings).Where("user_id = ?", user.Id).Select(); err != nil {
		return false
	}

	settingsMap := map[string]*bool{
		"GetEvents":        &settings.GetEvents,
		"DeletedMessages":  &settings.SaveDeletedMessages,
		"EditedMessages":   &settings.SaveEditedMessages,
		"SecretPhotos":     &settings.SaveSelfDistructingPhotos,
		"SecretVideoNotes": &settings.SaveSelfDistructingVideoNotes,
		"SecretVoices":     &settings.SaveSelfDistructingVoices,
		"SecretVideos":     &settings.SaveSelfDistructingVidoes, // ToDo: Fix typo in struct
	}

	for settingName, settingPtr := range settingsMap {
		if strings.Contains(data, settingName) {
			*settingPtr = strings.Split(data, "-")[2] == "true"
			break
		}
	}

	_, err := db.Model(&settings).Where("user_id = ?", user.Id).Update()
	return err == nil
}

var InstructionCallDataFilter = func(update tgbotapi.Update) bool {
	return equalCallbackData(update, "instruction")
}

var ToMainCallDataFilter = func(update tgbotapi.Update) bool {
	return equalCallbackData(update, "toMain")
}
