package main

import (
	"fmt"
)

func PrintData(values ...any) {
	fmt.Println("Printing Values")
	for _, value := range values {
		fmt.Println(value)
	}
	fmt.Println()
}

func main() {
	PrintData("Hello", "World")
	PrintData(1, 2, 3)
	PrintData(true, false)
	PrintData(1.2, 3.4)
}
