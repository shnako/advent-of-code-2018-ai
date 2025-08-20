/*
 * Day 13: Mine Cart Madness - Fresh Implementation
 * 
 * Starting completely from scratch to identify and fix the bug.
 */

package day13

import (
	"fmt"
	"sort"
	"strings"
)

type Direction2 int

const (
	North2 Direction2 = iota
	East2
	South2 
	West2
)

type Turn2 int

const (
	Left2 Turn2 = iota
	Straight2
	Right2
)

type Cart2 struct {
	X, Y      int
	Direction Direction2
	NextTurn  Turn2
	Alive     bool
}

type Grid [][]rune

func (g Grid) Get(x, y int) rune {
	if y < 0 || y >= len(g) || x < 0 || x >= len(g[y]) {
		return ' ' // Out of bounds
	}
	return g[y][x]
}

func NewSolutionFromScratch(input string) *SolutionFromScratch {
	return &SolutionFromScratch{input: strings.TrimSpace(input)}
}

type SolutionFromScratch struct {
	input string
}

func (s *SolutionFromScratch) ParseGrid() (Grid, []Cart2) {
	lines := strings.Split(s.input, "\n")
	grid := make(Grid, len(lines))
	var carts []Cart2
	
	for y, line := range lines {
		grid[y] = []rune(line)
		for x, char := range line {
			switch char {
			case '^':
				carts = append(carts, Cart2{X: x, Y: y, Direction: North2, NextTurn: Left2, Alive: true})
				grid[y][x] = '|' // Replace with vertical track
			case 'v':
				carts = append(carts, Cart2{X: x, Y: y, Direction: South2, NextTurn: Left2, Alive: true})
				grid[y][x] = '|' // Replace with vertical track
			case '<':
				carts = append(carts, Cart2{X: x, Y: y, Direction: West2, NextTurn: Left2, Alive: true})
				grid[y][x] = '-' // Replace with horizontal track
			case '>':
				carts = append(carts, Cart2{X: x, Y: y, Direction: East2, NextTurn: Left2, Alive: true})
				grid[y][x] = '-' // Replace with horizontal track
			}
		}
	}
	
	return grid, carts
}

func (s *SolutionFromScratch) isValidConnection(fromX, fromY int, fromTrack rune, direction Direction2, toX, toY int, toTrack rune, grid Grid) bool {
	// Check if movement from one track piece to another is valid
	switch fromTrack {
	case '|':
		// Vertical track connects north and south
		if direction == North2 {
			return toTrack == '|' || toTrack == '+' || toTrack == '/' || toTrack == '\\'
		} else if direction == South2 {
			return toTrack == '|' || toTrack == '+' || toTrack == '/' || toTrack == '\\'
		}
		return false
	case '-':
		// Horizontal track connects east and west  
		if direction == East2 {
			return toTrack == '-' || toTrack == '+' || toTrack == '/' || toTrack == '\\'
		} else if direction == West2 {
			return toTrack == '-' || toTrack == '+' || toTrack == '/' || toTrack == '\\'
		}
		return false
	case '+':
		// Intersection connects all directions
		return toTrack == '|' || toTrack == '-' || toTrack == '+' || toTrack == '/' || toTrack == '\\'
	case '/', '\\':
		// Curves connect based on their specific rules
		return toTrack == '|' || toTrack == '-' || toTrack == '+' || toTrack == '/' || toTrack == '\\'
	}
	return false
}

