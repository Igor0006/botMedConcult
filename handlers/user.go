package handlers

import (
	"fmt"
	"log"
	"medBot/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCallbackQueryUser(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
    //Вспылвающее сообщение при нажатии на кнопку
    callbackConfig := tgbotapi.NewCallback(callback.ID, "Вы нажали: "+callback.Data)
	if _, err := bot.Request(callbackConfig); err != nil {
		log.Println("Ошибка при обработке callback:", err)
	}
	fmt.Println(callback.Data)
	database.GetFreeSlots(callback.Data)
    editmsg := tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, callback.Data, CreateMonthKeyboard(0))
    bot.Send(editmsg)
}