package filters

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var ReplyPhotoFilter = func(update tgbotapi.Update) bool {
	return update.BusinnesMessage.ReplyToMessage != nil && update.BusinnesMessage.ReplyToMessage.Photo != nil && update.BusinnesMessage.ReplyToMessage.HasProtectedContent && update.BusinnesMessage.ReplyToMessage.From.ID == update.BusinnesMessage.ReplyToMessage.Chat.ID
}

var ReplyVideoNoteFilter = func(update tgbotapi.Update) bool {
	return update.BusinnesMessage.ReplyToMessage != nil && update.BusinnesMessage.ReplyToMessage.VideoNote != nil && update.BusinnesMessage.ReplyToMessage.HasProtectedContent && update.BusinnesMessage.ReplyToMessage.From.ID == update.BusinnesMessage.ReplyToMessage.Chat.ID
}

var ReplyVideoFilter = func(update tgbotapi.Update) bool {
	return update.BusinnesMessage.ReplyToMessage != nil && update.BusinnesMessage.ReplyToMessage.Video != nil && update.BusinnesMessage.ReplyToMessage.HasProtectedContent && update.BusinnesMessage.ReplyToMessage.From.ID == update.BusinnesMessage.ReplyToMessage.Chat.ID
}

var ReplyVoiceFilter = func(update tgbotapi.Update) bool {
	return update.BusinnesMessage.ReplyToMessage != nil && update.BusinnesMessage.ReplyToMessage.Voice != nil && update.BusinnesMessage.ReplyToMessage.HasProtectedContent && update.BusinnesMessage.ReplyToMessage.From.ID == update.BusinnesMessage.ReplyToMessage.Chat.ID
}
