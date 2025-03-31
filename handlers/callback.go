package handlers

import(
	"medBot/middleware"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCallbackQuery(bot *tgbotapi.BotAPI, callback * tgbotapi.CallbackQuery) {
	callbackConfig := tgbotapi.NewCallback(callback.ID, "")
	bot.Request(callbackConfig)
	var editmsg tgbotapi.EditMessageTextConfig
	switch callback.Data {
	case "prevMonth":
		
	}
	middleware.IsAdminMiddlewareCallback(HandleCallbackQueryAdmin, HandleCallbackQueryUser)(bot, callback)
}
