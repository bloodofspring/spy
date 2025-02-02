package models

import "fmt"

type Message struct {
	TgId       int `pg:",pk"`
	ChatId     int64
	FromUserId int64
	Text       string
}

func (m Message) String() string {
	return fmt.Sprintf("Message(TgId=%d)", m.TgId)
}
