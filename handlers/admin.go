package handlers
import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCallbackQueryAdmin(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
    //Вспылвающее сообщение при нажатии на кнопку
    callbackConfig := tgbotapi.NewCallback(callback.ID, "Вы нажали: "+callback.Data)
	if _, err := bot.Request(callbackConfig); err != nil {
		log.Println("Ошибка при обработке callback:", err)
	}
	if callback.Data == "shedule" {
		
	}
    editmsg := tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, callback.Data, CreateAdminSchedule())
    bot.Send(editmsg)
}