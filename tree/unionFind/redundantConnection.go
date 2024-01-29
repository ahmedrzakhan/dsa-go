package main

import "fmt"

/**
684. Redundant Connection

In this problem, a tree is an undirected graph that is connected and has no cycles.

You are given a graph that started as a tree with n nodes labeled from 1 to n, with one
additional edge added. The added edge has two different vertices chosen from 1 to n,
and was not an edge that already existed. The graph is represented as an array edges of
length n where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and
bi in the graph.

Return an edge that can be removed so that the resulting graph is a tree of n nodes. If
there are multiple answers, return the answer that occurs last in the input.

Example 1:

Input: edges = [[1,2],[1,3],[2,3]]
Output: [2,3]
Example 2:

Input: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
Output: [1,4]

Constraints:

n == edges.length
3 <= n <= 1000
edges[i].length == 2
1 <= ai < bi <= edges.length
ai != bi
There are no repeated edges.
The given graph is connected.
*/

/*
Brute force: create adjaceny list and dfs, like done in detecting cycels in graph
*/
type UnionFindRC struct {
	parent, rank []int
}

// TC - O(N*α(N)), SC - O(N)
func NewUnionFindRC(size int) *UnionFindRC {
	parent := make([]int, size+1)
	rank := make([]int, size+1)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFindRC{parent, rank}
}

func (uf *UnionFindRC) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFindRC) Union(x, y int) bool {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX == rootY {
		return false
	}
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	return true
}

func findRedundantConnection(edges [][]int) []int {
	uf := NewUnionFind(len(edges))
	for _, edge := range edges {
		if !uf.Union(edge[0], edge[1]) {
			return edge
		}
	}
	return nil
}

func mainRC() {
	edges := [][]int{{1, 2}, {1, 3}, {2, 3}}
	fmt.Println(findRedundantConnection(edges)) // Output: [2, 3]
}
