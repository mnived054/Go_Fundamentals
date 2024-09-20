package main

import (
	"errors"
	"fmt"
)

// The basic operations supported in the queue-
type Queue struct {
	Data []int
	Size int
}

// 1)Enqueue := Add an element to the end of the queue.
func (q *Queue) Enqueue(data int) {
	if len(q.Data) >= q.Size {
		fmt.Println("Queue is full")
		return
	}
	q.Data = append(q.Data, data)
}

// 2)Dequeue := Remove an element from the front of the queue.
func (q *Queue) Dequeue() {
	if len(q.Data) == 0 {
		fmt.Println("Queue is empty")
		return
	}
	q.Data = q.Data[1:]
}

// 3)IsFull:=Check if the queue is full.
func (q *Queue) IsFull() {
	if len(q.Data) >= q.Size {
		fmt.Printf("Queue is full \n\n")
		return
	}
	fmt.Printf("Queue is not full \n\n")
}

// 4)IsEmpty:= Check if the queue is empty.
func (q *Queue) IsEmpty() {
	if len(q.Data) == 0 {
		fmt.Printf("Queue is empty \n\n")
		return
	}
	fmt.Printf("Queue is not empty \n\n")
}

// 5)Peek:= Get the value if the front of the queue without removing it.
func (q *Queue) Peek() (int, error) {
	if len(q.Data) == 0 {
		return -1, errors.New("Queue is empty")
	}
	return q.Data[0], nil
}

func (q *Queue) DisplayData() {
	fmt.Println("Data of Queue:", q.Data)
	fmt.Println("Length of Queue:", len(q.Data))
	fmt.Println("Allowed Size of Queue:", q.Size)
	fmt.Println()
}

func main() {
	queue := Queue{[]int{}, 5}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.DisplayData()

	queue.Dequeue()
	queue.DisplayData()

	queue.IsFull()

	queue.IsEmpty()

	data, err := queue.Peek()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Value of the front of the queue:", data)

}
