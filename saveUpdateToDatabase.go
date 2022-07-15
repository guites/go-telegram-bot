package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)


func addUpdateToDatabase(db *sql.DB, newUpdate DatabaseUpdate) {
	stmt, _ := db.Prepare("INSERT INTO updates (UpdateId, Text, ChatId, FromId, First_Name, Last_Name) VALUES (?, ?, ?, ?, ?, ?)")
	stmt.Exec(nil, newUpdate.UpdateId, newUpdate.Text, newUpdate.ChatId, newUpdate.FromId, newUpdate.First_Name, newUpdate.Last_Name)
	defer stmt.Close()
}