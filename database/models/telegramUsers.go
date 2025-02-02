package models

import (
	"fmt"
)

type TelegramUser struct {
	Id   int
	TgId int64
	BusinessConnectionId string
}

func (p TelegramUser) String() string {
	return fmt.Sprintf("BotPeer(Id=%d)", p.TgId)
}

type Admin struct {
	Id             int
	UserId         int
	User           *TelegramUser `pg:"rel:has-one"`
	PermissionsLvl int           `pg:"default:1"`
}

func (a Admin) String() string {
	return fmt.Sprintf("Admin(UserId=%d, Lvl=%d)", a.User.TgId, a.PermissionsLvl)
}
