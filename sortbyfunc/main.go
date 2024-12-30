package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	words := []string{"bannana", "kiwi", "apple", "cherry"}

	lencmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(words, lencmp)

	fmt.Println(words)
}
