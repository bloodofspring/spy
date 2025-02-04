package models

import (
	"fmt"
)

type TelegramUser struct {
	TgId int64  `pg:",pk"`
	BusinessConnectionId string
}

func (p TelegramUser) String() string {
	return fmt.Sprintf("BotPeer(Id=%d)", p.TgId)
}

type Admin struct {
	Id             int
	UserTgId       int
	User           *TelegramUser `pg:"rel:has-one"`
	PermissionsLvl int           `pg:"default:1"`
}

func (a Admin) String() string {
	return fmt.Sprintf("Admin(UserId=%d, Lvl=%d)", a.User.TgId, a.PermissionsLvl)
}
