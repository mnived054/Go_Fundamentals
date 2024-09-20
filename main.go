package main

import (
	"datatypes/math"
	"fmt"
)

// goruotine
/*func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	go say("Hello")
	say("siri")
}*/
// func add() {
// 	var a int
// 	var b int
// 	fmt.Println("Enter first number")
// 	fmt.Scan(&a)
// 	fmt.Println("Enter second number")
// 	fmt.Scan(&b)

// 	fmt.Println("a+b=", a+b)
// }
/*func main() {

	//add()
	trail()
}*/

/*func trail() int {

	var d int
	var e int
	var f int
	var sum int

	fmt.Println("Enter first number")
	fmt.Scan(&d)
	fmt.Println("Enter second number")
	fmt.Scan(&e)
	sum = d + e
	fmt.Println("Enter third number")
	fmt.Scan(&f)

	fmt.Println("d+e*f=", sum*f)
	return sum * f

}*/

// func bubblesort(arr []int) {
// 	n := len(arr)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-i-1; j++ {

// 			if arr[j] > arr[j+1] {

// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}
// }

// func main() {
// 	var array = []int{2, 5, 1, 3, 52, 24}
// 	bubblesort(array)
// 	fmt.Println("sorted:", array)

// 	var array1 = []int{10, 30, 20, 50, 40}
// 	sort.Ints(array1)
// 	fmt.Println("sorted:", array1)

// }

//sorting array
/*func bubblesort(arr []int) {
	firstMax := math.MinInt16
	secondMax := math.MinInt16
	thirdmax := math.MinInt16
	for i := 0; i < len(arr); i++ {
		if arr[i] > firstMax {
			thirdmax = secondMax

			secondMax = firstMax
			firstMax = arr[i]

		} else if arr[i] > secondMax {
			thirdmax = secondMax
			secondMax = arr[i]

		} else if arr[i] > thirdmax {
			thirdmax = arr[i]
		}

	}

	fmt.Println(firstMax, secondMax, thirdmax)
}

func main() {
	//var array =[]int{20,40,30,1,5}
	bubblesort([]int{20, 40, 30, 1, 5})
}*/

//giving array size and target as input expecting output of indexnumber of array
/*func findindex(array []int, target int) (int, int, bool) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[i]+array[j] == target {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	target := 6

	index1, index2, found := findindex(array, target)

	if found {
		fmt.Printf("index %v and %v", array[index1], array[index2])
	} else {
		fmt.Printf("not found as per given target value")
	}

}*/
func main() {
	answer := math.Addition(100, 100)
	fmt.Println("Answer:", answer)
}
