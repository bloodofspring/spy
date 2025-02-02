package database

import (
	"main/database/models"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)


func InitDb(db *pg.DB) error {
	models := []interface{}{
		(*models.TelegramUser)(nil),
		(*models.BotPeer)(nil),
		(*models.UserSettings)(nil),
		(*models.Message)(nil),
		(*models.Chat)(nil),
    }

    for _, model := range models {
        err := db.Model(model).CreateTable(&orm.CreateTableOptions{
            Temp: false, // Временные таблицы
			IfNotExists: true,
        })
        if err != nil {
            return err
        }
    }
    return nil
}