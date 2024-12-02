package main

import (
	"fmt"
)

func Prime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	for i := 1; i <= 100; i++ {
		if Prime(i) {
			fmt.Printf("%d is a prime number\n", i)
		}
	}
}
