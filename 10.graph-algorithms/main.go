package main

// Problem Statement:
// Given an undirected graph, implement DFS to print all vertices in depth-first order.

import (
	"container/list"
	"fmt"
)

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

func bfs(graph Graph, start int, visited map[int]bool) {
	queue := list.New()
	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		node := queue.Front().Value.(int)
		fmt.Print(node, " ")
		queue.Remove(queue.Front())

		for _, neighbor := range graph.Nodes[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
			}
		}
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

	startNode = 0
	visited = make(map[int]bool)

	fmt.Println("BFS Traversal:")
	bfs(graph, startNode, visited)
}
