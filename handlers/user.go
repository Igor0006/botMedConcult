package handlers

import (
	"log"
	"medBot/database"
	"strings"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var currentDate string
var currentTime string
func HandleCallbackQueryUser(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	var editmsg tgbotapi.EditMessageTextConfig
	callbackConfig := tgbotapi.NewCallback(callback.ID, "")
	if _, err := bot.Request(callbackConfig); err != nil {
		log.Println("Ошибка при обработке callback:", err)
	}
	if strings.Split(callback.Data, "/")[0] == "calendar" {
		currentDate = strings.Split(callback.Data, "/")[1]
		fmt.Println(currentDate)
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, "Выберите время", CreateUserSchedule(currentDate))
	} else if strings.Split(callback.Data, "/")[0] == "userschedule"{
		currentTime = strings.Split(callback.Data, "/")[1]
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, "Выберите время", CreateConfirmKeyboard())
	}
	switch callback.Data {
	case "backToCalendar":
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, "Выберите дату", CreateMonthKeyboard(0))
	case "backToTime":
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, "Выберите время", CreateUserSchedule(currentDate))
	case "confirm":
		database.TakeTheTime(currentDate, currentTime)
		database.MakeAppointment(currentDate + " " + currentTime, int(callback.From.ID), callback.From.UserName)
		editmsg = tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, "Выберите дату", CreateMonthKeyboard(0))

	}
	database.GetFreeSlots(callback.Data)
    bot.Send(editmsg)
}