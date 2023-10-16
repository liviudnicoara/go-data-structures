package main

// Problem Statement:
// Given an undirected graph, implement DFS to print all vertices in depth-first order.

import "fmt"

type Graph struct {
	Nodes map[int][]int
}

func (g *Graph) addEdge(src, dest int) {
	g.Nodes[src] = append(g.Nodes[src], dest)
	g.Nodes[dest] = append(g.Nodes[dest], src)
}

func dfs(graph Graph, start int, visited map[int]bool) {
	if visited[start] {
		return
	}

	fmt.Print(start, " ")
	visited[start] = true

	for _, neighbor := range graph.Nodes[start] {
		dfs(graph, neighbor, visited)
	}
}

func main() {
	graph := Graph{
		Nodes: make(map[int][]int),
	}

	graph.addEdge(0, 1)
	graph.addEdge(0, 2)
	graph.addEdge(1, 3)
	graph.addEdge(2, 4)
	graph.addEdge(3, 5)
	graph.addEdge(4, 5)

	startNode := 0
	visited := make(map[int]bool)

	fmt.Println("DFS Traversal:")
	dfs(graph, startNode, visited)
}
