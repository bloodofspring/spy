package models

import "fmt"


type UserSettings struct {
	Id int64
	UserTgId int64
	User *TelegramUser `pg:"rel:has-one"`
	GetEvents bool
	SaveDeletedMessages bool
	SaveEditedMessages bool
	SaveSelfDistructingPhotos bool
	SaveSelfDistructingVideoNotes bool
	SaveSelfDistructingVoices bool
	SaveSelfDistructingVidoes bool  // ToDo: add this feature
}


func (s UserSettings) String() string {
	return fmt.Sprintf(
		"UserSettings(Id=%d, User=%s, %t, %t, %t, %t, %t, %t, %t)",
		s.Id, s.User.String(),
		s.GetEvents,
		s.SaveDeletedMessages,
		s.SaveEditedMessages,
		s.SaveSelfDistructingPhotos,
		s.SaveSelfDistructingVideoNotes,
		s.SaveSelfDistructingVoices,
		s.SaveSelfDistructingVidoes,
	)
}

