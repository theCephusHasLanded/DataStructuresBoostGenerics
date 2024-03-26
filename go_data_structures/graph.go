package go_data_structures

import (
	"golang.org/x/exp/constraints"
)

// Graph represents a graph with nodes of generic type T.
type Graph[T constraints.Ordered] struct {
	nodes map[T]struct{}         // Set of nodes to ensure uniqueness.
	edges map[T]map[T]struct{}   // Adjacency list to represent edges.
}

// NewGraph creates a new instance of a graph.
func NewGraph[T constraints.Ordered]() *Graph[T] {
	return &Graph[T]{
		nodes: make(map[T]struct{}),
		edges: make(map[T]map[T]struct{}),
	}
}

// insert adds a node to the graph and initializes its adjacency list if not present.
func (g *Graph[T]) insert(val T) {
	if _, exists := g.nodes[val]; !exists {
		g.nodes[val] = struct{}{}
		g.edges[val] = make(map[T]struct{})
	}
}

// remove deletes a node from the graph and removes it from other nodes' adjacency lists.
func (g *Graph[T]) remove(val T) {
	if _, exists := g.nodes[val]; exists {
		delete(g.nodes, val)
		delete(g.edges, val)
		for node := range g.edges {
			delete(g.edges[node], val)
		}
	}
}

// Empty checks if the graph is empty.
func (g *Graph[T]) Empty() bool {
	return len(g.nodes) == 0
}

// DepthFirstTraversal performs a depth-first traversal starting from root.
func (g *Graph[T]) DepthFirstTraversal(root T) []T {
	var result []T
	visited := make(map[T]bool)
	var dfs func(T)
	dfs = func(node T) {
		if visited[node] {
			return
		}
		visited[node] = true
		result = append(result, node)
		for neighbor := range g.edges[node] {
			dfs(neighbor)
		}
	}
	dfs(root)
	return result
}

// BreadthFirstTraversal performs a breadth-first traversal starting from root.
func (g *Graph[T]) BreadthFirstTraversal(root T) []T {
	var result []T
	visited := make(map[T]bool)
	queue := []T{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if visited[node] {
			continue
		}
		visited[node] = true
		result = append(result, node)
		for neighbor := range g.edges[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
	return result
}

// DepthFirstSearch searches for a value in the graph using depth-first search.
func (g *Graph[T]) DepthFirstSearch(val T) bool {
	visited := make(map[T]bool)
	var dfs func(T) bool
	dfs = func(node T) bool {
		if node == val {
			return true
		}
		visited[node] = true
		for neighbor := range g.edges[node] {
			if !visited[neighbor] && dfs(neighbor) {
				return true
			}
		}
		return false
	}
	for node := range g.nodes {
		if !visited[node] && dfs(node) {
			return true
		}
	}
	return false
}

// BreadthFirstSearch searches for a value in the graph using breadth-first search.
func (g *Graph[T]) BreadthFirstSearch(val T) bool {
	visited := make(map[T]bool)
	queue := []T{}
	for node := range g.nodes {
		queue = append(queue, node) // Start search from every node
		break                        // But actually start from the first node only, can be modified if needed
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == val {
			return true
		}
		if visited[node] {
			continue
		}
		visited[node] = true
		for neighbor := range g.edges[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
	return false
}

// Size returns the number of nodes in the graph.
func (g *Graph[T]) Size() int {
	return len(g.nodes)
}
