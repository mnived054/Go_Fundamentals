package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const file_name = "random1.txt"

func main() {
	information := "Hi hello this is new updated information 28-08-2024"
	data := []byte(information)
	err := ioutil.WriteFile(file_name, data, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Written completion is done")
}
