package main

import (
	"log"

	"golang.org/x/net/websocket"
)

// NewConnectionHandler creates a new client and stores it.
func NewConnectionHandler(connection *websocket.Conn) {
	log.Println("New Connection!!!")

	// Create a new Client object and add it to the clients array.
	newClient := Client{
		id:         GenerateUserId(),
		ip:         connection.Request().RemoteAddr,
		username:   "",
		ws: connection,
	}

	clients = append(clients, &newClient)

	newClient.StartListening()
}
