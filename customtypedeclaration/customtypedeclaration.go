package main

import "fmt"

type CustomString string
type CustomInt int
type CustomBoolean bool

func main() {
	var name CustomString = "Jake"
	fmt.Printf("Value of variable name %v and it's type %T \n", name, name)

	var age CustomInt = 25
	fmt.Printf("Value of variable name %v and it's type %T \n", age, age)

	var valid CustomBoolean = true
	fmt.Printf("Value of variable name %v and it's type %T \n", valid, valid)

}
