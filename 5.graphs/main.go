package main

import "fmt"

//Problem Statement:
// Given an undirected graph and two nodes, find the length of the shortest path between them.
// tip: use BFS (Breadth-First Search)

type Graph struct {
	Nodes map[int][]int
}

func (g *Graph) addEdge(src, dest int) {
	g.Nodes[src] = append(g.Nodes[src], dest)
	g.Nodes[dest] = append(g.Nodes[dest], src)
}

func shortestPath(g Graph, start, end int) int {
	q := make(chan int, len(g.Nodes))
	visited := map[int]bool{}
	distance := 0

	q <- start
	visited[start] = true

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := <-q

			if node == end {
				return distance
			}

			for _, neighbor := range g.Nodes[node] {
				if !visited[neighbor] {
					q <- neighbor
					visited[neighbor] = true
				}
			}
		}

		distance++
	}

	return -1
}

func main() {
	graph := Graph{
		Nodes: make(map[int][]int),
	}

	// Adding edges to the graph
	graph.addEdge(0, 1)
	graph.addEdge(0, 2)
	graph.addEdge(1, 3)
	graph.addEdge(2, 4)
	graph.addEdge(3, 5)
	graph.addEdge(4, 5)

	startNode := 0
	endNode := 5

	pathLength := shortestPath(graph, startNode, endNode)
	fmt.Printf("Shortest Path from %d to %d is %d\n", startNode, endNode, pathLength)
}
