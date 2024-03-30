package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // added "_" in the beginning of the import says to go formatter although I not use this package directly but I need it so it wouldn't delete it
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "event.db")
	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10) // configuring database to have at most 10 concurrent open connections
	DB.SetMaxIdleConns(5)
}
