package handlers

import (
	"medBot/lexicon"
	"medBot/middleware"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	var msg tgbotapi.MessageConfig
	switch update.Message.Text {
	case lexicon.Lex["init"]:
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите дату")
		msg.ReplyMarkup = CreateMonthKeyboard(0)
	case lexicon.Lex["getAppointment"]:
		middleware.IsAdminMiddlewareUpdate(AppointmentsAdmin, AppointmentsUser)(bot, update)
	}
	bot.Send(msg)
}