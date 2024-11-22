package storage

import (
	"database/sql"
	"errors"
)

var db *sql.DB

func InitDB() error {
	connStr := "name=postgres password=njasm786 dbname=banking-DB sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return errors.New("database connection is not alive")
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
