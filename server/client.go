package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/net/websocket"
)

var clients []*Client

type Client struct {
	id       string
	ip       string
	username string
	ws       *websocket.Conn
}

// StartListening method listens for messages from client.
func (client *Client) StartListening() {
	log.Println(client)
	buffer := make([]byte, 1024)

	for {
		n, err := client.ws.Read(buffer)

		// This error happens when the user disconnect
		if err != nil {
			ReleaseConnection(client)

			// Create Message
			exitMessage := Message{
				SenderID: "System",
				Username: "System",
				Message:  fmt.Sprintf("%s has left the chat", client.username),
			}

			// Send Message
			exitMessage.Post()
			break
		} else {
			HandleInputMessage(client, buffer[:n])
		}
	}
}

// ReleaseConnection function remove client from clients slice and ensure connection is closed.
func ReleaseConnection(client *Client) {
	log.Println("Release Connection:", client.username, client.id)

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

// GenerateUserId function generates a random 10 character id
func GenerateUserId() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOP0123456789")
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
