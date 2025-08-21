/*
 * Day 22: Mode Maze
 * 
 * Part 1: Calculate the total risk level for a cave system
 * Each region has a geologic index, erosion level, and type (rocky/wet/narrow)
 * Risk level is the sum of all region types in the rectangle from 0,0 to target
 * 
 * Part 2: Find the shortest path to the target with tool switching
 * Different tools can be used in different terrain types
 * Switching tools takes 7 minutes, moving takes 1 minute
 */

package day22

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

// Region types
const (
	rocky  = 0
	wet    = 1
	narrow = 2
)

// Tools
const (
	neither = 0
	torch   = 1
	gear    = 2
)

// Costs
const (
	moveCost   = 1
	switchCost = 7
)

// Search parameters
const (
	searchMargin = 50
)

type Solution struct {
	input string
}

func New(input string) *Solution {
	return &Solution{input: strings.TrimSpace(input)}
}

// State for pathfinding
type state struct {
	x, y, tool, time int
	index            int // for heap
}

// Priority queue implementation for Dijkstra's algorithm
type priorityQueue []*state

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].time < pq[j].time }
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*state)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (s *Solution) Part1() (int, error) {
	depth, targetX, targetY, err := s.parseInput()
	if err != nil {
		return 0, err
	}

	// Cache for erosion levels
	erosionLevel := make(map[string]int)

	// Calculate erosion level for a given coordinate
	var getErosionLevel func(x, y int) int
	getErosionLevel = func(x, y int) int {
		key := fmt.Sprintf("%d,%d", x, y)
		if level, exists := erosionLevel[key]; exists {
			return level
		}

		var geoIndex int
		if (x == 0 && y == 0) || (x == targetX && y == targetY) {
			geoIndex = 0
		} else if y == 0 {
			geoIndex = x * 16807
		} else if x == 0 {
			geoIndex = y * 48271
		} else {
			// Multiply erosion levels of adjacent regions
			leftErosion := getErosionLevel(x-1, y)
			topErosion := getErosionLevel(x, y-1)
			geoIndex = leftErosion * topErosion
		}

		level := (geoIndex + depth) % 20183
		erosionLevel[key] = level
		return level
	}

	// Calculate total risk level
	totalRisk := 0
	for y := 0; y <= targetY; y++ {
		for x := 0; x <= targetX; x++ {
			erosion := getErosionLevel(x, y)
			regionType := erosion % 3
			totalRisk += regionType
		}
	}

	return totalRisk, nil
}

func (s *Solution) Part2() (int, error) {
	depth, targetX, targetY, err := s.parseInput()
	if err != nil {
		return 0, err
	}

	// Cache for erosion levels
	erosionLevel := make(map[string]int)
	var getErosionLevel func(x, y int) int
	getErosionLevel = func(x, y int) int {
		key := fmt.Sprintf("%d,%d", x, y)
		if level, exists := erosionLevel[key]; exists {
			return level
		}

		var geoIndex int
		if (x == 0 && y == 0) || (x == targetX && y == targetY) {
			geoIndex = 0
		} else if y == 0 {
			geoIndex = x * 16807
		} else if x == 0 {
			geoIndex = y * 48271
		} else {
			leftErosion := getErosionLevel(x-1, y)
			topErosion := getErosionLevel(x, y-1)
			geoIndex = leftErosion * topErosion
		}

		level := (geoIndex + depth) % 20183
		erosionLevel[key] = level
		return level
	}

	// Get region type
	getRegionType := func(x, y int) int {
		return getErosionLevel(x, y) % 3
	}

	// Valid tools per region type:
	// Rocky: climbing gear or torch
	// Wet: climbing gear or neither
	// Narrow: torch or neither
	isValidTool := func(regionType, tool int) bool {
		switch regionType {
		case rocky:
			return tool == torch || tool == gear
		case wet:
			return tool == neither || tool == gear
		case narrow:
			return tool == neither || tool == torch
		}
		return false
	}

	// Dijkstra's algorithm with state = (x, y, tool)
	pq := make(priorityQueue, 0)
	heap.Init(&pq)

	// Start at 0,0 with torch equipped
	start := &state{x: 0, y: 0, tool: torch, time: 0}
	heap.Push(&pq, start)

	// Track visited states and best-known times
	visited := make(map[string]bool)
	best := make(map[string]int)

	// Directions: up, down, left, right
	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}

	// Search for shortest path
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*state)

		// Create state key
		stateKey := fmt.Sprintf("%d,%d,%d", current.x, current.y, current.tool)
		if visited[stateKey] {
			continue
		}
		visited[stateKey] = true

		// Check if we reached the target with torch equipped
		if current.x == targetX && current.y == targetY && current.tool == torch {
			return current.time, nil
		}

		// Try moving to adjacent regions
		for i := 0; i < 4; i++ {
			nx, ny := current.x+dx[i], current.y+dy[i]

			// Skip if out of bounds or negative
			if nx < 0 || ny < 0 {
				continue
			}

			// Don't explore too far from target
			if nx > targetX+searchMargin || ny > targetY+searchMargin {
				continue
			}

			nextRegionType := getRegionType(nx, ny)
			if isValidTool(nextRegionType, current.tool) {
				nextStateKey := fmt.Sprintf("%d,%d,%d", nx, ny, current.tool)
				if !visited[nextStateKey] {
					nt := current.time + moveCost
					if bt, ok := best[nextStateKey]; !ok || nt < bt {
						best[nextStateKey] = nt
						heap.Push(&pq, &state{x: nx, y: ny, tool: current.tool, time: nt})
					}
				}
			}
		}

		// Try switching tools at current position
		currentRegionType := getRegionType(current.x, current.y)
		for newTool := neither; newTool <= gear; newTool++ {
			if newTool != current.tool && isValidTool(currentRegionType, newTool) {
				nextStateKey := fmt.Sprintf("%d,%d,%d", current.x, current.y, newTool)
				if !visited[nextStateKey] {
					nt := current.time + switchCost
					if bt, ok := best[nextStateKey]; !ok || nt < bt {
						best[nextStateKey] = nt
						heap.Push(&pq, &state{x: current.x, y: current.y, tool: newTool, time: nt})
					}
				}
			}
		}
	}

	return -1, fmt.Errorf("no path found")
}

func (s *Solution) parseInput() (depth, targetX, targetY int, err error) {
	lines := strings.Split(s.input, "\n")
	if len(lines) < 2 {
		return 0, 0, 0, fmt.Errorf("invalid input format")
	}

	// Parse depth
	depthLine := strings.TrimSpace(lines[0])
	if !strings.HasPrefix(depthLine, "depth: ") {
		return 0, 0, 0, fmt.Errorf("invalid depth line: %s", depthLine)
	}
	depth, err = strconv.Atoi(strings.TrimPrefix(depthLine, "depth: "))
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid depth value: %v", err)
	}

	// Parse target
	targetLine := strings.TrimSpace(lines[1])
	if !strings.HasPrefix(targetLine, "target: ") {
		return 0, 0, 0, fmt.Errorf("invalid target line: %s", targetLine)
	}
	targetStr := strings.TrimPrefix(targetLine, "target: ")
	coords := strings.SplitN(targetStr, ",", 2)
	if len(coords) != 2 {
		return 0, 0, 0, fmt.Errorf("invalid target coordinates: %s", targetStr)
	}
	targetX, err = strconv.Atoi(strings.TrimSpace(coords[0]))
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid target X: %v", err)
	}
	targetY, err = strconv.Atoi(strings.TrimSpace(coords[1]))
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid target Y: %v", err)
	}

	return depth, targetX, targetY, nil
}