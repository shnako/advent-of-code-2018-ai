/*
 * Day 17: Reservoir Research
 *
 * Part 1: Simulate water flow from a spring through sand and clay veins.
 * Water flows down when possible, spreads horizontally when blocked by clay.
 * Count all water tiles (flowing and settled) within the Y range of clay veins.
 *
 * Part 2: Count only settled water tiles (tiles where water comes to rest).
 */

package day17

import (
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

type Solution struct {
	input string
	clay  map[Point]bool
	water map[Point]rune // '|' for flowing, '~' for settled
	minX  int
	maxX  int
	minY  int
	maxY  int
}

func New(input string) *Solution {
	return &Solution{
		input: strings.TrimSpace(input),
		clay:  make(map[Point]bool),
		water: make(map[Point]rune),
	}
}

func (s *Solution) Part1() (int, error) {
	err := s.parseInput()
	if err != nil {
		return 0, err
	}

	s.simulateWaterFlow()

	// Count all water tiles within Y bounds
	count := 0
	for p := range s.water {
		if p.Y >= s.minY && p.Y <= s.maxY {
			count++
		}
	}

	return count, nil
}

func (s *Solution) Part2() (int, error) {
	err := s.parseInput()
	if err != nil {
		return 0, err
	}

	// Reset water map for Part 2 
	s.water = make(map[Point]rune)
	s.simulateWaterFlow()

	// Count only settled water tiles ('~') within Y bounds
	count := 0
	for p, waterType := range s.water {
		if p.Y >= s.minY && p.Y <= s.maxY && waterType == '~' {
			count++
		}
	}

	return count, nil
}

func (s *Solution) parseInput() error {
	lines := strings.Split(s.input, "\n")
	
	// Regex patterns for parsing clay coordinates
	xPattern := regexp.MustCompile(`x=(\d+)(?:\.\.(\d+))?`)
	yPattern := regexp.MustCompile(`y=(\d+)(?:\.\.(\d+))?`)

	s.minX = 999999
	s.maxX = -999999
	s.minY = 999999
	s.maxY = -999999

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		xMatch := xPattern.FindStringSubmatch(line)
		yMatch := yPattern.FindStringSubmatch(line)

		if len(xMatch) < 2 || len(yMatch) < 2 {
			continue
		}

		// Parse X range
		x1, _ := strconv.Atoi(xMatch[1])
		x2 := x1
		if len(xMatch) > 2 && xMatch[2] != "" {
			x2, _ = strconv.Atoi(xMatch[2])
		}

		// Parse Y range
		y1, _ := strconv.Atoi(yMatch[1])
		y2 := y1
		if len(yMatch) > 2 && yMatch[2] != "" {
			y2, _ = strconv.Atoi(yMatch[2])
		}

		// Add clay tiles
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				s.clay[Point{x, y}] = true
				
				// Update bounds
				if x < s.minX {
					s.minX = x
				}
				if x > s.maxX {
					s.maxX = x
				}
				if y < s.minY {
					s.minY = y
				}
				if y > s.maxY {
					s.maxY = y
				}
			}
		}
	}

	return nil
}

func (s *Solution) simulateWaterFlow() {
	// Start water flow from spring at (500, 0)
	s.flowFrom(500, 0)
}

func (s *Solution) flowFrom(x, y int) bool {
	// If we're below the max Y, stop
	if y > s.maxY {
		return false
	}
	
	current := Point{x, y}
	if _, hasWater := s.water[current]; hasWater {
		return s.water[current] == '~'
	}

	// Mark this position as flowing water
	s.water[current] = '|'

	// Try to flow down first
	below := Point{x, y + 1}
	settledBelow := false
	
	if !s.clay[below] {
		settledBelow = s.flowFrom(x, y+1)
	} else {
		settledBelow = true // Clay acts as settled ground
	}

	if settledBelow {
		// Can't flow down, try to spread horizontally
		leftBlocked := s.fillHorizontal(x-1, y, -1)
		rightBlocked := s.fillHorizontal(x+1, y, 1)

		// If both sides are blocked, this becomes settled water
		if leftBlocked && rightBlocked {
			s.water[current] = '~'
			// Fill the entire row with settled water
			s.settleRow(x, y)
			return true
		}
	}

	return false
}

func (s *Solution) fillHorizontal(x, y, direction int) bool {
	current := Point{x, y}
	
	// If we hit clay, we're blocked
	if s.clay[current] {
		return true
	}

	// If we're beyond bounds, we can't be blocked
	if y > s.maxY {
		return false
	}

	// Mark as flowing water
	s.water[current] = '|'

	// Check if we can flow down from here
	below := Point{x, y + 1}
	if !s.clay[below] {
		settledBelow := s.flowFrom(x, y+1)
		if !settledBelow {
			return false // Water flows down, so we're not blocked
		}
	}

	// Continue filling horizontally
	return s.fillHorizontal(x+direction, y, direction)
}

func (s *Solution) settleRow(centerX, y int) {
	// Find the extent of water that should be settled on this row
	leftBound := centerX
	rightBound := centerX
	
	// Find leftmost water tile that should be settled
	for x := centerX - 1; x >= s.minX-1; x-- {
		if s.clay[Point{x, y}] {
			break
		}
		if _, hasWater := s.water[Point{x, y}]; !hasWater {
			break
		}
		leftBound = x
	}
	
	// Find rightmost water tile that should be settled
	for x := centerX + 1; x <= s.maxX+1; x++ {
		if s.clay[Point{x, y}] {
			break
		}
		if _, hasWater := s.water[Point{x, y}]; !hasWater {
			break
		}
		rightBound = x
	}
	
	// Settle all water tiles in this range
	for x := leftBound; x <= rightBound; x++ {
		if _, hasWater := s.water[Point{x, y}]; hasWater {
			s.water[Point{x, y}] = '~'
		}
	}
}