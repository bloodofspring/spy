package models

import "fmt"


type TelegramUser struct {
	TgId int64 `pg:",pk"`
	FirstName string
	LastName string
	Username string
}

func (u TelegramUser) String() string {
	return fmt.Sprintf("TelegramUser(ID=%d, FullName=%s, Username=%s)", u.TgId, u.FirstName + " " + u.LastName, u.Username)
}


type BotPeer struct {
	Id int64
	UserTgId int64
	User *TelegramUser `pg:"rel:has-one"`
}

func (p BotPeer) String() string {
	return fmt.Sprintf("BotPeer(Id=%d, User=%s)", p.Id, p.User.String())
}
