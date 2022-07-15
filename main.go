package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

func pollCallbacFn() {
	log.Printf("Callback function was called!")
}

type Chat struct {
	Id int `json:"id"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}


// parseTelegramResponse handles incoming update from the Telegram
func parseTelegramResponse(r *http.Response) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}
	return &update, nil
}


// longPollingHandler starts the request to the telegram API and executes the callback if the request finishes with updates
func longPollingHandler(timeout int) {

	// starts the polling request
	log.Printf("Attempting to start long poll with timeout %d...", timeout)

	var telegramPollUrl string = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_TOKEN") + "/getUpdates?timeout=" + strconv.Itoa(timeout) + "&offset=435559519"

	log.Printf("Telegram poll url: %s", telegramPollUrl)

	response, poll_err := http.Get(
		telegramPollUrl,
	)
	if poll_err != nil {
		log.Printf("Error polling the telegram server, %s", poll_err.Error())
		return
	}

	// parse incoming request
	var update, err = parseTelegramResponse(response)
	if err != nil {
		log.Printf("error parsing the telegram response, %s", err.Error())
		return
	}
	log.Printf("Received update from telegram: %s", update.Message.Text)
	// // sanitize input
	// var sanitizedSeed = sanitize(update.Message.Text)

	// // TODO: add the handling logic
	// var responseText string = "Test response! Beep."

	// // send a message back to telegram
	// var telegramResponseBody, errTelegram = sendTextToTelegramChat(update.Message.Chat.Id, responseText)
	// if errTelegram != nil {
	// 	log.Printf("got error %s from telegram, response body is %s", errTelegram.Error(), telegramResponseBody)
	// } else {
	// 	log.Printf("message %s successfuly distributed to chat id %d", responseText, update.Message.Chat.Id)
	// }
}


func main() {
	longPollingHandler(60)
}