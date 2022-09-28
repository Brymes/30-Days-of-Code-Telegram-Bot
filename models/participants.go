package models

import (
	"30DoC-Telegram-Bot/config"
	"errors"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
	"log"
)

func InitModels() {
	err := config.DBClient.AutoMigrate(&Participant{})
	if err != nil {
		log.Fatalln(err)
	}
}

func LoadPreviousSteps() map[int64]int {
	var (
		participants []Participant
	)
	result := config.DBClient.Model(&Participant{}).Where("step != ?", 0).Find(&participants)

	if result.Error != nil {
		log.Fatal("Cannot Load Previous steps")
	}

	chatIdSteps := make(map[int64]int, len(participants))
	for _, elem := range participants {
		chatIdSteps[elem.ChatID] = elem.Step
	}

	return chatIdSteps
}

type Participant struct {
	ChatID   int64  `gorm:"column:chat_id;unique"`
	FullName string `gorm:"column:full_name"`
	Email    string `gorm:"column:email;unique"`
	Phone    string `gorm:"column:phone"`
	School   string `gorm:"column:school"`
	Track    string `gorm:"column:track"`
	Step     int    `gorm:"column:step"`
}

func (participant Participant) SaveEmail() string {
	result := config.DBClient.Model(&Participant{}).Create(&participant)
	return handleDBErr(result)
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
	var perr *pgconn.PgError
	errors.As(result.Error, &perr)

	if perr != nil {
		switch perr.Code {
		case "23505":
			return "This User has already registered"
		default:
			return "There happens to be a problem Do contact support and resend last message"
		}
	} else if result.Error != nil {
		return "There happens to be a problem Do contact support and resend last message"
	}
	return ""
}
