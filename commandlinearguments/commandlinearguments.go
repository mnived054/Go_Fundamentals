package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	index := 1
	arg := os.Args[index]

	for index, arg := range argsWithProg {
		fmt.Println("Index of Arg:", index, "& value of Args:", arg)
	}

	fmt.Println("List of Arguments:", argsWithoutProg)
	fmt.Println("Argument with index", index, "is:", arg)
}
