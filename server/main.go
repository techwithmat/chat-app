package main

import (
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/ws", websocket.Handler(NewConnectionHandler))
	http.ListenAndServe(":3000", nil)
}
