package main

import "fmt"

// UnionFind structure
type UnionFind struct {
	par  map[int]int
	rank map[int]int
}

// NewUnionFind initializes a new UnionFind instance
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		par:  make(map[int]int),
		rank: make(map[int]int),
	}

	for i := 1; i <= n; i++ {
		uf.par[i] = i
		uf.rank[i] = 0
	}

	return uf
}

// Find with path compression
func (uf *UnionFind) Find(n int) int {
	if uf.par[n] != n {
		uf.par[n] = uf.Find(uf.par[n]) // Path compression
	}
	return uf.par[n]
}

// Union by rank, returns false if n1 and n2 are already connected
func (uf *UnionFind) Union(n1, n2 int) bool {
	p1, p2 := uf.Find(n1), uf.Find(n2)
	if p1 == p2 {
		return false
	}

	if uf.rank[p1] > uf.rank[p2] {
		uf.par[p2] = p1
	} else if uf.rank[p1] < uf.rank[p2] {
		uf.par[p1] = p2
	} else {
		uf.par[p1] = p2
		uf.rank[p2]++
	}

	return true
}

// Example to demonstrate the usage of UnionFind
func mainUF() {
	uf := NewUnionFind(5) // Create a UnionFind instance for 5 elements

	// Perform some unions
	fmt.Println("Union(1, 2):", uf.Union(1, 2)) // Connect 1 and 2
	fmt.Println("Union(2, 3):", uf.Union(2, 3)) // Connect 2 and 3
	fmt.Println("Union(4, 5):", uf.Union(4, 5)) // Connect 4 and 5

	// Check if two elements are connected
	fmt.Println("Find(1) == Find(3):", uf.Find(1) == uf.Find(3)) // Should be true, 1 and 3 are connected through 2
	fmt.Println("Find(1) == Find(4):", uf.Find(1) == uf.Find(4)) // Should be false, 1 and 4 are not connected
	fmt.Println("uf.Find(1)", uf.Find(1))
	// Try to union already connected components
	fmt.Println("Union(1, 3):", uf.Union(1, 3)) // Should be false, 1 and 3 are already connected
}
