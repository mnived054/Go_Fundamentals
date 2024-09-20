package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

const file_name = "person.json"

func ConvertStructToInterface(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return map[string]interface{}{}, err
	}
	err = json.Unmarshal(data, &newMap)
	return newMap, nil
}

func main() {
	person_1 := Person{"John", 23, "USA"}

	data, err := ConvertStructToInterface(person_1)
	if err != nil {
		log.Println(err)
	}
	file, err := json.MarshalIndent(data, "", "")
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(file_name, file, 0644)
	fmt.Println("Write operation completed")
}
