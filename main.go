package main

import (
	"medBot/handlers"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)


func main() {
    godotenv.Load()
    // database.AddFreeSlot("2025-03-20", "15:00:00")
	bot, _ := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))

    updateConfig := tgbotapi.NewUpdate(0)
    k := handlers.CreateMonthKeyboard(0)
    updates := bot.GetUpdatesChan(updateConfig)

    for update := range updates {
        if update.Message != nil {
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
            if update.Message.Text == "open" {
                msg.ReplyMarkup = k
            }
            bot.Send(msg)
        } else if update.CallbackQuery != nil {
            handlers.HandleCallbackQuery(bot, update.CallbackQuery)
        }
    }

}
