package bot

import BotAPILib "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var TracksKeyboard = BotAPILib.NewInlineKeyboardMarkup(
	BotAPILib.NewInlineKeyboardRow(
		BotAPILib.NewInlineKeyboardButtonData("Frontend(React)", "Frontend(React)"),
		BotAPILib.NewInlineKeyboardButtonData("Mobile(Flutter)", "Mobile(Flutter)"),
	),
	BotAPILib.NewInlineKeyboardRow(
		BotAPILib.NewInlineKeyboardButtonData("Python(Backend)", "Python(Backend)"),
		BotAPILib.NewInlineKeyboardButtonData("Python(Machine Learning)", "Python(Machine Learning)"),
	),
	BotAPILib.NewInlineKeyboardRow(
		BotAPILib.NewInlineKeyboardButtonData("Vanilla(Frontend)", "Vanilla(Frontend)"),
		BotAPILib.NewInlineKeyboardButtonData("Node.JS", "Node.JS"),
	),
	BotAPILib.NewInlineKeyboardRow(
		BotAPILib.NewInlineKeyboardButtonData("Game Development(Unity)", "Game Development(Unity)"),
	),
)

var TrackGroupLinks = map[string]string{
	"Frontend(React)":          "https://t.me/+H7jmlynMN38zZjU0",
	"Mobile(Flutter)":          "https://t.me/+tPAv0HtlNaw3Nzc0",
	"Python(Backend)":          "https://t.me/+hsmeMbvXlzsyOTU8",
	"Python(Machine Learning)": "https://t.me/+-XWmwQ_druY4YWE0",
	"Vanilla(Frontend)":        "https://t.me/+-vFNXPbm0IwxYjI0",
	"Node.JS":                  "https://t.me/+ZvAKAm7WdI8wYmQ0",
	"Game Development(Unity)":  "https://t.me/+QZY6tLO3Vi5hNGI0",
}
