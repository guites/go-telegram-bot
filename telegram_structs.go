package main

// ReceivedPayload is the wrapper around received telegram messages.
type ReceivedPayload struct {
	Ok bool `json:"ok"`
	Result []Update
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Message is a Telegram object that can be found in an update.
type Message struct {
	Text    string   `json:"text"`
	Chat    Chat     `json:"chat"`
	From    From     `json:"from"`
	Entity  []Entity `json:"entities"`
}

// A Telegram Chat indicates the conversation to which the message belongs.
type Chat struct {
	Id int `json:"id"`
}

// A Telegram From indicates the sender from which the message came.
type From struct {
	Id int `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name string `json:"last_name"`
}

// Message is a Telegram object that can be found in an update.
type Entity struct {
	Offset int    `json:"text"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}