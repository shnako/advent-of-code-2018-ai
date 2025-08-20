/*
 * Day 10: The Stars Align
 *
 * Part 1: Simulate moving points to find when they form the smallest bounding box
 * and create a readable message. Points move with constant velocity and at some
 * moment align to display text that can be visually read.
 *
 * Part 2: Find the exact time (number of seconds) when the message appears.
 * This requires detecting the moment of minimum bounding box area.
 */

package day10

import (
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x, y   int
	vx, vy int
}

type Solution struct {
	points []Point
}

func New(input string) *Solution {
	input = strings.ReplaceAll(strings.TrimSpace(input), "\r\n", "\n")
	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`position=<\s*(-?\d+),\s*(-?\d+)> velocity=<\s*(-?\d+),\s*(-?\d+)>`)

	var points []Point
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) == 5 {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			vx, _ := strconv.Atoi(matches[3])
			vy, _ := strconv.Atoi(matches[4])
			points = append(points, Point{x, y, vx, vy})
		}
	}

	return &Solution{points: points}
}

func (s *Solution) Part1() (string, error) {
	message, _ := s.findMessage()
	return message, nil
}

func (s *Solution) Part2() (int, error) {
	_, time := s.findMessage()
	return time, nil
}

// findMessage finds when points align to form the minimum bounding box and returns the message and time
func (s *Solution) findMessage() (string, int) {
	points := make([]Point, len(s.points))
	copy(points, s.points)

	minArea := int64(1<<63 - 1)
	bestTime := 0
	var bestPoints []Point

	// Look for the minimum bounding box area
	for time := 0; time < 100000; time++ {
		// Calculate bounding box
		minX, maxX := points[0].x, points[0].x
		minY, maxY := points[0].y, points[0].y

		for _, p := range points {
			if p.x < minX {
				minX = p.x
			}
			if p.x > maxX {
				maxX = p.x
			}
			if p.y < minY {
				minY = p.y
			}
			if p.y > maxY {
				maxY = p.y
			}
		}

		width := maxX - minX + 1
		height := maxY - minY + 1
		area := int64(width) * int64(height)

		if area < minArea {
			minArea = area
			bestTime = time
			bestPoints = make([]Point, len(points))
			copy(bestPoints, points)
		}

		// If area is getting larger again and we have a reasonable message size, we found it
		if time > bestTime+10 && height < 20 && width < 200 {
			break
		}

		// Move points forward one second
		for i := range points {
			points[i].x += points[i].vx
			points[i].y += points[i].vy
		}
	}

	// Generate the visual message from best points
	message := s.visualizePoints(bestPoints)
	return message, bestTime
}

// visualizePoints creates a visual representation of the points
func (s *Solution) visualizePoints(points []Point) string {
	if len(points) == 0 {
		return ""
	}

	// Find bounding box
	minX, maxX := points[0].x, points[0].x
	minY, maxY := points[0].y, points[0].y

	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	width := maxX - minX + 1
	height := maxY - minY + 1

	// Create grid
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	// Mark points
	for _, p := range points {
		x := p.x - minX
		y := p.y - minY
		if x >= 0 && x < width && y >= 0 && y < height {
			grid[y][x] = '#'
		}
	}

	// Convert to string
	var result strings.Builder
	for _, row := range grid {
		result.WriteString(string(row))
		result.WriteString("\n")
	}

	return strings.TrimSpace(result.String())
}
