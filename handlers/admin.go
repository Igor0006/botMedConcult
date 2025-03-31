package handlers

import (
	"fmt"
	"medBot/database"
	"strings"
	"time"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
var currentDayAdm string
func HandleCallbackQueryAdmin(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
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
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, "Выберите дату", CreateMonthKeyboard(0))
	default:
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, "Выберите время", CreateAdminSchedule(currentDayAdm))
	}
    bot.Send(editmsg)
}
func AppointmentsAdmin(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	data := database.GetAppointmentsAdmin()
	text := ""
	for _, el := range data {
		t, _ := time.Parse(time.RFC3339, el[0])
		text += fmt.Sprintf("@%s в %d:00 %s %d\n", el[1], t.Hour(), months[int(t.Month())], t.Day())
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	bot.Send(msg)
}