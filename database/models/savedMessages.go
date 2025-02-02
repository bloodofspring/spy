package models

import "fmt"


type Message struct {
	Id int64
	TgId int64 `pg:",pk"`
	FromUserTgId int64
	From *TelegramUser `pg:"rel:has-one"`
	ToUserTgId int64
	To *TelegramUser `pg:"rel:has-one"`
	ChatTgId int64
	Chat *Chat `pg:"rel:has-one"`
	Json string
}

func (m Message) String() string {
	return fmt.Sprintf("Message(TgId=%d)", m.TgId)
}