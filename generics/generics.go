package main

import (
	"fmt"
)

func PrintElement[T any](list []T) {
	fmt.Printf("%T type data\n", list)
	for _, data := range list {
		fmt.Println(data)
	}
}

func main() {
	var list_name = []string{"ABC", "CDE", "FGH"}
	var list_age = []int{10, 20, 30}
	PrintElement(list_name)
	PrintElement(list_age)
}
