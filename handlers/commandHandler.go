package handlers

import (

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	var msg tgbotapi.MessageConfig
	switch update.Message.Command() {
	case "start":
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyMarkup = CreateMonthKeyboard(0)
	}
	bot.Send(msg)
}