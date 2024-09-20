package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		ch1 <- 42
	}()

	go func() {
		data := <-ch2
		fmt.Println("Received from ch2:", data)
	}()

	select {
	case value := <-ch1:
		fmt.Println("Received from ch1:", value)
	case ch2 <- "Hello, Golang!":
		fmt.Println("Sent to ch2")
	}
}
