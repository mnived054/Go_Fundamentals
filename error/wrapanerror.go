package main

import (
	"errors"
	"fmt"
)

var errCritical = errors.New("oh critical error ->")

func check(number int) error {
	if number == 1 {
		return fmt.Errorf("it's odd")
	} else if number == 2 {
		return errCritical
	}
	return nil
}

func validation(number int) error {
	err := check(number)
	if err != nil {
		return fmt.Errorf("run error: %w", err)
	}
	return nil
}

func main() {
	for number := 1; number <= 5; number++ {
		fmt.Printf("validation %d... \n", number)
		err := validation(number)
		if err == errCritical {
			fmt.Println("oh no something has happened!")
		} else if err != nil {
			fmt.Println("some error is present...", err)
		} else {
			fmt.Println("valid number only...!")
		}
	}
}
