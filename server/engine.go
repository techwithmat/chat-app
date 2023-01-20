package main

import (
	"encoding/json"
	"fmt"
)

type InputMessage struct {
	Action   string `json:"action"`
	Username string `json:"username"`
	Message  string `json:"message,omitempty"`
}

// HandleInputMessage is the primary handler for the websocket connection which listen from messages the client.
func HandleInputMessage(from *Client, data []byte) {
	// Parse JSON
	var input InputMessage
	json.Unmarshal(data, &input)

	// Based on the action, handle it.
	switch input.Action {
	case "message":
		newMessage := Message{
			Username: from.username,
			Message:  input.Message,
		}

		newMessage.Post()
	case "connect":
		// Set client username.
		from.username = input.Username

		// Send to everyone that a user has joined to the chat.
		newMessageToRoom := Message{
			Username: "System",
			Message:  fmt.Sprintf("%s joined the chat", from.username),
		}

		newMessageToRoom.Post()
	}
}
