package bot

import (
	"30DoC-Telegram-Bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"sync"
)

var (
	ChatIDSteps = map[int64]int{}
	ChatIDMutex sync.Mutex
)

func BotController() {

	config.BotInstance.Debug = true

	log.Printf("Authorized on account %s", config.BotInstance.Self.UserName)

	botUpdateConfig := tgbotapi.NewUpdate(0)
	botUpdateConfig.Timeout = 60

	updates := config.BotInstance.GetUpdatesChan(botUpdateConfig)

	for update := range updates {
		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.IsCommand() {
			// Extract the command from the Message.
			switch update.Message.Command() {
			case "help":
				HelpHandler(update)
			case "tracks":
				TracksHandler(update)
			case "start":
				StartProcess(update)
			default:
				HelpHandler(update)
			}
		} else if update.Message != nil {
			ChatIDMutex.Lock()
			chatID := ChatIDSteps[update.Message.Chat.ID]
			ChatIDMutex.Unlock()

			switch chatID {

			case 0:
				HelpHandler(update)
			case 1:
				EmailHandler(update)
			case 2:
				FullNameHandler(update)
			case 3:
				PhoneHandler(update)
			case 4:
				SchoolHandler(update)
			case 5:
				SelectTrackHandler(update)

			}
		} else if update.CallbackQuery != nil {
			SelectTrackHandler(update)
		}

		if _, err := config.BotInstance.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
