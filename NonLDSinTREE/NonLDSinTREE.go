package main

import "fmt"

//The basic operations are mentioned as follow -
const AlphabetSize = 26

type Node struct {
	Children [AlphabetSize]*Node
	IsEnd    bool
}

type Tree struct {
	Root *Node
}

func InitTree() *Tree {
	return &Tree{Root: &Node{}}
}

//Insert - This operation is performed to add an element into the tree.
func (tree *Tree) Insert(word string) {
	currentNode := tree.Root
	for i := 0; i < len(word); i++ {
		charIndex := word[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			currentNode.Children[charIndex] = &Node{}
		}
		currentNode = currentNode.Children[charIndex]
	}
	currentNode.IsEnd = true
}

//Search - It is performed to search an element from the tree using the given key.
func (tree *Tree) Search(word string) bool {
	currentNode := tree.Root
	for i := 0; i < len(word); i++ {
		charIndex := word[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.Children[charIndex]
	}
	return currentNode.IsEnd
}

func main() {
	var tree = InitTree()

	var list_of_words = []string{
		"john", "jake", "james", "arya", "client",
	}
	for _, value := range list_of_words {
		tree.Insert(value)
	}

	fmt.Println("flek is present:", tree.Search("flek"))
	fmt.Println("John is present:", tree.Search("john"))
	fmt.Println("arya is present:", tree.Search("arya"))
	fmt.Println("amber is present:", tree.Search("amber"))

}
