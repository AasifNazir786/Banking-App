package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

type DB struct {
	*sql.DB
}

func InitDB() error {
	connStr := "user=postgres password=njasm786 dbname=banking-DB sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}
	return nil
}

func CloseDB() error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *sql.DB {
	return db
}
