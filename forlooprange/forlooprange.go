package main

import "fmt"

func main() {
	forlooprange()

}
// forloop using array which indicates range
func forlooprange() {
	rvariable := []string{"ABC", "Hello", "Learners"}
	for i, j := range rvariable {
		fmt.Println(i, j)
	}
}
