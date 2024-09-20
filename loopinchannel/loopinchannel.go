package main

import "fmt"

func main() {
	mychn1 := make(chan string)
	go func() {
		mychn1 <- "welcome"
		mychn1 <- "to"
		mychn1 <- "Course"

		close(mychn1)
	}()
	for res := range mychn1 {
		fmt.Println(res)
	}
}
