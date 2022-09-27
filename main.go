package main

import (
	"30DoC-Telegram-Bot/bot"
	"30DoC-Telegram-Bot/config"
)

func init() {
	config.InitDb()
	config.InitBot()
}

func main() {
	bot.BotController()
}
