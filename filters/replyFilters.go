package filters

import (
	"main/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var ReplyPhotoFilter = func(update tgbotapi.Update) bool {
	return update.BusinnesMessage.ReplyToMessage != nil && update.BusinnesMessage.ReplyToMessage.Photo != nil && update.BusinnesMessage.ReplyToMessage.HasProtectedContent && update.BusinnesMessage.ReplyToMessage.From.ID == update.BusinnesMessage.ReplyToMessage.Chat.ID
}

var ReceivePhotosFilter = func(update tgbotapi.Update) bool {
	settings, err := database.GetUserSettings(*update.BusinnesMessage)

	return err == nil && settings.GetEvents && settings.SaveSelfDistructingPhotos
}

var ReplyVideoNoteFilter = func(update tgbotapi.Update) bool {
	return update.BusinnesMessage.ReplyToMessage != nil && update.BusinnesMessage.ReplyToMessage.VideoNote != nil && update.BusinnesMessage.ReplyToMessage.HasProtectedContent && update.BusinnesMessage.ReplyToMessage.From.ID == update.BusinnesMessage.ReplyToMessage.Chat.ID
}

var ReceiveVideoNotesFilter = func(update tgbotapi.Update) bool {
	settings, err := database.GetUserSettings(*update.BusinnesMessage)

	return err == nil && settings.GetEvents && settings.SaveSelfDistructingVideoNotes
}

var ReplyVideoFilter = func(update tgbotapi.Update) bool {
	return update.BusinnesMessage.ReplyToMessage != nil && update.BusinnesMessage.ReplyToMessage.Video != nil && update.BusinnesMessage.ReplyToMessage.HasProtectedContent && update.BusinnesMessage.ReplyToMessage.From.ID == update.BusinnesMessage.ReplyToMessage.Chat.ID
}

var ReceiveVideosFilter = func(update tgbotapi.Update) bool {
	settings, err := database.GetUserSettings(*update.BusinnesMessage)

	return err == nil && settings.GetEvents && settings.SaveSelfDistructingVideos
}

var ReplyVoiceFilter = func(update tgbotapi.Update) bool {
	return update.BusinnesMessage.ReplyToMessage != nil && update.BusinnesMessage.ReplyToMessage.Voice != nil && update.BusinnesMessage.ReplyToMessage.HasProtectedContent && update.BusinnesMessage.ReplyToMessage.From.ID == update.BusinnesMessage.ReplyToMessage.Chat.ID
}

var ReceiveVoicesFilter = func(update tgbotapi.Update) bool {
	settings, err := database.GetUserSettings(*update.BusinnesMessage)

	return err == nil && settings.GetEvents && settings.SaveSelfDistructingVoices
}
