package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassowrd(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	var password string
	fmt.Print("enter the password")
	fmt.Scan(&password)

	fmt.Print("Hashing is Done \n")

	hash, _ := HashPassowrd(password)

	fmt.Println("password:", password)

	fmt.Println("hash:", hash)

	match := CheckHashPassword(password, hash)

	fmt.Println("match:", match)
}
