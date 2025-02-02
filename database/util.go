package database

import (
	models "main/database/models"
)

func GetOrCreateUser(tgId int64, businessConnectionId string) (models.TelegramUser, error) {
	db := Connect()
	defer db.Close()

	user := &models.TelegramUser{}
	err := db.Model(user).Where("tg_id = ?", tgId).Select()
	
	if err == nil {
		return *user, nil
	}

	user = &models.TelegramUser{
		TgId:                 tgId,
		BusinessConnectionId: businessConnectionId,
	}
	_, err = db.Model(user).Insert()
	if err != nil {
		return models.TelegramUser{}, err
	}

	settings := &models.UserSettings{
		UserId: user.Id,
	}
	_, err = db.Model(settings).Insert()
	if err != nil {
		return *user, err
	}

	return *user, nil
}
