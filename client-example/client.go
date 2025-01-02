package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	// Dial the WebSocket endpoint
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// Send a message to the server
	err = c.WriteMessage(websocket.TextMessage, []byte("Hello from client!"))
	if err != nil {
		log.Println("write:", err)
		return
	}

	// Read message from server
	_, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}

	fmt.Println("Received:", string(message))
}
