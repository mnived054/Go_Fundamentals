package main

import (
	"fmt"
)

func Travers(arr *[7]string) {
	fmt.Println("This is Traversal method:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr [ %d ] = %s", i, arr[i])
	}
	fmt.Printf("\n\n")
}

func Search(arr *[7]string, data string) {
	fmt.Println("Search Method:")
	for i := 0; i < len(arr); i++ {
		if arr[i] == data {
			fmt.Printf("Element %s found at %d position \n\n", data, i)
			return
		}
	}
	fmt.Printf("Element does not found. \n\n")
}

func Insert(arr [7]string, data string, position int) {
	fmt.Println("This is Insert method:")
	var newarr [9]string
	for i := 0; i < len(arr); i++ {
		if i < position {
			newarr[i] = arr[i]
		} else if i == position {
			newarr[i] = data
			newarr[i+1] = arr[i]
		} else {
			newarr[i+1] = arr[i]
		}
	}
	fmt.Println("After performing Insert method:", newarr)
	fmt.Printf("\n")
}

func delete(arr [7]string, position int) {
	fmt.Println("This is delete method:")
	if position < 0 || position > len(arr) {
		fmt.Printf("Possiton out of index. \n\n")
		return
	}

	var index = 0
	for i := range arr {
		if position != i {
			arr[index] = arr[i]
			index++
		}
	}

	newarr := arr[:index]
	fmt.Println("After Performing Delete Method:", newarr)
	fmt.Printf("\n")
}

func Update(arr [7]string, data string, position int) {
	fmt.Println("This is Update method:")
	for i := 0; i < len(arr); i++ {
		if i == position {
			arr[i] = data
		}
	}

	fmt.Println("After Performing Update Method:", arr)
	fmt.Printf("\n")
}

func main() {
	var arr = [7]string{"A", "B", "C", "D", "E", "F", "G"}

	Travers(&arr)
	Search(&arr, "B")
	Insert(arr, "H", 5)
	delete(arr, 3)
	Update(arr, "B", 2)
	Update(arr, "C", 3)
}
