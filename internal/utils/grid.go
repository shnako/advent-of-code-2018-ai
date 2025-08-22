package utils

import "strings"

// Point represents a 2D coordinate.
type Point struct {
	X, Y int
}

// Grid represents a 2D grid of runes with utility methods.
type Grid struct {
	Width, Height int
	Data          [][]rune
}

// NewGrid creates a new Grid from string input.
func NewGrid(input string) *Grid {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	height := len(lines)
	if height == 0 {
		return &Grid{Width: 0, Height: 0, Data: [][]rune{}}
	}

	width := len(lines[0])
	data := make([][]rune, height)
	for i, line := range lines {
		data[i] = []rune(line)
	}

	return &Grid{
		Width:  width,
		Height: height,
		Data:   data,
	}
}

// Get returns the rune at the given point, or 0 if out of bounds.
func (g *Grid) Get(p Point) rune {
	if p.X < 0 || p.X >= g.Width || p.Y < 0 || p.Y >= g.Height {
		return 0
	}
	return g.Data[p.Y][p.X]
}

// Set sets the rune at the given point if it's within bounds.
func (g *Grid) Set(p Point, val rune) {
	if p.X >= 0 && p.X < g.Width && p.Y >= 0 && p.Y < g.Height {
		g.Data[p.Y][p.X] = val
	}
}

// InBounds returns true if the point is within the grid boundaries.
func (g *Grid) InBounds(p Point) bool {
	return p.X >= 0 && p.X < g.Width && p.Y >= 0 && p.Y < g.Height
}

// Find returns all points containing the target rune.
func (g *Grid) Find(target rune) []Point {
	points := []Point{}
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if g.Data[y][x] == target {
				points = append(points, Point{x, y})
			}
		}
	}
	return points
}

// String returns the grid as a string representation.
func (g *Grid) String() string {
	lines := make([]string, g.Height)
	for i, row := range g.Data {
		lines[i] = string(row)
	}
	return strings.Join(lines, "\n")
}

// Add returns the sum of two points.
func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

// Sub returns the difference of two points.
func (p Point) Sub(other Point) Point {
	return Point{p.X - other.X, p.Y - other.Y}
}

// Manhattan returns the Manhattan distance between two points.
func (p Point) Manhattan(other Point) int {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}

// Neighbors4 returns the four cardinal neighbors.
func (p Point) Neighbors4() []Point {
	return []Point{
		p.Add(North),
		p.Add(South),
		p.Add(East),
		p.Add(West),
	}
}

// Neighbors8 returns all eight neighbors including diagonals.
func (p Point) Neighbors8() []Point {
	return []Point{
		p.Add(North),
		p.Add(South),
		p.Add(East),
		p.Add(West),
		p.Add(NorthEast),
		p.Add(NorthWest),
		p.Add(SouthEast),
		p.Add(SouthWest),
	}
}

// Common direction vectors for grid navigation.
var (
	North     = Point{0, -1}
	South     = Point{0, 1}
	East      = Point{1, 0}
	West      = Point{-1, 0}
	NorthEast = Point{1, -1}
	NorthWest = Point{-1, -1}
	SouthEast = Point{1, 1}
	SouthWest = Point{-1, 1}

	Cardinals = []Point{North, South, East, West}
	AllDirs   = []Point{North, South, East, West, NorthEast, NorthWest, SouthEast, SouthWest}
)
