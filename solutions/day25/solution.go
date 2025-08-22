package day25

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shnako/advent-of-code-2018-ai/internal/utils"
)

type Point struct {
	x, y, z, t int
}


func manhattanDistance(p1, p2 Point) int {
	return utils.Abs(p1.x-p2.x) + utils.Abs(p1.y-p2.y) + utils.Abs(p1.z-p2.z) + utils.Abs(p1.t-p2.t)
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
		
		x, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			continue // Skip invalid coordinate lines
		}
		y, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			continue // Skip invalid coordinate lines
		}
		z, err := strconv.Atoi(strings.TrimSpace(parts[2]))
		if err != nil {
			continue // Skip invalid coordinate lines
		}
		t, err := strconv.Atoi(strings.TrimSpace(parts[3]))
		if err != nil {
			continue // Skip invalid coordinate lines
		}
		
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