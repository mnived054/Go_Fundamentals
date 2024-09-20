package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func (p person) details() string {
	return "hello my name is " + p.Name
}

func (p *person) HaveBirthday() {
	p.Age++
}
func main() {
	p := person{Name: "Nived", Age: 24}
	fmt.Println(p.Name)
	fmt.Println(p.Age)

	fmt.Println(p.details())
}
