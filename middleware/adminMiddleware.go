package middleware

import (
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"strconv"
)

func IsAdminMiddleware(nextAdmin func(botAPI *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery), 
					nextUser func(botAPI *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery)) func 
(botAPI *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {
	return func(botAPI *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {
		if strconv.Itoa(int(callbackQuery.From.ID)) == os.Getenv("ADMIN_ID") {
			nextAdmin(botAPI, callbackQuery)
		} else {
			nextUser(botAPI, callbackQuery)
		}
	}
}
