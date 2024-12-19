package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	ch1 := make(chan string)

	go func() {
		element := "nived"
		ch <- element
	}()
	go func() {
		ch1 <- <-ch
	}()

	val := <-ch1

	fmt.Println(val)

}
