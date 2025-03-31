package handlers

import (
	"fmt"
	"medBot/database"
	"strconv"
	"time"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
var months = map[int]string{
	1: "Январь", 2: "Февраль", 3: "Март", 4: "Апрель", 5: "Май", 6: "Июнь", 7: "Июль", 8: "Август", 9: "Сентябрь", 10: "Октябрь", 11: "Ноябрь", 12: "Декабрь",
}
func CreateUserSchedule(date string) tgbotapi.InlineKeyboardMarkup {
	arr := database.GetFreeSlots(date)
	var keyboard [][]tgbotapi.InlineKeyboardButton
	for i := 0; i < len(arr); i++{
		parsedTime, _ := time.Parse("15:04:05", arr[i])
		keyboard = append(keyboard, []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(parsedTime.Hour()) + ":00-" + strconv.Itoa(parsedTime.Hour() + 1) + ":00", 
			"userschedule/" + arr[i])})
	}
	keyboard = append(keyboard, []tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonData("↩Назад", "backToCalendar")})
	return tgbotapi.NewInlineKeyboardMarkup(keyboard...)
}
func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
func CreateAdminSchedule(date string) tgbotapi.InlineKeyboardMarkup {
	t := time.Date(0, 0, 0, 8, 0, 0, 0, time.Now().Location())
	arr := database.GetFreeSlots(date)
	fmt.Println(arr)
	var keyboard [][]tgbotapi.InlineKeyboardButton
	for i := 0; i < 7; i++ {
		t = t.Add(1 * time.Hour)
		var keyboardrow []tgbotapi.InlineKeyboardButton

		if !contains(arr, fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())) {
			keyboardrow = append(keyboardrow, tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(t.Hour()) + ":00-" + strconv.Itoa(t.Hour() + 1) + ":00", 
			"schedule/" + strconv.Itoa(t.Hour()) + ":00:00"))
		}
		t = t.Add(1 * time.Hour)
		if !contains(arr, fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())) {
			keyboardrow = append(keyboardrow, tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(t.Hour()) + ":00-" + strconv.Itoa(t.Hour() + 1) + ":00", 
			"schedule/" + strconv.Itoa(t.Hour()) + ":00:00"))
		}
		keyboard = append(keyboard, keyboardrow)
	}
	keyboard = append(keyboard, []tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonData("↩Назад", "backToCalendar")})
	return tgbotapi.NewInlineKeyboardMarkup(keyboard...)
}
func CreateConfirmKeyboard() tgbotapi.InlineKeyboardMarkup{
	var keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✅Подтвердить", "confirm"),
			tgbotapi.NewInlineKeyboardButtonData("↩Назад", "backToTime"),
		),
	)
	return keyboard
}
func CreateUserAppointment() tgbotapi.InlineKeyboardMarkup {
	var keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("❌Отмена записи", "cancelApp"),
			tgbotapi.NewInlineKeyboardButtonData("↩Назад", "backToCalendar"),
		),
	)
	return keyboard
}
func NoAppointments() tgbotapi.InlineKeyboardMarkup {
	var keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("↩Назад", "backToCalendar"),
		),
	)
	return keyboard
}
func CreateMonthKeyboard(monthstep int) tgbotapi.InlineKeyboardMarkup {
    now := time.Now()
	year, month, _ := now.Date()

	// Первый день текущего месяца
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location()).AddDate(0, monthstep, 0)

	// Определяем день недели первого дня месяца (0 - воскресенье, 1 - понедельник, и т.д.)
	weekday := int(firstOfMonth.Weekday())
	if weekday == 0 {
		weekday = 7 // Воскресенье становится 7
	}

	prevMonthDays := (weekday - 1) % 7
	if prevMonthDays < 0 {
		prevMonthDays += 7
	}

	// Начинаем с предыдущего месяца
	currentDay := firstOfMonth.AddDate(0, 0, -prevMonthDays)

    var keyboard [][]tgbotapi.InlineKeyboardButton
	keyboard = append(keyboard, []tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonData(months[int(month) + monthstep], " ")})
    for i := 0; i < 6; i++ {
        var keyboardrow []tgbotapi.InlineKeyboardButton
        for j := 0; j < 7; j++ {
            s := fmt.Sprintf("%2d ", currentDay.Day())
			if len(database.GetFreeSlots(string(currentDay.Format("2006-01-02")))) != 0 {
				s = "| " + s + "|"
			}
            keyboardrow = append(keyboardrow, tgbotapi.NewInlineKeyboardButtonData(s, "calendar/" + string(currentDay.Format("2006-01-02"))))
            currentDay = currentDay.AddDate(0, 0, 1)

        }
        keyboard = append(keyboard, keyboardrow)
    }
	keyboard = append(keyboard, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("<<", "prevMonth"),
		tgbotapi.NewInlineKeyboardButtonData(">>", "nextMonth"),
	))
    return tgbotapi.NewInlineKeyboardMarkup(keyboard...)
}
