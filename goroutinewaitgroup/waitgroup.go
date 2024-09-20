package main

import (
	"fmt"
	"sync"
)

func printNumbers(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("Goroutine %d: %v\n", id, i)
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go printNumbers(&wg, i)
	}
	wg.Wait()
}
