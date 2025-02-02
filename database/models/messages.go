package models

import "fmt"

type Message struct {
	TgId       int `pg:",pk"`
	BusinessConnectionId string
	Text       string
}

func (m Message) String() string {
	return fmt.Sprintf("Message(TgId=%d)", m.TgId)
}
