package models

import "fmt"


type Message struct {
	Id int64
	TgId int64 `pg:"pk"`
	From *TelegramUser `pg:"rel:has-one"`
	To *TelegramUser `pg:"rel:has-one"`
	Chat *Chat `pg:"rel:has-one"`
	Json string
}

func (m Message) String() string {
	return fmt.Sprintf("Message(TgId=%d)", m.TgId)
}
