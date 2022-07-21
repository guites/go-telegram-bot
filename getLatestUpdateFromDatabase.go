package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func getLatestUpdateFromDatabase (db *sql.DB) (DatabaseUpdate, error){
	var latestUpdate DatabaseUpdate
	err := db.
	QueryRow("SELECT * FROM updates ORDER BY UpdateId desc;").
	Scan(
		&latestUpdate.UpdateId,
		&latestUpdate.Text,
		&latestUpdate.ChatId,
		&latestUpdate.FromId,
		&latestUpdate.First_Name,
		&latestUpdate.Last_Name,
		&latestUpdate.Type,
		&latestUpdate.Offset,
		&latestUpdate.Length,
	)
	return latestUpdate, err
}