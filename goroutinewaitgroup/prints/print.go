package main

import "fmt"

func main() {
	variable()
}

func variable() {
	//example for println
	var a int
	a = 10
	fmt.Println(a)

	//example for printf (formatting output)
	var animal string = " Elephant"
	fmt.Printf("animal:%v\n", animal)
}
