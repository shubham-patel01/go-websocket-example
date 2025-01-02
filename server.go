package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// UpgradeRequest upgrades the initial GET request from the browser to a WebSocket connection.
func upgradeRequest(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader := websocket.Upgrader{} // Use a default upgrader
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return nil, err
	}
	return conn, nil
}

// HandleWebsocket handles the WebSocket connection.
func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgradeRequest(w, r)
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		// Read message from client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read:", err)
			break
		}

		fmt.Println("Received:", string(message))

		// Send message back to client
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Write:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebsocket)
	fmt.Println("Server starting on port 8080")
	log.Printf("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
