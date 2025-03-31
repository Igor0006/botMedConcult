package handlers

import (
	"medBot/lexicon"
	"medBot/middleware"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
func HandleCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	var msg tgbotapi.MessageConfig
	switch update.Message.Command() {
	case "start":
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, lexicon.Lex["greeting"])
		msg.ReplyMarkup = Menu
	case  "getAppointment":
		middleware.IsAdminMiddlewareUpdate(AppointmentsAdmin, AppointmentsUser)(bot, update)
	}
	bot.Send(msg)
}	