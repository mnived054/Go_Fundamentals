package main

import (
	"fmt"
)

func sumAndDivide(x, y float64) float64 {
	var sum = x
	for i := x; i < y; i++ {
		sum = sum + 1 + i
	}

	return sum / (x + y)

}

func main() {

	fmt.Println(sumAndDivide(4, 10))
}
