package utils

import "strings"

type Point struct {
	X, Y int
}

type Grid struct {
	Width, Height int
	Data          [][]rune
}

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

func (g *Grid) Get(p Point) rune {
	if p.X < 0 || p.X >= g.Width || p.Y < 0 || p.Y >= g.Height {
		return 0
	}
	return g.Data[p.Y][p.X]
}

func (g *Grid) Set(p Point, val rune) {
	if p.X >= 0 && p.X < g.Width && p.Y >= 0 && p.Y < g.Height {
		g.Data[p.Y][p.X] = val
	}
}

func (g *Grid) InBounds(p Point) bool {
	return p.X >= 0 && p.X < g.Width && p.Y >= 0 && p.Y < g.Height
}

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

func (g *Grid) String() string {
	lines := make([]string, g.Height)
	for i, row := range g.Data {
		lines[i] = string(row)
	}
	return strings.Join(lines, "\n")
}

func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

func (p Point) Sub(other Point) Point {
	return Point{p.X - other.X, p.Y - other.Y}
}

func (p Point) Manhattan(other Point) int {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}

func (p Point) Neighbors4() []Point {
	return []Point{
		p.Add(North),
		p.Add(South),
		p.Add(East),
		p.Add(West),
	}
}

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

// Common directions
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