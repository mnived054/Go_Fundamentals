package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

var filename = "data.csv"

func ReadCsvFile(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("Unable to read input file %s \nError: %s", filepath, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Printf("Unable to parse file as CSV for %s\nError: %s", filepath, err)
	}
	return records
}

func main() {
	records := ReadCsvFile(filename)
	fmt.Println(records)
}
