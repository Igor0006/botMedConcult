package handlers

import (
	"log"
	"medBot/database"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
var currentDayAdm string
func HandleCallbackQueryAdmin(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
    //Вспылвающее сообщение при нажатии на кнопку
    callbackConfig := tgbotapi.NewCallback(callback.ID, "")
	if _, err := bot.Request(callbackConfig); err != nil {
		log.Println("Ошибка при обработке callback:", err)
	}
	if strings.Split(callback.Data, "/")[0] == "calendar" {
		currentDayAdm = strings.Split(callback.Data, "/")[1]
	} else if strings.Split(callback.Data, "/")[0] == "schedule" {
		database.AddFreeSlot(currentDayAdm, strings.Split(callback.Data, "/")[1])
	} 
	var editmsg tgbotapi.EditMessageTextConfig
	switch callback.Data {
	case "backToCalendar":
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, callback.Data, CreateMonthKeyboard(0))
	default:
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, callback.Data, CreateAdminSchedule())
	}
    bot.Send(editmsg)
}