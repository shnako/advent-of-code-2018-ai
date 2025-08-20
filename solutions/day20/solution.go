/*
 * Day 20: A Regular Map
 * 
 * Part 1: Parse a regex pattern describing room connections and find the furthest room.
 * The regex represents all possible paths through a facility, where N/S/E/W indicate movement.
 * Parentheses group options, pipes separate alternatives, and empty alternatives are allowed.
 * We build a graph of all rooms and doors, then use BFS to find shortest paths to all rooms.
 * 
 * Part 2: Count how many rooms have a shortest path from the start that passes through at least 1000 doors.
 * Using the same BFS result from Part 1, we count rooms with distance >= 1000.
 */

package day20

import (
	"strings"
)

type Solution struct {
	input string
}

func New(input string) *Solution {
	return &Solution{input: strings.TrimSpace(input)}
}

type Point struct {
	x, y int
}

func (s *Solution) Part1() (int, error) {
	regex := s.input
	if regex[0] != '^' || regex[len(regex)-1] != '$' {
		return 0, nil
	}
	regex = regex[1 : len(regex)-1] // Remove ^ and $
	
	// Build the map
	doors := make(map[Point]map[Point]bool)
	distances := s.explore(regex, doors)
	
	// Find the maximum distance
	maxDist := 0
	for _, dist := range distances {
		if dist > maxDist {
			maxDist = dist
		}
	}
	
	return maxDist, nil
}

func (s *Solution) Part2() (int, error) {
	regex := s.input
	if regex[0] != '^' || regex[len(regex)-1] != '$' {
		return 0, nil
	}
	regex = regex[1 : len(regex)-1] // Remove ^ and $
	
	// Build the map
	doors := make(map[Point]map[Point]bool)
	distances := s.explore(regex, doors)
	
	// Count rooms with distance >= 1000
	count := 0
	for _, dist := range distances {
		if dist >= 1000 {
			count++
		}
	}
	
	return count, nil
}

func (s *Solution) explore(regex string, doors map[Point]map[Point]bool) map[Point]int {
	// Parse the regex and build the graph
	stack := []Point{{0, 0}}
	current := Point{0, 0}
	
	for i := 0; i < len(regex); i++ {
		switch regex[i] {
		case 'N':
			next := Point{current.x, current.y - 1}
			s.addDoor(doors, current, next)
			current = next
		case 'S':
			next := Point{current.x, current.y + 1}
			s.addDoor(doors, current, next)
			current = next
		case 'E':
			next := Point{current.x + 1, current.y}
			s.addDoor(doors, current, next)
			current = next
		case 'W':
			next := Point{current.x - 1, current.y}
			s.addDoor(doors, current, next)
			current = next
		case '(':
			stack = append(stack, current)
		case '|':
			current = stack[len(stack)-1]
		case ')':
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	
	// BFS to find distances to all rooms
	distances := make(map[Point]int)
	distances[Point{0, 0}] = 0
	queue := []Point{{0, 0}}
	
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		
		if neighbors, ok := doors[pos]; ok {
			for neighbor := range neighbors {
				if _, visited := distances[neighbor]; !visited {
					distances[neighbor] = distances[pos] + 1
					queue = append(queue, neighbor)
				}
			}
		}
	}
	
	return distances
}

func (s *Solution) addDoor(doors map[Point]map[Point]bool, from, to Point) {
	if doors[from] == nil {
		doors[from] = make(map[Point]bool)
	}
	if doors[to] == nil {
		doors[to] = make(map[Point]bool)
	}
	doors[from][to] = true
	doors[to][from] = true
}