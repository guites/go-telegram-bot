package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func addUpdateToDatabase(db *sql.DB, newUpdate DatabaseUpdate) {
	stmt, _ := db.Prepare("INSERT INTO updates (UpdateId, Text, ChatId, FromId, First_Name, Last_Name) VALUES (?, ?, ?, ?, ?, ?)")
	res, err := stmt.Exec(newUpdate.UpdateId, newUpdate.Text, newUpdate.ChatId, newUpdate.FromId, newUpdate.First_Name, newUpdate.Last_Name)
	if err != nil {
		log.Fatal((err.Error()))
	}
	defer stmt.Close()
	log.Print(res)
}