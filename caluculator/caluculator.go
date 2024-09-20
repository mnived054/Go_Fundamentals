package main

import "fmt"

func caluculate(num1 float64, operator string, num2 float64) float64 {
	result := 0.0

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Cannot divide by zero")
		}
	default:
		fmt.Println("Invalid operation")
	}
	return result
}

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter the first number:")
	fmt.Scanln(&num1)

	fmt.Print("Choose the operator(+,-,*,/):")
	fmt.Scanln(&operator)

	fmt.Print("Enter the second number:")
	fmt.Scanln(&num2)

	result := caluculate(num1, operator, num2)
	fmt.Printf("Result : %v\n", result)
}
