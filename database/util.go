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

func UpdateBusinessConnectionId(user models.TelegramUser, new string) error {
	db := Connect()
	defer db.Close()

	_, err := db.Model((*models.Message)(nil)).
		Set("business_connection_id = ?", new).
		Where("business_connection_id = ?", user.BusinessConnectionId).
		Update()
	if err != nil {
		return err
	}

	user.BusinessConnectionId = new
	_, err = db.Model(&user).Column("business_connection_id").WherePK().Update()

	return err
}

