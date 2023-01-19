package main

import (
	"net/http"

	"golang.org/x/net/websocket"
)

// NewConnectionHandler creates a new client and stores it.
func NewConnectionHandler(connection *websocket.Conn) {
	// Create a new Client object and add it to the clients array.
	newClient := Client{
		username: "", // username will be assigned after the connection is successfully
		ws:       connection,
	}

	clients = append(clients, &newClient)
	newClient.StartListening()
}

func main() {
	http.Handle("/ws", websocket.Handler(NewConnectionHandler))
	http.ListenAndServe(":3000", nil)
}
