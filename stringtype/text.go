package main

import (
	"fmt"
	"strings"
)

func main() {
	text()
	splitstring()
}
func text() {
	fmt.Print("These are text datatype\n")
	str := "Welcome to Go programming\n"
	fmt.Printf("Lenght is: %d\n", len(str))
	fmt.Printf("\nString is: %s\n", str)
	fmt.Printf("\nType is: %T\n", str)
}

func splitstring() {
	str := "Nived"
	strings.Split(str, ",")
	fmt.Println(str)
}
