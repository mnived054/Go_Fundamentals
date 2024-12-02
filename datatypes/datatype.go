package main

import "fmt"

func main() {
	numeric()
}

func numeric() {
	//  Integer types
	fmt.Print("These are numeric datatype\n")
	var age int = 28
	var count int64 = 1000000000
	fmt.Println("Age:\n", age)
	fmt.Println("Count:\n", count)
}
