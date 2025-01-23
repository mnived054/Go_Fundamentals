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

	specificDate := time.Date(2025, time.January, 23, 0, 0, 0, 0, time.UTC)
	fmt.Println(specificDate)
	fmt.Println(specificDate.Unix())
	fmt.Println(specificDate.UnixMilli())
	fmt.Println(specificDate.UnixNano())

	fmt.Println(time.Unix(specificDate.Unix(), 0))
	fmt.Println(time.Unix(0, specificDate.UnixNano()))
}
