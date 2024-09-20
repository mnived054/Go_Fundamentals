package main

import "fmt"

func main() {
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
