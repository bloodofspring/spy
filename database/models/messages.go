package models

import "fmt"

type Message struct {
	TgId       int `pg:",pk"`
	ChatId     int64
	FromUserId int64
	FromUser   *TelegramUser `pg:"rel:has-one"`
}

func (m Message) String() string {
	return fmt.Sprintf("Message(TgId=%d)", m.TgId)
}
