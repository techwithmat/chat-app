package main

import (
	"encoding/json"
	"fmt"
)

// HandleInputMessage is the primary handler for the websocket connection which listen from messages the client
func HandleInputMessage(from *Client, data []byte) {
	// Parse JSON
	var input map[string]string
	json.Unmarshal(data, &input)

	// Based on the action, handle it
	switch input["action"] {
	case "post_message":
		newMessage := Message{
			SenderID: from.id,
			Username: from.username,
			Message:  input["message"],
		}

		newMessage.Post()

	case "initial_connection":
		// Set client username.
		from.username = input["username"]

		// Create a message to notificate the user his ID.
		newIdMessage := Message{
			SenderID: "System",
			Username: "System",
			Message:  fmt.Sprintf("Your ID: %v", from.id),
		}
		newIdMessage.BroadcastTo(from)

		// Send chat history to the new user.
		chatHistory, _ := json.Marshal(messages)
		from.ws.Write(chatHistory)

		// Send to everyone that a user has joined to the chat.
		newMessageToRoom := Message{
			SenderID: "System",
			Username: "System",
			Message:  fmt.Sprintf("%s joined the chat", from.username),
		}

		newMessageToRoom.Post()
	}
}
