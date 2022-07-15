package main

import "database/sql"

const create string = `
  CREATE TABLE IF NOT EXISTS updates (
  UpdateId INTEGER NOT NULL PRIMARY KEY,
  Text TEXT NOT NULL,
  ChatId INTEGER NOT NULL,
  FromId INTEGER NOT NULL,
  First_Name TEXT NOT NULL,
  Last_Name TEXT
);`

const file string = "updates.db"

func createUpdatesTable() (error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return err
	}
	defer db.Close()
	if _, err := db.Exec(create); err != nil {
		return err
	}
	return nil
}