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

	newmap1 := make([]map[string]int, 4)

	for i := range newmap1 {
		newmap1[i] = make(map[string]int)
	}

	newmap1[0]["Key1"] = 10
	newmap1[1]["Key2"] = 20
	newmap1[2]["Key3"] = 30
	newmap1[3]["Key4"] = 40

	fmt.Println(newmap1)
	fmt.Println(newmap1[0])
	fmt.Println(newmap1[1])
}
