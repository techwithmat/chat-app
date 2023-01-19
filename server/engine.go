package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	log.Printf("New %v from: %v -> %v", input.Action, from.username, input.Message)

	// Based on the action, handle it.
	switch input.Action {
	case "post_message":
		newMessage := Message{
			Username: from.username,
			Message:  input.Message,
		}

		newMessage.Post()

	case "initial_connection":
		// Set client username.
		from.username = input.Username

		// Send chat history to the new user.
		chatHistory, _ := json.Marshal(messages)
		from.ws.Write(chatHistory)

		// Send to everyone that a user has joined to the chat.
		newMessageToRoom := Message{
			Username: "System",
			Message:  fmt.Sprintf("%s joined the chat", from.username),
		}

		newMessageToRoom.Post()
	}
}
