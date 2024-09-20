package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var SECRET_KEY = "THIS IS SECRET KEY"

type User struct {
	Id    uuid.UUID
	Email string
}

func GenerateJWT(user *User) (string, error) {
	var mySigningKey = []byte(SECRET_KEY)
	var token = jwt.New(jwt.SigningMethodHS256)
	var claim = token.Claims.(jwt.MapClaims)

	claim["authorized"] = true
	claim["email"] = user.Email
	claim["id"] = user.Id
	claim["exp"] = time.Now().Add(time.Minute * 20).Unix()

	token_string, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Printf("something went wrong: %s", err.Error())
		return "", err
	}
	return token_string, nil
}

func main() {
	user_id := uuid.New()
	email := "user1@gmail.com"

	fmt.Println("User info:")
	fmt.Printf("Id: %s\nEmail: %s\n", user_id.String(), email)

	user := &User{
		Id:    user_id,
		Email: email,
	}
	token, err := GenerateJWT(user)
	if err != nil {
		return
	}
	fmt.Println("Generated Token:", token)
}
