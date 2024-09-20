package main

import "fmt"

func main() {
	dayofWeek := 2
	switch dayofWeek {
	case 1:
		fmt.Println("Monday")
		fallthrough
	case 2:
		fmt.Println("Tuesday")
		fallthrough
	case 3:
		fmt.Println("Wednesday")
		fallthrough
	case 4:
		fmt.Println("Thursday")
		fallthrough
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Invalid Day")
	}
}
