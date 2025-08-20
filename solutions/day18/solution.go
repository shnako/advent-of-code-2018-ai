/*
 * Day 18: Settlers of The North Pole
 *
 * Part 1: Simulate cellular automata on a 50x50 grid for 10 minutes.
 * Each cell can be open ground (.), trees (|), or lumberyard (#).
 * Rules: Open -> trees (3+ adjacent trees), trees -> lumberyard (3+ adjacent lumberyards),
 * lumberyard -> open (unless adjacent to both trees and lumberyard).
 * Resource value = number of wooded acres * number of lumberyards.
 *
 * Part 2: Same simulation but for 1,000,000,000 minutes (needs cycle detection).
 */

package day18

import (
	"strings"
)

type Solution struct {
	input string
	grid  [][]rune
	width int
	height int
}

func New(input string) *Solution {
	return &Solution{input: strings.TrimSpace(input)}
}

func (s *Solution) Part1() (int, error) {
	s.parseInput()
	
	// Simulate for 10 minutes
	for minute := 0; minute < 10; minute++ {
		s.simulateMinute()
	}
	
	return s.calculateResourceValue(), nil
}

func (s *Solution) Part2() (int, error) {
	s.parseInput()
	
	// For Part 2, we need cycle detection since 1,000,000,000 minutes is too many to simulate
	seen := make(map[string]int)
	minute := 0
	target := 1000000000
	
	for minute < target {
		// Convert grid to string for cycle detection
		state := s.gridToString()
		if prevMinute, exists := seen[state]; exists {
			// Found a cycle! Calculate where we'll be at the target minute
			cycleLength := minute - prevMinute
			remaining := target - minute
			skipCycles := remaining / cycleLength
			minute += skipCycles * cycleLength
		} else {
			seen[state] = minute
		}
		
		if minute < target {
			s.simulateMinute()
			minute++
		}
	}
	
	return s.calculateResourceValue(), nil
}

func (s *Solution) parseInput() {
	lines := strings.Split(s.input, "\n")
	s.height = len(lines)
	if s.height > 0 {
		s.width = len(lines[0])
	}
	
	s.grid = make([][]rune, s.height)
	for i, line := range lines {
		s.grid[i] = []rune(line)
	}
}

func (s *Solution) simulateMinute() {
	newGrid := make([][]rune, s.height)
	for i := range newGrid {
		newGrid[i] = make([]rune, s.width)
	}
	
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			// Additional bounds check for debugging
			if y >= len(s.grid) || x >= len(s.grid[0]) {
				continue
			}
			newGrid[y][x] = s.getNextState(x, y)
		}
	}
	
	s.grid = newGrid
}

func (s *Solution) getNextState(x, y int) rune {
	// Bounds checking to prevent panic
	if y >= len(s.grid) || x >= len(s.grid[y]) {
		// Return default state if out of bounds
		return '.'
	}
	current := s.grid[y][x]
	trees, lumberyards, _ := s.countAdjacent(x, y)
	
	switch current {
	case '.': // Open ground
		if trees >= 3 {
			return '|' // Becomes trees
		}
	case '|': // Trees
		if lumberyards >= 3 {
			return '#' // Becomes lumberyard
		}
	case '#': // Lumberyard
		if lumberyards >= 1 && trees >= 1 {
			return '#' // Remains lumberyard
		} else {
			return '.' // Becomes open
		}
	}
	
	return current // No change
}

func (s *Solution) countAdjacent(x, y int) (trees, lumberyards, open int) {
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}
	
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if nx >= 0 && nx < s.width && ny >= 0 && ny < s.height {
			// Additional bounds checking to prevent panic
			if ny >= len(s.grid) || nx >= len(s.grid[ny]) {
				continue
			}
			switch s.grid[ny][nx] {
			case '.':
				open++
			case '|':
				trees++
			case '#':
				lumberyards++
			}
		}
	}
	
	return
}

func (s *Solution) calculateResourceValue() int {
	trees := 0
	lumberyards := 0
	
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			switch s.grid[y][x] {
			case '|':
				trees++
			case '#':
				lumberyards++
			}
		}
	}
	
	return trees * lumberyards
}

func (s *Solution) gridToString() string {
	var sb strings.Builder
	for _, row := range s.grid {
		sb.WriteString(string(row))
	}
	return sb.String()
}