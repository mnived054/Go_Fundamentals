package main

import "fmt"

func main() {

	//Note:- slices are similar has arrays but array contain with fixed sixe and slices doesnt declare particular index size
	arr1 := []int{1, 1, 1}
	fmt.Println(arr1)
	fmt.Println(len(arr1))
	fmt.Printf("capacity %v", cap(arr1))

	//slicing

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[5:]
	d := a[:4]
	e := a[2:4]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
