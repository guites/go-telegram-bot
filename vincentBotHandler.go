// telegramCallbackHandler sends a message back to the chat which contacted the bot
func telegramCallbackHandler(w http.ResponseWriter, r *http.Request) {

	// parse incoming request
	var update, err = parseTelegramRequest(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}

	// sanitize input
	var sanitizedSeed = sanitize(update.Message.Text)

	// TODO: add the handling logic
	var responseText string = "Test response! Beep."

	// send a message back to telegram
	var telegramResponseBody, errTelegram = sendTextToTelegramChat(update.Message.Chat.Id, responseText)
	if errTelegram != nil {
		log.Printf("got error %s from telegram, response body is %s", errTelegram.Error(), telegramResponseBody)
	} else {
		log.Printf("message %s successfuly distributed to chat id %d", responseText, update.Message.Chat.Id)
	}
}
