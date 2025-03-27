package handlers

import (
	"fmt"
	"log"
	"medBot/database"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var currentDateString string
var currentTimeString string
func HandleCallbackQueryUser(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	floatMessage:= " "
	var editmsg tgbotapi.EditMessageTextConfig
	if strings.Split(callback.Data, "/")[0] == "calendar" {
		currentDateString = strings.Split(callback.Data, "/")[1]
		// fmt.Println(currentDate)
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, "Выберите время", CreateUserSchedule(currentDateString))
	} else if strings.Split(callback.Data, "/")[0] == "userschedule"{
		currentTimeString = strings.Split(callback.Data, "/")[1]
		t, _ := time.Parse("15:04:05", currentTimeString)
		d, _ := time.Parse("2006-01-02", currentDateString)
		s := fmt.Sprintf("Вы уверены что хотите записатья на %d:00 %s %d", t.Hour(), months[int(d.Month())], d.Day() )
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, s, CreateConfirmKeyboard())
	}
	switch callback.Data {
	case "backToCalendar":
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, "Выберите дату", CreateMonthKeyboard(0))
	case "backToTime":
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, "Выберите время", CreateUserSchedule(currentDateString))
	case "confirm":
		if database.CanMakeAppointment(int(callback.From.ID)) {
			database.TakeTheTime(currentDateString, currentTimeString)
			database.MakeAppointment(currentDateString + " " + currentTimeString, int(callback.From.ID), callback.From.UserName)
			floatMessage = "Запись успешно создана"
		} else {
			floatMessage = "Вы не можете записаться повторно"
		}
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, "Выберите дату", CreateMonthKeyboard(0))
	case "cancelApp":
		database.DeleteAppointment(int(callback.From.ID))
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, "Выберите дату", CreateMonthKeyboard(0))
	}


	callbackConfig := tgbotapi.NewCallback(callback.ID, floatMessage)
	if _, err := bot.Request(callbackConfig); err != nil {
		log.Println("Ошибка при обработке callback:", err)
	}
    bot.Send(editmsg)
}
func HandleCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	var msg tgbotapi.MessageConfig
	switch update.Message.Command() {
	case "start":
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите дату")
		msg.ReplyMarkup = CreateMonthKeyboard(0)
	case  "getAppointment":
		userid := int(update.Message.From.ID)

		t := database.GetAppointmentUser(userid)
		if t.Year() == 1 {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "У вас нету записей")
			msg.ReplyMarkup = NoAppointments()
		} else {
			s := fmt.Sprintf("Ваша запись на %d:00 %s %d", t.Hour(), months[int(t.Month())], t.Day())
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, s)
			msg.ReplyMarkup = CreateUserAppointment()
		}
	}
	bot.Send(msg)
}