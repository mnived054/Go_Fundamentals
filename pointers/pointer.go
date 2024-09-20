package main

import "fmt"

func main() {
	pointers()
}

func pointers() {
	var firstName string
	var lastName string
	var email string
	var tickets int

	fmt.Println("Enter your firstname please: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastName please: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email please: ")
	fmt.Scan(&email)

	fmt.Println("Enter how many tickets do you want to purchase:")
	fmt.Scan(&tickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, tickets, email)
}
