package day25

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z, t int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(p1, p2 Point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z) + abs(p1.t-p2.t)
}

type UnionFind struct {
	parent map[int]int
	rank   map[int]int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		parent: make(map[int]int),
		rank:   make(map[int]int),
	}
}

func (uf *UnionFind) makeSet(x int) {
	if _, exists := uf.parent[x]; !exists {
		uf.parent[x] = x
		uf.rank[x] = 0
	}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) union(x, y int) {
	rootX := uf.find(x)
	rootY := uf.find(y)
	
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

func (uf *UnionFind) countSets() int {
	roots := make(map[int]bool)
	for node := range uf.parent {
		roots[uf.find(node)] = true
	}
	return len(roots)
}

func parseInput(input string) []Point {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	points := make([]Point, 0, len(lines))
	
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 4 {
			continue
		}
		
		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		z, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
		t, _ := strconv.Atoi(strings.TrimSpace(parts[3]))
		
		points = append(points, Point{x, y, z, t})
	}
	
	return points
}

func Part1(input string) string {
	points := parseInput(input)
	
	// Create Union-Find structure
	uf := NewUnionFind()
	
	// Initialize each point as its own set
	for i := range points {
		uf.makeSet(i)
	}
	
	// Connect points that are within Manhattan distance of 3
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if manhattanDistance(points[i], points[j]) <= 3 {
				uf.union(i, j)
			}
		}
	}
	
	// Count the number of constellations
	return fmt.Sprintf("%d", uf.countSets())
}

func Part2(input string) string {
	// Day 25 has no Part 2 - it's just a Christmas star!
	return "Merry Christmas!"
}