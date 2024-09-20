package main

import "fmt"

//The basic operationns are mentioned as follows -
type Node struct {
	Data      int
	LeftNode  *Node
	RightNode *Node
}

type Tree struct {
	Root   *Node
	Lenght int
}

//Inserting an element
func (node *Node) Insert(data int) {
	if node.Data < data {
		if node.RightNode == nil {
			node.RightNode = &Node{Data: data}
		} else {
			node.RightNode.Insert(data)
		}
	} else if node.Data > data {
		if node.LeftNode == nil {
			node.LeftNode = &Node{Data: data}
		} else {
			node.LeftNode.Insert(data)
		}
	}
}

//Searchimg for an element
func (node *Node) Search(data int) bool {
	if node == nil {
		return false
	}

	if node.Data < data {
		return node.RightNode.Search(data)
	} else if node.Data > data {
		return node.LeftNode.Search(data)
	}

	return true
}

func PrintTreeData(node *Node) {
	if node == nil {
		fmt.Println()
		return
	}

	PrintTreeData(node.LeftNode)
	fmt.Printf("%d", node.Data)
	PrintTreeData(node.RightNode)
}

func main() {
	tree := &Node{Data: 100}
	tree.Insert(20)
	tree.Insert(50)
	tree.Insert(1000)
	tree.Insert(200)
	data := 200
	fmt.Printf("Data %d available in Trees: %t \n", data, tree.Search(data))
	PrintTreeData(tree)
}
