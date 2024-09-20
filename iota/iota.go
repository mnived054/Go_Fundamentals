package main

import "fmt"

const (
	constant_0 = iota
	constant_1 = iota
	constant_2 = iota
)

func main() {
	fmt.Printf("constant_0 has type: %T and it's value: %v \n", constant_0, constant_0)
	fmt.Printf("constant_0 has type: %T and it's value: %v \n", constant_1, constant_1)
	fmt.Printf("constant_0 has type: %T and it's value: %v \n", constant_2, constant_2)
}
