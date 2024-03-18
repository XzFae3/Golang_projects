package main

import (
	"fmt"
)

// La structure du Graph;
// Graph G; non-orienté
// Noeuds

type Graph struct {
	nodes map[string][]string
}

func graph() *Graph {
	return &Graph{
		nodes: make(map[string][]string),
	}
}

func (graph *Graph) addVertex(vertex string) {
	graph.nodes[vertex] = []string{}
}

func (graph *Graph) addEdges(extremite1, extremite2 string) {
	graph.nodes[extremite1] = append(graph.nodes[extremite1], extremite2)
	graph.nodes[extremite2] = append(graph.nodes[extremite2], extremite1)
}

func (graph *Graph) print() {
	for nodes, neighboors := range graph.nodes {
		fmt.Println("->", nodes, neighboors)
	}
}
