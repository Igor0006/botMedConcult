package handlers

import(
	"medBot/middleware"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCallbackQuery(bot *tgbotapi.BotAPI, callback * tgbotapi.CallbackQuery) {
	middleware.IsAdminMiddleware(HandleCallbackQueryAdmin, HandleCallbackQueryUser)(bot, callback)
}