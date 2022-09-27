package bot

import (
	"30DoC-Telegram-Bot/config"
	BotAPILib "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func sendText(update BotAPILib.Update, text string) {
	msg := BotAPILib.NewMessage(update.Message.Chat.ID, "")
	msg.ReplyToMessageID = update.Message.MessageID

	msg.Text = text
	_, err := config.BotInstance.Send(msg)

	if err != nil {
		log.Println(err)
	}
}
func sendSpecialKeyboard(update BotAPILib.Update) {
	msg := BotAPILib.NewMessage(update.Message.Chat.ID, "Kindly select your preferred track")

	msg.ReplyMarkup = TracksKeyboard

	// Send the message.
	if _, err := config.BotInstance.Send(msg); err != nil {
		log.Println(err)
	}
}

func EditChatIDSteps(chatID int64, step int) {
	ChatIDMutex.Lock()
	if step == 6 {
		delete(ChatIDSteps, chatID)
	} else {
		ChatIDSteps[chatID] = step
	}
	ChatIDMutex.Unlock()
}
