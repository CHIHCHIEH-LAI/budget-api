package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Open(config Config) error {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Name)
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open the database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping the database: %v", err)
	}

	log.Println("Connected to the database")
	return nil
}

func Close() error {
	err := DB.Close()
	if err != nil {
		return fmt.Errorf("failed to close the database connection")
	}

	log.Println("Closed the database connection")
	return nil
}
