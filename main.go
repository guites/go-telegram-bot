package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
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

func testLOL() {
	db, db_err := sql.Open("sqlite3", "./updates.db")
	if db_err != nil {
		log.Fatal(db_err)
	}
	newUpdate := DatabaseUpdate{
		UpdateId: 123123,
		Text: "/lembrete where tf this going",
		ChatId: 213123,
		FromId: 44,
		First_Name: "guites",
		Last_Name: "hackerman millionaire",
		Type: "bot_command",
		Length: 9,
		Offset: 0,
	}
	handleRespondingLogic(db, newUpdate)
}

func runBot() {
	for {
		db, db_err := sql.Open("sqlite3", "./updates.db")
		if db_err != nil {
			log.Println(db_err)
			time.Sleep(3 * time.Second)
		}
		// get latest UpdateId from database
		latestUpdate, update_err := getLatestUpdateFromDatabase(db)
		
		var latestOffset int
		if update_err != nil {
			log.Println(update_err)
			latestOffset = 0
		} else {
			latestOffset = latestUpdate.UpdateId + 1
		}
		var response = longPollingHandler(60, latestOffset, telegram_bot_token)

		// parse incoming request
		var receivedPayload, err = parseTelegramResponse(response)

		if err != nil {
			log.Printf("error parsing the telegram response, %s", err.Error())
		}

		for i, s := range receivedPayload.Result {
			newUpdate := DatabaseUpdate{
				UpdateId: s.UpdateId,
				Text: s.Message.Text,
				ChatId: s.Message.Chat.Id,
				FromId: s.Message.From.Id,
				First_Name: s.Message.From.First_Name,
				Last_Name: s.Message.From.Last_Name,
				
			}
			if len(s.Message.Entity) > 0 {
				newUpdate.Type = s.Message.Entity[0].Type
				newUpdate.Length = s.Message.Entity[0].Length
				newUpdate.Offset = s.Message.Entity[0].Offset
			}
			addUpdateToDatabase(db, newUpdate)
			log.Printf("message #%d (%d) saved to db", i, s.UpdateId)
			handleRespondingLogic(db, newUpdate)

		}
		db.Close()
	}
}

func addComands() {
	var name string
	var callback string

	db, db_err := sql.Open("sqlite3", "./updates.db")
	if db_err != nil {
		log.Println(db_err)
		time.Sleep(3 * time.Second)
	}

	// TODO: validate for empty input, min length, command starts with slash etc
	fmt.Print("Insert command name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Command name is", name)

	fmt.Print("Insert command callback function: ")
	fmt.Scanf("%s", &callback)
	fmt.Println("Callback function name:", callback)

	newCommand := DatabaseCommand{
		Name: name,
		Callback: callback,
	}

	addCommandToDatabase(db, newCommand)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db_err := createTables()
	if db_err != nil {
		log.Fatalf("Could not start sqlite3 database: %s", db_err.Error())
	}

	cmdArgs := os.Args[1:]
	if len(cmdArgs) < 1 {
		log.Fatal("Available options: [bot|add_commands]")
	}
	chosenOption := cmdArgs[0]
	switch chosenOption {
	case "bot":
		runBot()
	case "add_commands":
		addComands()
	case "test":
		testLOL()
	default:
		fmt.Println("Available options: [bot|add_commands]")
	}
}

// responseText = s.Message.Text + " ~beep"
// var telegramResponseBody, errTelegram = sendTextToTelegramChat(s.Message.Chat.Id, responseText)
// if errTelegram != nil {
// 	log.Printf("Got error %s from telegram on message %d, response body is %s", errTelegram.Error(), i, telegramResponseBody)
// } else {
// 	log.Printf("message %s successfuly distributed to chat id %d", responseText, s.Message.Chat.Id)
// }