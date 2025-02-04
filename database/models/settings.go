package models

import "fmt"

type UserSettings struct {
	Id                            int
	UserId                        int
	User                          *TelegramUser `pg:"rel:has-one"`
	GetEvents                     bool          `pg:"default:true"`
	SaveDeletedMessages           bool          `pg:"default:true"`
	SaveEditedMessages            bool          `pg:"default:true"`
	SaveSelfDistructingPhotos     bool          `pg:"default:true"`
	SaveSelfDistructingVideoNotes bool          `pg:"default:true"`
	SaveSelfDistructingVoices     bool          `pg:"default:true"`
	SaveSelfDistructingVideos     bool          `pg:"default:true"` // ToDo: add this feature
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
		s.SaveSelfDistructingVideos,
	)
}
