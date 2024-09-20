package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello!")
		time.Sleep(time.Millisecond * 500)
	}

}

func sayBye() {
	for i := 0; i < 5; i++ {
		fmt.Println("Byee")
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go sayHello()
	go sayBye()
	for i := 0; i < 5; i++ {
		fmt.Println("Hi!")
		time.Sleep(time.Millisecond * 500)
	}
	time.Sleep(time.Second)
}
