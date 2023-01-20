package main

import (
	"encoding/json"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

// Post method broadcast message to all clients.
func (m *Message) Post() {
	m.Broadcast()
}

// Broadcast method sends message to all clients in client slice.
func (m *Message) Broadcast() {
	for _, client := range clients {
		m.BroadcastTo(client)
	}
}

// BroadcastTo method sends message to passed client.
func (m *Message) BroadcastTo(to *Client) {
	byteMessage, _ := json.Marshal(m)

	to.ws.Write(byteMessage)
}