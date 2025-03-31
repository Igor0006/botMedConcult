package handlers

import(
	"medBot/middleware"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
var monthstep int

func HandleCallbackQuery(bot *tgbotapi.BotAPI, callback * tgbotapi.CallbackQuery) {
	var editmsg tgbotapi.EditMessageTextConfig
	switch callback.Data {
	case "prevMonth":
		monthstep--
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, "Выберите дату", CreateMonthKeyboard(monthstep))
	case "nextMonth":
		monthstep++
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, "Выберите дату", CreateMonthKeyboard(monthstep))
	default:
		middleware.IsAdminMiddlewareCallback(HandleCallbackQueryAdmin, HandleCallbackQueryUser)(bot, callback)
	}
	bot.Send(editmsg)
}
