package database

import (
	"database/sql"
	"fmt"
	"log"

	- "github.com/lib/pq"
)

var DB *sql.DB

func Open() error {
	var err error
	connStr := "user=postgres password=postgres dbname=budget_manager sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Errorf("Failed to connect to the database")
	}

	if err = DB.Ping(); err != nil {
		fmt.Errorf("Failed to ping the database")
	}

	log.Println("Connected to the database")
	return nil
}

func Close() error {
	err := DB.Close()
	if err != nil {
		fmt.Errorf("Failed to close the database connection")
	}

	log.Println("Closed the database connection")
	return nil
}
