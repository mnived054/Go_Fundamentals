package main

import (
	"fmt"
	"strconv"
)

// Define an interface
type BoxOffice interface {
	Industry() string
}

// Define a Movie type
type Movie struct {
	Mname string
	Mid   int
}

// Define a Director type
type Director struct {
	Dname string
}

// Implement the Speak method for the Movie type
func (m Movie) Industry() string {
	return "Movie Name :" + m.Mname + " MID :" + strconv.Itoa(m.Mid)
}

// Implement the Speak method for the Director type
func (d Director) Industry() string {
	return "Director Name: " + d.Dname
}

func main() {
	// Create instances of Movie and Director
	puspha1 := Movie{Mname: "Pushpa-1", Mid: 50}

	d := Director{Dname: "Sukumar"}

	m1 := Movie{Mname: "Pushpa-2"}
	mid1 := Movie{Mid: 60}
	d1 := Director{Dname: "Sukumar"}

	// Create a slice of BoxOffice type (which can hold any type that implements the Speak method)
	var BoxOffices []BoxOffice = []BoxOffice{puspha1, d, m1, mid1, d1}

	fmt.Println(BoxOffices)

}
