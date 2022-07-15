package main

// a DatabaseUpdate horizontalizes all fields from a sent/received telegram update
type DatabaseUpdate struct {
	UpdateId   int
	Text       string
	ChatId     int
	FromId     int
	First_Name string
	Last_Name  string
}