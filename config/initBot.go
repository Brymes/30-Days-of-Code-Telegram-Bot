package config

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

var BotInstance *tgbotapi.BotAPI

// FIle to Initialize Connetions to External Service
func InitBot() {
	var err error
	token := os.Getenv("TELEGRAM_API_TOKEN")

	if token == "" {
		log.Fatal("Kindly Set Telegram Bot token as environment variable with name: TELEGRAM_API_TOKEN ")
	}

	BotInstance, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Error Initializing Telegram Bot: %v", err)
	}
}
