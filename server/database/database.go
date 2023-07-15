package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var dbConnection *sql.DB

func Connect(fileName string) {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
	dbConnection = db
}

func GetConnection() *sql.DB {
	if dbConnection == nil {
		log.Fatal("no database connection found")
	}
	return dbConnection
}
