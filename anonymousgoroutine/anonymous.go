package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main functions : Start")
	go func() {
		fmt.Println("Anonymous Goroutine : Running")
		time.Sleep(2 * time.Second)
		fmt.Println("Anonymous Goroutine: Done")
	}()
	fmt.Println("Main function: Continuing")
	time.Sleep(3 * time.Second)
	fmt.Println("Main function: End")
}
