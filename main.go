package main

import (
    "database/sql"
	"medBot/handlers"
	"os"
    "fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
    _ "github.com/lib/pq"
)


func main() {
    godotenv.Load()
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    "locqlhost", 5432, os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
defer db.Close()
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
