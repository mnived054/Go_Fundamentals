package main

import "fmt"

func main() {
	ifelsestatement()
}

func ifelsestatement() {
	temperature := 28
	if temperature > 30 {
		fmt.Println("It's a hot day.")
	} else {
		fmt.Println("It's not a hot day")
	}
}
