package main

import "database/sql"

const create string = `
	CREATE TABLE IF NOT EXISTS updates (
		UpdateId INTEGER NOT NULL PRIMARY KEY,
		Text TEXT NOT NULL,
		ChatId INTEGER NOT NULL,
		FromId INTEGER NOT NULL,
		First_Name TEXT NOT NULL,
		Last_Name TEXT,
		Type TEXT,
		Offset INTEGER,
		Length INTEGER
	);
	CREATE TABLE IF NOT EXISTS commands (
		Id INTEGER NOT NULL PRIMARY KEY,
		Name TEXT NOT NULL,
		Callback TEXT NOT NULL
	);
	INSERT INTO commands (Name, Callback) VALUES ("/echo", "EchoHandlerFunc");
	`

const file string = "updates.db"

func createTables() (error) {
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