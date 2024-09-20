package main

import (
	"fmt"
)

func main() {
	ds()
}

// array declaration for int
func ds() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	marks := [3]int{90, 80, 70}
	fmt.Printf("%v \n", marks)
	//single dimensional of array declaration of string
	var arr1 [3]string
	arr1[0] = "Hi"
	arr1[1] = "hello"
	arr1[2] = "world"

	fmt.Println(arr1)

	//single dimensional of array declaration of string
	var arr2 = [3]string{"Apple", "Ball", "Cat"}

	arr3 := arr2
	arr3[0] = "AMerica"
	arr3[1] = "Newyork"
	fmt.Println(arr2)
	fmt.Println(arr3)

	//multidimensional array
	var matrix2d [2][2]int
	matrix2d[0] = [2]int{1, 0}
	matrix2d[1] = [2]int{0, 1}

	fmt.Println(matrix2d)

}
