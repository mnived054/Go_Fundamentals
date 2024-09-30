//By assigning any type in its generic form, we can make
//use of the same code for different functions, we have assigned
//"T" as an any keyword, this any  type allows us to parse different.
//types of variables within the same function.

package main

import "fmt"

type GenericSlice[T any] []T

func Print[T any](g GenericSlice[T]) {
	for _, V := range g {
		fmt.Printf("Value : %v & Type: %T\n", V, V)
	}
	fmt.Println()
}

func main() {
	a := GenericSlice[int]{1, 2, 3}
	b := GenericSlice[string]{"a", "b", "c"}
	c := GenericSlice[bool]{true, false}
	d := GenericSlice[float32]{1.5, 27.9, 32.4}
	e := GenericSlice[interface{}]{1, "a", true, 13.5}
	f := GenericSlice[struct{ name string }]{{"a"}, {"b"}, {"c"}}
	g := GenericSlice[complex64]{1 + 2i, 3 + 4i, 5 + 6i}

	Print(a)
	Print(b)
	Print(c)
	Print(d)
	Print(e)
	Print(f)
	Print(g)

}
