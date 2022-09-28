package bot

import (
	"30DoC-Telegram-Bot/config"
	"30DoC-Telegram-Bot/models"
	"fmt"
	BotAPILib "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/mail"
)

func TracksHandler(update BotAPILib.Update) {
	text := `
	Kindly Find the Links to our Tracks below
	
	Frontend(React):          https://t.me/+H7jmlynMN38zZjU0 ,
	Mobile(Flutter):          https://t.me/+tPAv0HtlNaw3Nzc0 ,
	Python(Backend):          https://t.me/+hsmeMbvXlzsyOTU8 ,
	Python(Machine Learning): https://t.me/+-XWmwQ_druY4YWE0 ,
	Vanilla(Frontend):        https://t.me/+-vFNXPbm0IwxYjI0 ,
	Node.JS:                  https://t.me/+ZvAKAm7WdI8wYmQ0 ,
	Game Development(Unity):  https://t.me/+QZY6tLO3Vi5hNGI0 ,
	`
	sendText(update, text)

}

func HelpHandler(update BotAPILib.Update) {
	text := `
	Hi Welcome to 30 Days of Code,
	
	We have the following tracks
	
	Frontend(React),
	Mobile(Flutter),
	Python(Backend),
	Python(Machine Learning),
	Vanilla(Frontend),
	Node.JS,
	Game Development(Unity),
	
	To register kindly run /register .
	`

	sendText(update, text)
}

func StartProcess(update BotAPILib.Update) {
	text := `Kindly send your Email Address`
	EditChatIDSteps(update.Message.Chat.ID, 1)
	sendText(update, text)
}

func EmailHandler(update BotAPILib.Update) {
	_, err := mail.ParseAddress(update.Message.Text)
	if err != nil {
		sendText(update, "Invalid Email Address supplied. Please enter a valid email address")
		return
	}

	message := models.Participant{
		Step:   2,
		ChatID: update.Message.Chat.ID,
		Email:  update.Message.Text,
	}.SaveEmail()

	if message == "This User has already registered" {
		if ChatIDSteps[update.Message.Chat.ID] >= 1 {
			EditChatIDSteps(update.Message.Chat.ID, 2)
			sendText(update, "Kindly enter your Full Name")
			return
		} else {
			sendText(update, message)
			return
		}
	} else if message != "" {
		sendText(update, message)
		return
	}

	EditChatIDSteps(update.Message.Chat.ID, 2)
	sendText(update, "Kindly enter your Full Name")
}

func FullNameHandler(update BotAPILib.Update) {
	message := models.Participant{
		Step:     3,
		ChatID:   update.Message.Chat.ID,
		FullName: update.Message.Text,
	}.SaveName()

	if message != "" {
		sendText(update, message)
		return
	}

	EditChatIDSteps(update.Message.Chat.ID, 3)
	sendText(update, "Kindly enter your Phone Number")
}

func PhoneHandler(update BotAPILib.Update) {
	length := len(update.Message.Text)
	if length < 11 || length > 15 {
		sendText(update, "Invalid Phone Number Supplied. Please enter a valid phone number")
		return
	}

	message := models.Participant{
		Step:   4,
		ChatID: update.Message.Chat.ID,
		Phone:  update.Message.Text,
	}.SavePhone()

	if message != "" {
		sendText(update, message)
		return
	}

	EditChatIDSteps(update.Message.Chat.ID, 4)
	sendText(update, "Kindly enter the name of your School")
}

func SchoolHandler(update BotAPILib.Update) {
	message := models.Participant{
		Step:   5,
		ChatID: update.Message.Chat.ID,
		School: update.Message.Text,
	}.SaveSchool()

	if message != "" {
		sendText(update, message)
		return
	}

	EditChatIDSteps(update.Message.Chat.ID, 5)
	sendSpecialKeyboard(update)
}

func SelectTrackHandler(update BotAPILib.Update) {
	if update.CallbackQuery == nil {
		sendSpecialKeyboard(update)
	}

	log.Println(update.CallbackData())
	message := models.Participant{
		Step:   6,
		ChatID: update.CallbackQuery.Message.Chat.ID,
		Track:  update.CallbackData(),
	}.SaveTrack()

	if message != "" {
		sendText(update, message)
		return
	}

	EditChatIDSteps(update.CallbackQuery.Message.Chat.ID, 6)
	text := fmt.Sprintf("Your information has been saved. \n Kindly join the group chat for your selected Track \n %v", TrackGroupLinks[update.CallbackData()])

	msg := BotAPILib.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.CallbackQuery.Message.MessageID

	_, err := config.BotInstance.Send(msg)

	if err != nil {
		log.Println(err)
	}
}
