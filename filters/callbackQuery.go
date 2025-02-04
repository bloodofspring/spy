package filters

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func equalCallbackData(update tgbotapi.Update, callData string) bool {
	return update.CallbackQuery.Data == callData
}

var SettingsCallDataFilter = func(update tgbotapi.Update) bool {
	return equalCallbackData(update, "settings")
}

var InstructionCallDataFilter = func(update tgbotapi.Update) bool {
	return equalCallbackData(update, "instruction")
}

var ToMainCallDataFilter = func(update tgbotapi.Update) bool {
	return equalCallbackData(update, "toMain")
}
