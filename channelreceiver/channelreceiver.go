package main

import "fmt"

func SendValues(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go SendValues(ch)

	for value := range ch {
		fmt.Println("Recieved:", value)
	}
}
