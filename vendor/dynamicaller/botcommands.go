package dynamicaller

import (
	"log"
	"telegrambotting"
)

func (dynamicaller Dynamicaller) EchoHandlerFunc(arg0 string, chatId int) (res0 string, res1 string, err error) {
	log.Println("Calling EchoHandlerFunc! with args", arg0)
	if arg0 == "" {
		telegrambotting.SendTextToTelegramChat(chatId, "Usage: /echo Message to be echoed back")
	} else {
		telegrambotting.SendTextToTelegramChat(chatId, "Echoing \""+arg0+"\"")
	}
	return "ok", "another ok", nil
}