package models

import "fmt"

type Message struct {
	TgId         int `pg:",pk"`
	BusinessConnectionId string
	FromUserTgId int64
	ToUserTgId   int64
	ToUser       *TelegramUser `pg:"rel:has-one"`
	Text         string
}

func (m Message) String() string {
	return fmt.Sprintf("Message(TgId=%d)", m.TgId)
}
