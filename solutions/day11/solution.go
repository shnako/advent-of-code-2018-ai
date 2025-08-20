/*
 * Day 11: Chronal Charge
 * 
 * Part 1: Find the 3x3 square with the largest total power in a 300x300 grid.
 * Each fuel cell's power level is calculated using a specific formula involving
 * rack ID, coordinates, and grid serial number.
 * 
 * Part 2: Find the square of ANY size (1x1 to 300x300) with the largest total power.
 * Requires efficient computation using summed-area tables for optimal performance.
 */

package day11

import (
	"strconv"
	"strings"
)

type Solution struct {
	serialNumber int
	grid         [301][301]int // 1-indexed for convenience
	summedArea   [301][301]int // Summed-area table for efficient range queries
}

func New(input string) *Solution {
	input = strings.ReplaceAll(strings.TrimSpace(input), "\r\n", "\n")
	serialNumber, _ := strconv.Atoi(input)
	
	solution := &Solution{serialNumber: serialNumber}
	solution.calculateGrid()
	solution.buildSummedAreaTable()
	
	return solution
}

func (s *Solution) Part1() (string, error) {
	maxPower := -1000
	maxX, maxY := 0, 0
	
	// Check all possible 3x3 squares
	for x := 1; x <= 298; x++ {
		for y := 1; y <= 298; y++ {
			power := s.getSquareSum(x, y, 3)
			if power > maxPower {
				maxPower = power
				maxX, maxY = x, y
			}
		}
	}
	
	return strconv.Itoa(maxX) + "," + strconv.Itoa(maxY), nil
}

func (s *Solution) Part2() (string, error) {
	maxPower := -1000000
	maxX, maxY, maxSize := 0, 0, 0
	
	// Check all possible square sizes
	for size := 1; size <= 300; size++ {
		for x := 1; x <= 301-size; x++ {
			for y := 1; y <= 301-size; y++ {
				power := s.getSquareSum(x, y, size)
				if power > maxPower {
					maxPower = power
					maxX, maxY, maxSize = x, y, size
				}
			}
		}
	}
	
	return strconv.Itoa(maxX) + "," + strconv.Itoa(maxY) + "," + strconv.Itoa(maxSize), nil
}

// calculateGrid computes the power level for each fuel cell
func (s *Solution) calculateGrid() {
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			s.grid[x][y] = s.getPowerLevel(x, y)
		}
	}
}

// getPowerLevel calculates the power level for a fuel cell at (x,y)
func (s *Solution) getPowerLevel(x, y int) int {
	rackID := x + 10
	powerLevel := rackID * y
	powerLevel += s.serialNumber
	powerLevel *= rackID
	
	// Keep only hundreds digit
	hundredsDigit := (powerLevel / 100) % 10
	
	return hundredsDigit - 5
}

// buildSummedAreaTable creates a summed-area table for efficient range sum queries
func (s *Solution) buildSummedAreaTable() {
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			s.summedArea[x][y] = s.grid[x][y] + 
				s.summedArea[x-1][y] + 
				s.summedArea[x][y-1] - 
				s.summedArea[x-1][y-1]
		}
	}
}

// getSquareSum returns the sum of a square with top-left corner at (x,y) and given size
func (s *Solution) getSquareSum(x, y, size int) int {
	x2 := x + size - 1
	y2 := y + size - 1
	
	return s.summedArea[x2][y2] - 
		s.summedArea[x-1][y2] - 
		s.summedArea[x2][y-1] + 
		s.summedArea[x-1][y-1]
}