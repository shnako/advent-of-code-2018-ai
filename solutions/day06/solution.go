/*
 * Day 6: Chronal Coordinates
 *
 * Part 1: Find the largest finite area in a coordinate grid using Manhattan distance.
 * For each point in the grid, determine which input coordinate is closest (using Manhattan
 * distance). Count the area for each coordinate, but exclude coordinates with infinite
 * areas (those that extend to the edge of the grid).
 *
 * Part 2: Find the size of the region containing all locations which have a total distance
 * to all given coordinates of less than 10000. For each point, calculate the sum of Manhattan
 * distances to all coordinates and count how many locations have a total distance < 10000.
 */

package day06

import (
	"strconv"
	"strings"

	"github.com/shnako/advent-of-code-2018-ai/internal/utils"
)


type Solution struct {
	coordinates []utils.Point
}

func New(input string) *Solution {
	// Handle both Unix (\n) and Windows (\r\n) line endings
	input = strings.ReplaceAll(strings.TrimSpace(input), "\r\n", "\n")
	lines := strings.Split(input, "\n")
	coordinates := make([]utils.Point, 0, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, ", ")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			continue // Skip invalid coordinate lines
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			continue // Skip invalid coordinate lines
		}
		coordinates = append(coordinates, utils.Point{X: x, Y: y})
	}

	return &Solution{coordinates: coordinates}
}

func (s *Solution) Part1() (int, error) {
	// Find the bounding box
	minX, maxX := s.coordinates[0].X, s.coordinates[0].X
	minY, maxY := s.coordinates[0].Y, s.coordinates[0].Y

	for _, coord := range s.coordinates {
		if coord.X < minX {
			minX = coord.X
		}
		if coord.X > maxX {
			maxX = coord.X
		}
		if coord.Y < minY {
			minY = coord.Y
		}
		if coord.Y > maxY {
			maxY = coord.Y
		}
	}

	// Track which coordinates have infinite areas
	infiniteAreas := make(map[int]bool)

	// Count areas for each coordinate
	areas := make(map[int]int)

	// Check each point in the grid
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			closest := s.findClosestCoordinate(utils.Point{X: x, Y: y})
			if closest != -1 {
				areas[closest]++

				// If this point is on the edge, mark the coordinate as having infinite area
				if x == minX || x == maxX || y == minY || y == maxY {
					infiniteAreas[closest] = true
				}
			}
		}
	}

	// Find the largest finite area
	maxArea := 0
	for coordIndex, area := range areas {
		if !infiniteAreas[coordIndex] && area > maxArea {
			maxArea = area
		}
	}

	return maxArea, nil
}

func (s *Solution) Part2() (int, error) {
	return s.countRegionSize(10000), nil
}

// countRegionSize counts locations where the sum of Manhattan distances to all coordinates is less than maxDistance
func (s *Solution) countRegionSize(maxDistance int) int {
	// Find the bounding box (with some padding to ensure we capture the full region)
	minX, maxX := s.coordinates[0].X, s.coordinates[0].X
	minY, maxY := s.coordinates[0].Y, s.coordinates[0].Y

	for _, coord := range s.coordinates {
		if coord.X < minX {
			minX = coord.X
		}
		if coord.X > maxX {
			maxX = coord.X
		}
		if coord.Y < minY {
			minY = coord.Y
		}
		if coord.Y > maxY {
			maxY = coord.Y
		}
	}

	// Add padding to ensure we don't miss edge cases
	padding := maxDistance / len(s.coordinates)
	minX -= padding
	maxX += padding
	minY -= padding
	maxY += padding

	count := 0

	// Check each point in the expanded grid
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			totalDistance := 0
			point := utils.Point{X: x, Y: y}

			// Calculate sum of distances to all coordinates
			for _, coord := range s.coordinates {
				totalDistance += manhattanDistance(point, coord)
			}

			if totalDistance < maxDistance {
				count++
			}
		}
	}

	return count
}

// findClosestCoordinate returns the index of the closest coordinate to the given point
// Returns -1 if there's a tie (multiple coordinates at the same minimum distance)
func (s *Solution) findClosestCoordinate(point utils.Point) int {
	minDistance := -1
	closestIndex := -1
	tied := false

	for i, coord := range s.coordinates {
		distance := manhattanDistance(point, coord)

		if minDistance == -1 || distance < minDistance {
			minDistance = distance
			closestIndex = i
			tied = false
		} else if distance == minDistance {
			tied = true
		}
	}

	if tied {
		return -1
	}

	return closestIndex
}

// manhattanDistance calculates the Manhattan distance between two points
func manhattanDistance(a, b utils.Point) int {
	return utils.Abs(a.X-b.X) + utils.Abs(a.Y-b.Y)
}

