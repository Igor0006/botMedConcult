package main

import (
	"medBot/handlers"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)


func main() {
    godotenv.Load()
	bot, _ := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
    updateConfig := tgbotapi.NewUpdate(0)
    updates := bot.GetUpdatesChan(updateConfig)

    for update := range updates {
        if update.CallbackQuery != nil {
            handlers.HandleCallbackQuery(bot, update.CallbackQuery)
        }else if update.Message.IsCommand() {
            handlers.HandleCommand(bot, &update)
        } else {
            handlers.HandleMessage(bot, &update)
        }
    }

}
