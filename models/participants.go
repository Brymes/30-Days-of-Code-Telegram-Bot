package models

import (
	"30DoC-Telegram-Bot/config"
	"gorm.io/gorm"
)

type Participant struct {
	ChatID   int64  `gorm:"column:chat_id;unique"`
	FullName string `gorm:"column:full_name"`
	Email    string `gorm:"column:email;unique"`
	Phone    string `gorm:"column:phone"`
	School   string `gorm:"column:school"`
	Track    string `gorm:"column:track"`
}

func (participant Participant) GetChatID() {

}

func (participant Participant) SaveEmail() string {
	result := config.DBClient.Model(&Participant{}).Create(&participant)
	if result.Error != nil {
		return "There happens to be a problem Do contact support and consider joining any of the track links at /tracks"
	}
	return ""
}

func (participant Participant) SaveName() string {
	result := config.DBClient.Model(&Participant{}).Where("chat_id = ?", participant.ChatID).Update("full_name", participant.FullName)
	return handleDBErr(result)
}

func (participant Participant) SavePhone() string {
	result := config.DBClient.Model(&Participant{}).Where("chat_id = ?", participant.ChatID).Update("phone", participant.Phone)
	return handleDBErr(result)
}

func (participant Participant) SaveSchool() string {
	result := config.DBClient.Model(&Participant{}).Where("chat_id = ?", participant.ChatID).Update("school", participant.School)
	return handleDBErr(result)
}

func (participant Participant) SaveTrack() string {
	result := config.DBClient.Model(&Participant{}).Where("chat_id = ?", participant.ChatID).Update("track", participant.Track)
	return handleDBErr(result)
}

func handleDBErr(result *gorm.DB) string {
	if result.Error != nil {
		return "There happens to be a problem Do contact support and resend last message"
	}
	return ""
}
