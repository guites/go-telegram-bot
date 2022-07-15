package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// parseTelegramResponse handles incoming update from Telegram
func parseTelegramResponse(r *http.Response) (*ReceivedPayload, error) {

	// TODO: handle non http 200 cases

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err.Error())
	}

	var data ReceivedPayload
	json.Unmarshal(body, &data)

	return &data, nil
}

func printParsedResponse(receivedPayload *ReceivedPayload) {
	log.Printf("Received Ok: %t\n", receivedPayload.Ok)
	for i, s := range receivedPayload.Result {
		log.Printf(
			"Message #%d: \n"+
			"{\n"+
			"    'UpdateId': '%d',\n"+
			"    'Message.Text': '%s',\n"+
			"    'Message.From.Id': '%d',\n"+
			"    'Message.From.First_Name': '%s',\n"+
			"    'Message.From.Last_Name': '%s',\n"+
			"}\n",
			i,
			s.UpdateId,
			s.Message.Text,
			s.Message.From.Id,
			s.Message.From.First_Name,
			s.Message.From.Last_Name,
		)
	}
}