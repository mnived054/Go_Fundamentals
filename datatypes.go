package main

import "fmt"

/*func main() {
	print()
	variables()
	numeric()
	text()
	pointers()
	boolean()
	ifstatement()
	ifelsestatement()
	switchstatement()
	forloop()
	forlooprange()
	ds()
}*/

//different types of prints in console
//example for println
func print() {
	var a int
	a = 10
	fmt.Println(a)

	//example for printf (formatting output)
	var animal string = " Elephant"
	fmt.Printf("animal:%v\n", animal)
}

//types of variable declaration in go
func variables() {
	//type 1 variable declaration for string
	var yourName string = "Nived"
	fmt.Printf("Hi1 this is %v\n", yourName)

	//type 1 variable declaration for int
	var yourAge int = 23
	fmt.Printf("Iam %v years old\n", yourAge)

	//type 2 variable declaration for string
	yournewName := "Nived1"
	fmt.Printf("hi this is %v\n", yournewName)

	//type 2 variable declaration for int
	yournewAge := 24
	fmt.Printf("Iam %v year old now\n", yournewAge)

	//type 3 variable declaration for string
	var youroldName string
	youroldName = "Nived2"
	fmt.Printf("this is old name %v\n", youroldName)

	//type 3 variable declaration for int
	var youroldAge int
	youroldAge = 25
	fmt.Printf("I will be %v year old by next july\n", youroldAge)

}

// numeric data type
func numeric() {
	//  Integer types
	fmt.Print("These are numeric datatype\n")
	var age int = 28
	var count int64 = 1000000000
	fmt.Println("Age:\n", age)
	fmt.Println("Count:\n", count)
}

//text data type
func text() {
	fmt.Print("These are text datatype\n")
	str := "Welcome to Go programming\n"
	fmt.Printf("Lenght is: %d\n", len(str))
	fmt.Printf("\nString is: %s\n", str)
	fmt.Printf("\nType is: %T\n", str)
}

//boolean data type
func boolean() {
	fmt.Print("These are boolean datatype\n")
	var boolValue bool
	boolValue = false
	fmt.Println(boolValue)
}

//pointers using &symbol
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

//if statement
func ifstatement() {
	num := 48
	if num < 50 {
		fmt.Printf("The entered number is less than 50\n")
	}
}

//ifelse statement
func ifelsestatement() {
	temperature := 28
	if temperature > 30 {
		fmt.Println("It's a hot day.")
	} else {
		fmt.Println("It's not a hot day")
	}
}

//switch statement
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

//for loop
func forloop() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Welcome to the Program\n")
	}
}

//for loop with range
func forlooprange() {
	rvariable := []string{"ABC", "Hello", "Learners"}
	for i, j := range rvariable {
		fmt.Println(i, j)
	}
}

//data structures in Golang
//array declaration of int
func ds() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	//array declaration of string
	var arr1 [3]string
	arr1[0] = "Hi"
	arr1[1] = "hello"
	arr1[2] = "world"

	fmt.Println(arr1)
}
