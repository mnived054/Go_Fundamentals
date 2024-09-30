package main

import (
	"fmt"
)

type User string
type Person string

func (u User) chechClass(class string) string {
	return fmt.Sprint(u, "belongs to class:", class)
}

func check(v interface{}) bool {
	_, has := v.(interface{ chechClass(string) string })
	return has
}

func main() {
	var user User = "john"
	var person Person = "jake"
	fmt.Println(check(user))
	fmt.Println(check(person))
}
