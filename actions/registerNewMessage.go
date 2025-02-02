package actions

// import (
// 	"encoding/json"
// 	"main/database"

// 	models "main/database/models"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// )

// type RegisterMessage struct {
// 	Name   string
// 	Client tgbotapi.BotAPI
// }

// func (e RegisterMessage) Run(update tgbotapi.Update) error {
// 	if update.BusinnesMessage.From.ID == update.BusinnesMessage.Chat.ID {
// 		return nil
// 	}

// 	db := database.Connect()
// 	defer db.Close()

// 	// 1. Get or create sender user
// 	fromUser := &models.TelegramUser{
// 		TgId:      update.BusinnesMessage.From.ID,
// 		FirstName: update.BusinnesMessage.From.FirstName,
// 		LastName:  update.BusinnesMessage.From.LastName,
// 		Username:  update.BusinnesMessage.From.UserName,
// 	}
// 	_, err := db.Model(fromUser).
// 		Where("tg_id = ?", fromUser.TgId).
// 		SelectOrInsert()
// 	if err != nil {
// 		return err
// 	}

// 	// 2. Get or create recipient user
// 	toUser := &models.TelegramUser{
// 		TgId:      update.BusinnesMessage.Chat.ID,
// 		FirstName: update.BusinnesMessage.Chat.FirstName,
// 		LastName:  update.BusinnesMessage.Chat.LastName,
// 		Username:  update.BusinnesMessage.Chat.UserName,
// 	}
// 	_, err = db.Model(toUser).
// 		Where("tg_id = ?", toUser.TgId).
// 		SelectOrInsert()
// 	if err != nil {
// 		return err
// 	}

// 	// 3. Get or create chat
// 	chat := &models.Chat{
// 		TgId:         update.BusinnesMessage.Chat.ID,
// 		WithUserTgId: toUser.TgId,
// 	}
// 	_, err = db.Model(chat).
// 		Where("tg_id = ?", chat.TgId).
// 		SelectOrInsert()
// 	if err != nil {
// 		return err
// 	}

// 	updateJson, err := json.Marshal(update)
// 	if err != nil {
// 		return err
// 	}

// 	msg := &models.Message{
// 		TgId: update.BusinnesMessage.MessageID,
// 		FromUserTgId: fromUser.TgId,
// 		ToUserTgId: toUser.TgId,
// 		ChatTgId: chat.TgId,
// 		Json: string(updateJson),
// 	}

// 	_, err = db.Model(msg).Insert()

// 	return err
// }

// func (e RegisterMessage) GetName() string {
// 	return e.Name
// }
