// Package utils provides common utility functions for Advent of Code solutions.
package utils

// UnionFind implements a disjoint-set data structure with path compression and union by rank.
// This is useful for efficiently tracking connected components.
type UnionFind struct {
	parent map[int]int
	rank   map[int]int
}

// NewUnionFind creates a new UnionFind structure.
func NewUnionFind() *UnionFind {
	return &UnionFind{
		parent: make(map[int]int),
		rank:   make(map[int]int),
	}
}

// MakeSet adds a new element to the UnionFind structure.
func (uf *UnionFind) MakeSet(x int) {
	if _, exists := uf.parent[x]; !exists {
		uf.parent[x] = x
		uf.rank[x] = 0
	}
}

// Find returns the root of the set containing x, with path compression.
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

// Union merges the sets containing x and y using union by rank.
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	
	if rootX == rootY {
		return
	}
	
	// Union by rank
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
}

// Connected returns true if x and y are in the same set.
func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// CountSets returns the number of disjoint sets.
func (uf *UnionFind) CountSets() int {
	roots := make(map[int]bool)
	for node := range uf.parent {
		roots[uf.Find(node)] = true
	}
	return len(roots)
}