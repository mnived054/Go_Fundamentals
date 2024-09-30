package main

import (
	"fmt"
	"math/rand"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("Function: %s took %v\n", name, time.Since(start))
	}
}

func Add() {
	defer timer("Add")()
	number, answer := 5, 0

	for i := 0; i < number; i++ {
		answer += rand.Intn(100)
	}
}

func Multiply() {
	defer timer("Multiply")()
	number, answer := 5, 0

	for i := 0; i < number; i++ {
		answer *= rand.Intn(100)
	}
}

func main() {
	Add()
	Multiply()
}
