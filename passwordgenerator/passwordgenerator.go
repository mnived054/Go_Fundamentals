package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const Cap = "ABCCDEFGHIJKLMNOPQRSTUVWXYZ"
const Small = "abcdefghijklmnopqrstuvwxyz"
const Number = "0123456789"
const SpecialChars = "!@#$%&_"
const LengthofPassword = 8

func convertToSlice(s string) []string {
	return strings.Split(s, "")
}

func mergeAllCharacter() []string {
	var allCharacter []string
	allCharacter = append(allCharacter, convertToSlice(Cap)...)
	allCharacter = append(allCharacter, convertToSlice(Small)...)
	allCharacter = append(allCharacter, convertToSlice(Number)...)
	allCharacter = append(allCharacter, convertToSlice(SpecialChars)...)
	return allCharacter
}

func main() {
	var password = ""
	slice_of_password_characters := mergeAllCharacter()

	for i := 0; i < LengthofPassword; i++ {
		random_data := slice_of_password_characters[rand.Intn(len(slice_of_password_characters))]
		password += random_data
	}
	fmt.Println("Password:", password)
}
