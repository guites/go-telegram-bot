package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
)

var envVars EnvHolder = loadEnv()
var telegram_bot_token string = envVars.getVar("TELEGRAM_BOT_TOKEN")

// longPollingHandler starts the request to the telegram API returns the reponse
func longPollingHandler(timeout int, offset int, telegram_bot_token string) (response *http.Response) {

	// starts the polling request
	log.Printf("Attempting to start long poll with timeout %d and offset %d...", timeout, offset)
 
	var telegramPollUrl string = "https://api.telegram.org/bot" + telegram_bot_token + "/getUpdates?timeout=" + strconv.Itoa(timeout) + "&offset=" + strconv.Itoa(offset)

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

	db_err := createUpdatesTable()
	if db_err != nil {
		log.Fatalf("Could not start sqlite3 database: %s", db_err.Error())
	}

	// open database connection
	db, db_err := sql.Open("sqlite3", "./updates.db")
	if db_err != nil {
		log.Fatal(db_err)
	}
	defer db.Close()

	// get latest UpdateId from database
	latestUpdate := getLatestUpdateFromDatabase(db)
	lastestOffset := latestUpdate.UpdateId + 1
	var response = longPollingHandler(60, lastestOffset, telegram_bot_token)

	// parse incoming request
	var receivedPayload, err = parseTelegramResponse(response)

	if err != nil {
		log.Printf("error parsing the telegram response, %s", err.Error())
		return
	}

	printParsedResponse(receivedPayload)

	// // TODO: add response logic
	// var responseText string

	// // for each received message, save the UpdateId to the database and send a reply
	for i, s := range receivedPayload.Result {
		newUpdate := DatabaseUpdate{
			UpdateId: s.UpdateId,
			Text: s.Message.Text,
			ChatId: s.Message.Chat.Id,
			FromId: s.Message.From.Id,
			First_Name: s.Message.From.First_Name,
			Last_Name: s.Message.From.Last_Name,
		}
		addUpdateToDatabase(db, newUpdate)
		log.Printf("message %d saved to db", i)
		// responseText = s.Message.Text + " ~beep"
		// var telegramResponseBody, errTelegram = sendTextToTelegramChat(s.Message.Chat.Id, responseText)
		// if errTelegram != nil {
		// 	log.Printf("Got error %s from telegram on message %d, response body is %s", errTelegram.Error(), i, telegramResponseBody)
		// } else {
		// 	log.Printf("message %s successfuly distributed to chat id %d", responseText, s.Message.Chat.Id)
		// }
	}

}