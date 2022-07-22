package main

// a DatabaseUpdate horizontalizes all fields from a sent/received telegram update
type DatabaseUpdate struct {
	UpdateId   int
	Text       string
	ChatId     int
	FromId     int
	First_Name string
	Last_Name  string
	Type       string
	Offset     int
	Length     int
}


// a DatabaseCommands relates a command name to a callback function
type DatabaseCommand struct {
	Id       int
	Name     string
	Callback string
}