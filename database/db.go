package database

import (
	"database/sql"
	"fmt"
	"os"
	"log"
	pq "github.com/lib/pq"
)

func HasFreeSlots(date string) bool{
	connStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", 
	os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	db, _ := sql.Open("postgres", connStr)
	var f bool
	db.QueryRow("SELECT EXISTS (SELECT 1 FROM schedule WHERE date=$1)", date).Scan(&f)
	return f
}

func AddFreeSlot(date string, time string) {
	fmt.Println(date, time)
	connStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", 
	os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	db, _ := sql.Open("postgres", connStr)
	err := db.Ping()
	if err != nil {
		log.Fatal("Ошибка проверки подключения:", err)
	}
    defer db.Close()
	
	if HasFreeSlots(date) {
		query := `
		UPDATE schedule
		SET free_slots = array_append(free_slots, $1::TIME)
		WHERE date = $2`
		_, err := db.Exec(query, time, date)
		if err != nil {
			log.Fatal(err)
		}
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
}