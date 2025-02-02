package models

import "fmt"


type Chat struct {
	TgId int64 `pg:"pk"`
	WithUser *TelegramUser `pg:"rel:has-one"`
}

func (c Chat) String() string {
	return fmt.Sprintf("Chat(TgId=%d, With=%s)", c.TgId, c.WithUser.String())
}