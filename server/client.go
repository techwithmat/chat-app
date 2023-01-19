package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

var clients []*Client

type Client struct {
	username string
	ws       *websocket.Conn
}

// StartListening method listens for messages from client.
func (client *Client) StartListening() {
	log.Println("New User starts listening")
	
	buffer := make([]byte, 1024)

	for {
		n, err := client.ws.Read(buffer)

		// This error happens when the user disconnect
		if err != nil {
			ReleaseConnection(client)

			// Create Message
			exitMessage := Message{
				Username: "System",
				Message:  fmt.Sprintf("%s has left the chat", client.username),
			}

			// Send Message
			exitMessage.Post()
			break
		}

		HandleInputMessage(client, buffer[:n])
	}
}

// ReleaseConnection function remove client from clients slice and ensure connection is closed.
func ReleaseConnection(client *Client) {
	log.Println("Release Connection:", client.username)

	// Remove from clients slice.
	index := -1
	for idx, value := range clients {
		if client == value {
			index = idx
			break
		}
	}

	if index >= 0 {
		clients = append(clients[:index], clients[index+1:]...)
	}

	log.Println("Connection active:", len(clients))
	client.ws.Close()
}
