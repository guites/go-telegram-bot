package main

// ReceivedPayload is the wrapper around received telegram messages.
type ReceivedPayload struct {
	Ok bool `json:"ok"`
	Result []Update
}
