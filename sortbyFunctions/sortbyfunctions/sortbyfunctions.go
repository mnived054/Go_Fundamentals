package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UsersByAge []User

func (u UsersByAge) Len() int {
	return len(u)
}

func (u UsersByAge) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u UsersByAge) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func main() {
	users := []User{
		{
			Name: "Jason",
			Age:  54,
		},
		{
			Name: "Monica",
			Age:  11,
		},
		{
			Name: "John",
			Age:  56,
		},
		{
			Name: "Jake",
			Age:  45,
		},
		{
			Name: "Robin",
			Age:  30,
		},
	}

	sort.Sort(UsersByAge(users))
	fmt.Println("Sorted users by age: ", users)
}
