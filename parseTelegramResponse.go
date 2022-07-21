package main

import (
	"encoding/json"
	"io/ioutil"
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