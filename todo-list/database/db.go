package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

var DB *sql.DB // Global DB variable

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:vijay@tcp(localhost:3306)/to_do")
	if err != nil {
		log.Fatal(err)
	}

	// Test the database connection
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}
