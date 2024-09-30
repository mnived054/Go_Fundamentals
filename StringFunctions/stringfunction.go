package main

import (
	"fmt"
	String "strings"
)

func main() {
	fmt.Println("ToLower:   ", String.ToLower("JAKE"))
	fmt.Println("ToUpper:	", String.ToUpper("jake"))
	fmt.Println("Contains:	", String.Contains("jeremy", "re"))
	fmt.Println("Index:	    ", String.Index("jeremy", "e"))
	fmt.Println("Count:		", String.Count("jeremy", "e"))
	fmt.Println("Replace:	", String.Replace("Absolute", "o", "0", -1))
	fmt.Println("Replace:	", String.Replace("Absolute", "b", "0", 1))
	fmt.Println("HasPrefix:	", String.HasPrefix("jake", "ja"))
	fmt.Println("HasSuffix:	", String.HasSuffix("jake", "ke"))
	fmt.Println("Join:		", String.Join([]string{"1", "2", "3"}, "-"))
	fmt.Println("Repeat:	", String.Repeat("R", 5))
	fmt.Println("Split:	    ", String.Split("a-b-c-d-e", "-"))
}
