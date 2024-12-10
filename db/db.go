package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	db.Close()
}

func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	)`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}
