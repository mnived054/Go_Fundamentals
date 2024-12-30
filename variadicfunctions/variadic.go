package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println("Sum of 1,2,3:", sum(1, 2, 3))
	fmt.Println("Sum of 5,10:", sum(5, 10))
	fmt.Println("sum with no arguments:", sum())
}
