package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}

	return db.Ping()
}

func CreatePillarTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS pillars(
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT NOT NULL,
		"description" TEXT
	);`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Pillars Table Created")
}