func (s *SolutionFromScratch) moveCart(cart *Cart2, grid Grid) {
	oldX, oldY := cart.X, cart.Y
	oldTrack := grid.Get(oldX, oldY)
	
	// Add debugging for carts near the problem area
	if cart.X > 125 || cart.Y < 10 {
		fmt.Printf("DEBUG: Cart at (%d,%d) dir %d on track '%c' about to move\n", 
			cart.X, cart.Y, cart.Direction, oldTrack)
	}
	
	// Special tracking for the problematic cart path
	if cart.X == 127 && cart.Y >= 0 && cart.Y <= 10 {
		fmt.Printf("TRACE: Cart at (127,%d) dir %d on track '%c'\n", cart.Y, cart.Direction, oldTrack)
	}
	
	// Calculate new position
	newX, newY := cart.X, cart.Y
	switch cart.Direction {
	case North2:
		newY--
	case South2:
		newY++
	case West2:
		newX--
	case East2:
		newX++
	}
	
	// Check bounds
	newTrack := grid.Get(newX, newY)
	if newTrack == ' ' {
		fmt.Printf("PANIC: Cart trying to move out of bounds from (%d,%d) to (%d,%d)\n", 
			oldX, oldY, newX, newY)
		panic("Cart went out of bounds!")
	}
	
	// Check connectivity - but allow for debugging
	if !s.isValidConnection(oldX, oldY, oldTrack, cart.Direction, newX, newY, newTrack, grid) {
		fmt.Printf("WARNING: Invalid connection: Cart at (%d,%d) track '%c' dir %d trying to move to (%d,%d) track '%c'\n", 
			oldX, oldY, oldTrack, cart.Direction, newX, newY, newTrack)
		// For now, continue to allow debugging - this will crash later when going out of bounds
	}
	
	// Movement is valid, update cart position
	cart.X, cart.Y = newX, newY
	track := newTrack
	
	// Handle track-specific behavior
	switch track {
	case '/':
		// Forward slash curve
		switch cart.Direction {
		case North2:
			cart.Direction = East2
		case South2:
			cart.Direction = West2
		case West2:
			cart.Direction = South2
		case East2:
			cart.Direction = North2
		}
	case '\\':
		// Backslash curve  
		switch cart.Direction {
		case North2:
			cart.Direction = West2
		case South2:
			cart.Direction = East2
		case West2:
			cart.Direction = North2
		case East2:
			cart.Direction = South2
		}
	case '+':
		// Intersection - turn based on next turn state
		switch cart.NextTurn {
		case Left2:
			cart.Direction = Direction2((int(cart.Direction) + 3) % 4) // Turn left
		case Straight2:
			// No change in direction
		case Right2:
			cart.Direction = Direction2((int(cart.Direction) + 1) % 4) // Turn right
		}
		// Cycle to next turn state
		cart.NextTurn = Turn2((int(cart.NextTurn) + 1) % 3)
	}
	// For '-' and '|' tracks, no direction change needed
}

func (s *SolutionFromScratch) Part1Debug() (string, error) {
	grid, carts := s.ParseGrid()
	
	fmt.Printf("Initial setup: %d carts\n", len(carts))
	for i, cart := range carts {
		fmt.Printf("Cart %d: (%d,%d) dir %d on track '%c'\n", 
			i, cart.X, cart.Y, cart.Direction, grid.Get(cart.X, cart.Y))
	}
	
	for tick := 0; tick < 100000; tick++ {
		// Sort carts by position (top to bottom, left to right)
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].Y == carts[j].Y {
				return carts[i].X < carts[j].X
			}
			return carts[i].Y < carts[j].Y
		})
		
		for i := range carts {
			// Move cart
			oldX, oldY := carts[i].X, carts[i].Y
			s.moveCart(&carts[i], grid)
			
			if tick < 5 {
				fmt.Printf("Tick %d: Cart %d moved from (%d,%d) to (%d,%d)\n", 
					tick, i, oldX, oldY, carts[i].X, carts[i].Y)
			}
			
			// Check for collision
			for j := range carts {
				if i != j && carts[i].X == carts[j].X && carts[i].Y == carts[j].Y {
					return fmt.Sprintf("%d,%d", carts[i].X, carts[i].Y), nil
				}
			}
		}
	}
	
	return "", fmt.Errorf("no collision found")
}

func (s *SolutionFromScratch) Part2Debug() (string, error) {
	grid, carts := s.ParseGrid()
	
	for tick := 0; tick < 100000; tick++ {
		// Get alive carts and sort by position
		var aliveCarts []int
		for i := range carts {
			if carts[i].Alive {
				aliveCarts = append(aliveCarts, i)
			}
		}
		
		sort.Slice(aliveCarts, func(i, j int) bool {
			ci, cj := aliveCarts[i], aliveCarts[j]
			if carts[ci].Y == carts[cj].Y {
				return carts[ci].X < carts[cj].X
			}
			return carts[ci].Y < carts[cj].Y
		})
		
		// Move each cart and check for collisions
		for _, cartIdx := range aliveCarts {
			if !carts[cartIdx].Alive {
				continue // Skip if crashed earlier this tick
			}
			
			s.moveCart(&carts[cartIdx], grid)
			
			// Check for collision with any other alive cart
			for j := range carts {
				if j != cartIdx && carts[j].Alive && 
				   carts[cartIdx].X == carts[j].X && carts[cartIdx].Y == carts[j].Y {
					carts[cartIdx].Alive = false
					carts[j].Alive = false
					break
				}
			}
		}
		
		// Count remaining carts
		aliveCount := 0
		var lastCart *Cart2
		for i := range carts {
			if carts[i].Alive {
				aliveCount++
				lastCart = &carts[i]
			}
		}
		
		if aliveCount == 1 {
			return fmt.Sprintf("%d,%d", lastCart.X, lastCart.Y), nil
		}
		if aliveCount == 0 {
			return "", fmt.Errorf("no carts remaining")
		}
	}
	
	return "", fmt.Errorf("simulation timed out")
}