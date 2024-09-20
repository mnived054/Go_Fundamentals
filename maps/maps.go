package main

import (
	"fmt"
)

func main() {
	exmp1()
}
func exmp1() {
	myMap := map[string]int{
		"a": 10,
		"b": 20,
		"c": 30,
	}

	fmt.Println(myMap)
	myMap["d"] = 40
	myMap["a"] = 50

	fmt.Println(myMap)

	delete(myMap, "d")
	fmt.Println(myMap)

}
