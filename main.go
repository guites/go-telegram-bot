package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

// longPollingHandler starts the request to the telegram API and executes the callback if the request finishes with updates
func longPollingHandler(timeout int, offset int, telegram_bot_token string) (response *http.Response) {

	// starts the polling request
	log.Printf("Attempting to start long poll with timeout %d and offset %d...", timeout, offset)

	// TODO: get latest UpdateId from database, to use (+1) as offset
 
	var telegramPollUrl string = "https://api.telegram.org/bot" + telegram_bot_token + "/getUpdates?timeout=" + strconv.Itoa(timeout) + "&offset=435559519"

	log.Printf("Telegram poll url: %s", telegramPollUrl)

	response, poll_err := http.Get(
		telegramPollUrl,
	)

	if poll_err != nil {
		log.Fatalf("Error polling the telegram server, %s", poll_err.Error())
	}

	return response

}


func main() {
	// TODO: read TELEGRAM_BOT_TOKEN from .env file
	var telegram_bot_token string = os.Getenv("TELEGRAM_BOT_TOKEN")
	if telegram_bot_token == "" {
		log.Fatal("Undefined TELEGRAM_BOT_TOKEN environment variable.")
	}

	db_err := createUpdatesTable()
	if db_err != nil {
		log.Fatalf("Could not start sqlite3 database: %s", db_err.Error())
	}

	// TODO: get latest UpdateId from database

	var response = longPollingHandler(60, 435559520, telegram_bot_token)

	// parse incoming request
	var receivedPayload, err = parseTelegramResponse(response)

	if err != nil {
		log.Printf("error parsing the telegram response, %s", err.Error())
		return
	}

	printParsedResponse(receivedPayload)

	// // TODO: add response logic
	var responseText string

	// // for each received message, save the UpdateId to the database and send a reply
	for i, s := range receivedPayload.Result {
		responseText = s.Message.Text + " ~beep"
		var telegramResponseBody, errTelegram = sendTextToTelegramChat(s.Message.Chat.Id, responseText)
		if errTelegram != nil {
			log.Printf("Got error %s from telegram on message %d, response body is %s", errTelegram.Error(), i, telegramResponseBody)
		} else {
			log.Printf("message %s successfuly distributed to chat id %d", responseText, s.Message.Chat.Id)
		}
	}

}