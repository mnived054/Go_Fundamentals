package main

import (
	"errors"
	"fmt"
)

// There are some basic operations that allow us to perform different actions on a stack.
type Stack struct {
	StackData []int
	Length    int
}

// 1)IsEmpty:= Check if the stack is empty.
func (st *Stack) IsEmpty() bool {
	return len(st.StackData) == 0
}

// 2)IsFull:= Check if the stack is full.
func (st *Stack) IsFull() bool {
	return len(st.StackData) >= st.Length
}

// 3)Push:= Add an element to the top of a stack.
func (st *Stack) Push(data int) error {
	if st.IsFull() {
		return errors.New("Stack is full")
	}

	st.StackData = append(st.StackData, data)
	return nil
}

// 4)Pop:=  Remove an element from the top of a stack.
func (st *Stack) Pop() (int, error) {
	if st.IsEmpty() {
		return -1, errors.New("Stack is empty")
	}

	data := st.StackData[len(st.StackData)-1]
	st.StackData = st.StackData[:len(st.StackData)-1]
	return data, nil
}

// 5)Peek:= Get the value of the top element without removing it.
func (st *Stack) Peek() (int, error) {
	if st.IsEmpty() {
		return -1, errors.New("Stack is empty")
	}
	return st.StackData[len(st.StackData)-1], nil
}
func DisplayData(st Stack) {
	fmt.Println("Stack Data:", st.StackData)
	fmt.Printf("Lenght of Stack Data: %d \n\n", st.Length)
}

func main() {
	stack := Stack{[]int{}, 4}

	_, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	}
	DisplayData(stack)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	DisplayData(stack)

	err = stack.Push(5)
	if err != nil {
		fmt.Println(err)
	}
	DisplayData(stack)

	data, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Poped data:", data)
	DisplayData(stack)

	data, err = stack.Peek()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Last Element:", data)
	DisplayData(stack)
}
