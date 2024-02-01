package main

import "fmt"

type UnionFindCC struct {
	parent, rank []int
}

// brute force

// TC - O(α(N)) / O(E), SC - O(N)
func NewUnionFindCC(size int) *UnionFindCC {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}
	return &UnionFindCC{parent, rank}
}

func (uf *UnionFindCC) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path Compression
	}
	return uf.parent[x]
}

func (uf *UnionFindCC) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

func countComponents(n int, edges [][]int) int {
	uf := NewUnionFind(n)
	for _, edge := range edges {
		uf.Union(edge[0], edge[1])
	}

	componentCount := 0
	for i := 0; i < n; i++ {
		if uf.Find(i) == i {
			componentCount++
		}
	}
	return componentCount
}

func main() {
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}
	fmt.Println(countComponents(5, edges)) // Output: 2
}
