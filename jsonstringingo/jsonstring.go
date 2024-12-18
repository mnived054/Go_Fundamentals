package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {

	person := Person{
		Name:    "Nived",
		Age:     24,
		Address: "hyd",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal("Error encoding JSON: ", err)
	}
	fmt.Println("Encoded JSON string:")
	fmt.Println(string(jsonData))

	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		log.Fatal("Error decoding JSON: ", err)
	}
	fmt.Println("\nDecoded struct: ")
	fmt.Printf("Name: %s, Age:%d, Address: %s\n", decodedPerson.Name, decodedPerson.Age, decodedPerson.Address)
}
