package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "2000"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Println("Server starting...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error Listening:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Listening on" + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")

	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Getting Error; ", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go getClient(conn)
	}
}

func getClient(conn net.Conn) {
	buffer := make([]byte, 1024)
	msgLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received:", string(buffer[:msgLen]))
	_, err = conn.Write([]byte("Thanks! Got your message:" + string(buffer[:msgLen])))
	conn.Close()
}
