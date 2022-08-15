package telegrambotting

import (
	"envholder"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var envVars envholder.EnvHolder = envholder.LoadEnv()
var telegram_bot_token string = envVars.GetVar("TELEGRAM_BOT_TOKEN")

// sendTextToTelegramChat sends a text message to the Telegram chat identified by its chat Id
func SendTextToTelegramChat(chatId int, text string) (string, error) {
	
	log.Printf("Sending %s to chat_id: %d", text, chatId)
	var telegramApi string = "https://api.telegram.org/bot" + telegram_bot_token + "/sendMessage"
	
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		},
	)

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}
