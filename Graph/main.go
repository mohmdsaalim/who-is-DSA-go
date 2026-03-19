package main

import "fmt"

type Graph struct {
	vertices map[int][]int
}

// Add Vertex
func (g *Graph) AddVertex(v int) {
	if g.vertices == nil {
		g.vertices = make(map[int][]int)
	}
	g.vertices[v] = []int{}
}

// Add Edge (undirected)
func (g *Graph) AddEdge(v1, v2 int) {
	g.vertices[v1] = append(g.vertices[v1], v2)
	g.vertices[v2] = append(g.vertices[v2], v1)
}

// Print
func (g *Graph) Print() {
	for v, adj := range g.vertices {
		fmt.Println(v, "->", adj)
	}
}

// BFS
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	visited[start] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		fmt.Print(curr, " ")

		for _, n := range g.vertices[curr] {
			if !visited[n] {
				visited[n] = true
				queue = append(queue, n)
			}
		}
	}
	fmt.Println()
}

// DFS
func (g *Graph) DFS(v int, visited map[int]bool) {
	visited[v] = true
	fmt.Print(v, " ")

	for _, n := range g.vertices[v] {
		if !visited[n] {
			g.DFS(n, visited)
		}
	}
}

func main() {
	g := &Graph{}

	// vertices
	for i := 1; i <= 5; i++ {
		g.AddVertex(i)
	}

	// edges
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 5)

	g.Print()

	fmt.Print("BFS: ")
	g.BFS(1)

	fmt.Print("DFS: ")
	g.DFS(1, make(map[int]bool))
}