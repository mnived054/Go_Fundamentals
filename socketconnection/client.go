// package socketconnection

// import (
// 	"fmt"
// 	"net"
// )

// const (
// 	SERVER_HOST = "localhost"
// 	SERVER_PORT = "2000"
// 	SERVER_TYPE = "tcp"
// )

// func main() {
// 	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()

// 	_, err = conn.Write([]byte("Hello Server! Greetings."))
// 	buffer := make([]byte, 1024)
// 	msgLen, err := conn.Read(buffer)

// 	if err != nil {
// 		fmt.Println("Error reading:", err.Error())
// 	}

// 	fmt.Println("Received:", string(buffer[:msgLen]))

// }
