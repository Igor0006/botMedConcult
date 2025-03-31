package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	pq "github.com/lib/pq"
)


func getDb() *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", 
	os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	db, _ := sql.Open("postgres", connStr)
	err := db.Ping()
	if err != nil {
		log.Fatal("Ошибка проверки подключения:", err)
	}
	return db
}
func HasFreeSlots(date string) bool{
	db := getDb()
	defer db.Close()
	var f bool
	db.QueryRow("SELECT EXISTS (SELECT 1 FROM schedule WHERE date=$1)", date).Scan(&f)
	return f
}

func GetFreeSlots(date string) []string {
	db := getDb()
	defer db.Close()
	clean()
	query := `SELECT free_slots FROM schedule where date = $1`
	arr := []string{}
	db.QueryRow(query, date).Scan(pq.Array(&arr))
	return arr	
}
func clean() {
	db := getDb()
	defer db.Close()
	now := time.Now().Format("2006-01-02")
	db.Exec(`DELETE FROM schedule WHERE date < $1`, now)
	now = time.Now().Add(-time.Hour).Format("2006-01-02 15:04:05")
	db.Exec(`DELETE FROM appointments WHERE time < $1`, now)

}
func TakeTheTime(date string, time string) {
	db := getDb()
	defer db.Close()
	query1 := `
		UPDATE schedule
		SET free_slots = array_remove(free_slots, $1)
		WHERE date = $2
	`
	db.Exec(query1, time, date)
	query2 := `
	UPDATE schedule 
	SET occupied_slots = array_append(occupied_slots, $1::TIME)
	WHERE date = $2`
	db.Exec(query2, time, date)
}
func MakeAppointment(time string, userid int, contact string) {
	db := getDb()
	defer db.Close()
	query := `
		INSERT INTO appointments (time, contact, userid)
		VALUES($1, $2, $3)
	`
	db.Exec(query, time, contact, userid)
}
func GetAppointmentUser(userid int) time.Time{
	db := getDb()
	defer db.Close()
	query := `
	SELECT time From appointments
	WHERE userid = $1
	`
	var timeString string
	db.QueryRow(query, userid).Scan(&timeString)
	parsedTime, _ := time.Parse("2006-01-02T15:04:05Z07:00", timeString)
	return parsedTime
}
func GetAppointmentsAdmin() [][]string{
	db := getDb()
	defer db.Close()
	rows, _ := db.Query(`SELECT time, contact FROM appointments`)
	data := [][]string{}
	for rows.Next() {
		var timeString string
		var username string
		rows.Scan(&timeString, &username)
		data = append(data, []string{timeString, username})
	}
	return data
}
func DeleteAppointment(userid int) {
	db := getDb()
	defer db.Close()
	var timeString string
	db.QueryRow(`SELECT time FROM appointments WHERE userid = $1`, userid).Scan(&timeString)
	parsedTime, _ := time.Parse("2006-01-02T15:04:05Z07:00", timeString)
	time := parsedTime.Format("15:04:05")
	date := parsedTime.Format("2006-01-02")
	AddFreeSlot(date, time)
	db.Exec(`DELETE FROM appointments WHERE userid = $1`, userid)
}
func CanMakeAppointment(userid int) bool{
	db := getDb()
	defer db.Close()
	query := `
	SELECT EXISTS (SELECT 1 FROM appointments where userid = $1)
	`
	var f bool
	db.QueryRow(query, userid).Scan(&f)
	return !f
}
func AddFreeSlot(date string, time string) {
	db := getDb()
    defer db.Close()
	
	if HasFreeSlots(date) {
		query := `
		UPDATE schedule
		SET free_slots = array_append(free_slots, $1::TIME)
		WHERE date = $2`
		db.Exec(query, time, date)
		query = `UPDATE schedule
		SET occupied_slots = array_remove(occupied_slots, $1)
		WHERE date = $2`
		db.Exec(query, time, date)
	} else {
		freeSlots := []string{time}
		occupiedSlots := []string{}
		freeSlotsArray := pq.StringArray(freeSlots)
		occupiedSlotsArray := pq.StringArray(occupiedSlots)
		query := `
		INSERT INTO schedule (date, free_slots, occupied_slots)
		VALUES ($1, $2, $3)`
		db.Exec(query, date, freeSlotsArray, occupiedSlotsArray)
	}
	fmt.Println("Free slot was added", date, time)
}