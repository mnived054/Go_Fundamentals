package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	fmt.Println(currentTime)
	fmt.Println(currentTime.Unix())
	fmt.Println(currentTime.UnixMilli())
	fmt.Println(currentTime.UnixNano())

	fmt.Println(time.Unix(currentTime.Unix(), 0))
	fmt.Println(time.Unix(0, currentTime.UnixNano()))
}
