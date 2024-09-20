package main

import "fmt"

func main() {
	switchstatement()
}

func switchstatement() {
	day := "Thursday"

	switch day {
	case "Tuesday":
		fmt.Println("It's still early in the week.")
	case "Wednesday":
		fmt.Println("It's the middle of the week.")
	case "Thursday":
		fmt.Println("The weekend is approaching.")
	default:
		fmt.Println("Its the weekend!")

	}
}
