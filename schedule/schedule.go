package main

import (
	"fmt"
	"runtime"
	"time"
)

const iterations = 5

func firstRoutine() {
	for i := 1; i < iterations; i++ {
		fmt.Println("1st Goroutine - Counter", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func secondRuntime() {
	for i := 1; i < iterations; i++ {
		fmt.Println("2nd Goroutine - Counter", i)
		time.Sleep(500 * time.Millisecond)
	}
}
func main() {
	go firstRoutine()
	go secondRuntime()

	runtime.Gosched()

	time.Sleep(2 * time.Second)
}
