package main

import "fmt"

func main() {
	for a := 1; a <= 6; a++ {
		if a == 7 {
			continue
		}
		fmt.Println(a)
	}
}
