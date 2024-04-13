package main

import (
	"log"

	"database/sql"
)

var db *sql.DB

func connectDB() {
	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Println("error opening sql: " + err.Error())
	}
}
