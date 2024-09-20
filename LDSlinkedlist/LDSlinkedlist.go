package main

import (
	"errors"
	"fmt"
)

// The basic operations that are supposed by a list are mentioned as follows-
type Node struct {
	Value int
	Next  *Node
}

type Linkedlist struct {
	Length int
	Head   *Node
}

// Insert := This operation is performed to add an element into the list.
func (Linked_list *Linkedlist) Insert(Value int) {
	var new_node = Node{}
	new_node.Value = Value

	if Linked_list.Length == 0 {
		Linked_list.Length++
		Linked_list.Head = &new_node
		return
	}

	node := Linked_list.Head
	for i := 0; i < Linked_list.Length; i++ {
		if node.Next == nil {
			node.Next = &new_node
			Linked_list.Length++
			break
		}
		node = node.Next
	}
}

// Search := It is performed to search an element from the list using the given key.
func (Linked_list *Linkedlist) Search(Value int) (int, error) {
	node := Linked_list.Head

	for i := 0; i < Linked_list.Length; i++ {
		if node.Value == Value {
			return i, nil
		}
		node = node.Next
	}
	return -1, errors.New("value not found")
}

// Display := It is performed to display the elements of the list.
func (Linked_list *Linkedlist) Display() {
	if Linked_list.Length == 0 {
		fmt.Println("No nodes in list")
		return
	}
	ptr := Linked_list.Head
	for i := 0; i < Linked_list.Length; i++ {
		fmt.Println("Node:", ptr)
		ptr = ptr.Next
	}
	fmt.Printf("Lenght of linked list : %d \n\n", Linked_list.Length)
}

func (Linked_list *Linkedlist) GetAt(pos int) *Node {
	node := Linked_list.Head

	if pos < 0 {
		return node
	}
	if pos >= Linked_list.Length {
		return nil
	}
	for i := 0; i < pos; i++ {
		node = node.Next
	}

	return node
}

// Delete := It is performed to delete an operation from the list.
func (Linked_list *Linkedlist) Delete(Value int) error {
	node := Linked_list.Head

	if Linked_list.Length == 0 {
		return errors.New("linked list is empty")
	}

	for i := 0; i < Linked_list.Length; i++ {
		if node.Value == Value {
			if i > 0 {
				prevNode := Linked_list.GetAt(i - 1)
				prevNode.Next = Linked_list.GetAt(i).Next
			} else {
				Linked_list.Head = node.Next
			}
			Linked_list.Length--
			return nil
		}
		node = node.Next
	}
	return errors.New("node not found")
}

func main() {

	var Linked_list = Linkedlist{}
	Linked_list.Insert(1)
	Linked_list.Insert(2)
	Linked_list.Insert(3)

	Linked_list.Display()

	valueToBeSearch := 2
	positiom, err := Linked_list.Search(valueToBeSearch)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Position of value %d is %d \n\n", valueToBeSearch, positiom)
	}

	valueToBeDelete := 1
	err = Linked_list.Delete(valueToBeDelete)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Node of value: %d is deleted \n\n", valueToBeDelete)
	}

	Linked_list.Display()
}
