package main

import "fmt"

//struct 1
type vehicle struct {
	Name string
	Type string
}

//struct 2
type car struct {
	vehicle
	Maxspeed float32
	fueltype string
}

func main() {
	c := car{}
	c.Name = "Ferrari"
	c.Type = "Convertable"
	c.Maxspeed = 250
	c.fueltype = "Diesel"
	fmt.Println(c)
}
