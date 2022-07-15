package main

// A Telegram From indicates the sender from which the message came.
type From struct {
	Id int `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name string `json:"last_name"`
}
