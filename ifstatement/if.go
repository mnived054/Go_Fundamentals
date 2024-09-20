package main

import "fmt"

func main() {
	ifstatement()
}

func ifstatement() {
	num := 48
	if num < 50 {
		fmt.Printf("The entered number is less than 50\n")
	}
}
