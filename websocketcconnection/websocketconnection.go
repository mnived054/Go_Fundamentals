package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocketHandler(writer http.ResponseWriter, request *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client Connected.....")

	Reader(ws)
}

func Reader(conn *websocket.Conn) {
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Message from Client:", string(msg))

		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", WebSocketHandler)
	log.Println("WebSocket started on port 5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
