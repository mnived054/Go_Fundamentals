package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const filename = "random.txt"

func main() {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Content of file" + filename + "\n\n")
	fmt.Println(string(data))
}
