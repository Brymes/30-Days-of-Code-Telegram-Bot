package bot

import (
	"30DoC-Telegram-Bot/config"
	"30DoC-Telegram-Bot/models"
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
	models.InitModels()
	ChatIDSteps = models.LoadPreviousSteps()

	log.Printf("Authorized on account %s", config.BotInstance.Self.UserName)

	botUpdateConfig := tgbotapi.NewUpdate(0)
	botUpdateConfig.Timeout = 60

	updates := config.BotInstance.GetUpdatesChan(botUpdateConfig)

	for update := range updates {
		//Handle Bot Restart and Stop
		if update.CallbackQuery != nil {
			go SelectTrackHandler(update)
			continue
		} else if update.Message == nil {
			EditChatIDSteps(update.MyChatMember.Chat.ID, 0)
			continue
		} else if update.Message.Text == "/start" {
			go HelpHandler(update)
			continue
		}

		if update.Message.IsCommand() {
			// Extract the command from the Message.
			switch update.Message.Command() {
			case "help":
				go HelpHandler(update)
				continue
			case "tracks":
				go TracksHandler(update)
				continue
			case "start":
				go HelpHandler(update)
				continue
			case "register":
				StartProcess(update)
			default:
				go HelpHandler(update)
				continue
			}
		} else if update.Message != nil {
			ChatIDMutex.Lock()
			chatID := ChatIDSteps[update.Message.Chat.ID]
			ChatIDMutex.Unlock()

			switch chatID {

			case 0:
				go HelpHandler(update)
				continue
			case 1:
				go EmailHandler(update)
				continue
			case 2:
				go FullNameHandler(update)
				continue
			case 3:
				go PhoneHandler(update)
				continue
			case 4:
				go SchoolHandler(update)
				continue
			case 5:
				go SelectTrackHandler(update)
				continue

			}
		}
	}
}
