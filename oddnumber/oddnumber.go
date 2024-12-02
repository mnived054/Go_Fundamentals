package main

import (
	"fmt"
)

func OddNumber(n int) bool {
	return n%2 != 0
}

func main() {
	for i := 1; i <= 100; i++ {
		if OddNumber(i) {
			fmt.Printf("%d is odd number\n", i)
		}
	}
}
