package main

import (
	"fmt"
)

// The basic operations are mentioned as follows -
type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	Key      int
	Adjacent []*Vertex
}

// Add/Remove Vertex:= Add or remove a vertex in a graph.
func (graph *Graph) AddVertex(data int) {
	if Contains(graph.Vertices, data) {
		fmt.Printf("Vertex %v is an existing key. \n", data)
		return
	}
	graph.Vertices = append(graph.Vertices, &Vertex{Key: data})
}

// Add/Remove Edge:= Add or remove an edge between two vertices.
func (graph *Graph) AddEdge(from, to int) {
	from_vertex := graph.Getvertex(from)
	to_vertex := graph.Getvertex(to)

	if from_vertex == nil || to_vertex == nil {
		fmt.Printf("Invalid edge (%v->%v)", from, to)
		return
	} else if Contains(from_vertex.Adjacent, to) {
		fmt.Printf("Edge (%v->%v) Exists.", from, to)
		return
	}
	from_vertex.Adjacent = append(from_vertex.Adjacent, to_vertex)
}

// find the path from one vertex to another vertex.
func (graph *Graph) Getvertex(data int) *Vertex {
	for i, value := range graph.Vertices {
		if data == value.Key {
			return graph.Vertices[i]
		}
	}
	return nil
}

func Contains(list_of_vertex []*Vertex, data int) bool {
	for _, value := range list_of_vertex {
		if data == value.Key {
			return true
		}
	}
	return false
}

// Check if the graph contains a given value.
func (graph *Graph) Print() {
	for _, v := range graph.Vertices {
		fmt.Printf("Vertex %v: ", v.Key)
		for _, v := range v.Adjacent {
			fmt.Printf("%v", v.Key)
		}
		fmt.Println()
	}
}

func main() {
	graph := &Graph{}

	for i := 0; i < 5; i++ {
		graph.AddVertex(i)

	}
	graph.AddEdge(2, 1)
	graph.AddEdge(4, 3)
	graph.AddEdge(4, 1)
	graph.AddEdge(1, 4)
	graph.AddEdge(3, 2)

	graph.Print()
}
