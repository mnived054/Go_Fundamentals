package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := "This Is Go Code"
	fmt.Println("Data in string format:", data)

	hash := sha256.New()
	hash.Write([]byte(data))
	fmt.Printf("Hashed Data: %x\n", hash.Sum(nil))
}
