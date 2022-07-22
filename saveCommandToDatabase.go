package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func addCommandToDatabase(db *sql.DB, newCommand DatabaseCommand) {
	stmt, _ := db.Prepare("INSERT INTO commands (Name, Callback) VALUES (?, ?)")

	res, err := stmt.Exec(
		newCommand.Name,
		newCommand.Callback,
	)

	if err != nil {
		log.Fatal((err.Error()))
	}
	defer stmt.Close()
	log.Print(res)
}