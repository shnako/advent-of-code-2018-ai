/*
 * Day 13: Mine Cart Madness - Clean Implementation
 * 
 * Completely fresh approach focusing on:
 * 1. Simple, clear data structures
 * 2. Exact adherence to problem specification
 * 3. Careful collision detection order
 */

package day13

import (
	"fmt"
	"sort"
	"strings"
)

// Simple position struct
type Pos struct {
	X, Y int
}

// Direction constants - using clockwise ordering for easy turning
const (
	NORTH = 0
	EAST  = 1
	SOUTH = 2
	WEST  = 3
)

// Turn states for intersections
const (
	TURN_LEFT = 0
	GO_STRAIGHT = 1
	TURN_RIGHT = 2
)

// Cart structure
type CartClean struct {
	Pos       Pos
	Dir       int
	TurnState int
	ID        int // For debugging
	Alive     bool
}

type SolutionClean struct {
	input string
}

func NewClean(input string) *SolutionClean {
	return &SolutionClean{input: strings.TrimSpace(input)}
}

func (s *SolutionClean) parseInput() ([]string, []CartClean) {
	lines := strings.Split(s.input, "\n")
	var carts []CartClean
	cartID := 0
	
	// Create a copy of lines for grid, replacing carts with appropriate track
	grid := make([]string, len(lines))
	for i, line := range lines {
		runes := []rune(line)
		for j, char := range runes {
			switch char {
			case '^':
				carts = append(carts, CartClean{
					Pos: Pos{X: j, Y: i},
					Dir: NORTH,
					TurnState: TURN_LEFT,
					ID: cartID,
					Alive: true,
				})
				cartID++
				runes[j] = '|' // Replace with vertical track
			case 'v':
				carts = append(carts, CartClean{
					Pos: Pos{X: j, Y: i},
					Dir: SOUTH,
					TurnState: TURN_LEFT,
					ID: cartID,
					Alive: true,
				})
				cartID++
				runes[j] = '|' // Replace with vertical track
			case '<':
				carts = append(carts, CartClean{
					Pos: Pos{X: j, Y: i},
					Dir: WEST,
					TurnState: TURN_LEFT,
					ID: cartID,
					Alive: true,
				})
				cartID++
				runes[j] = '-' // Replace with horizontal track
			case '>':
				carts = append(carts, CartClean{
					Pos: Pos{X: j, Y: i},
					Dir: EAST,
					TurnState: TURN_LEFT,
					ID: cartID,
					Alive: true,
				})
				cartID++
				runes[j] = '-' // Replace with horizontal track
			}
		}
		grid[i] = string(runes)
	}
	
	return grid, carts
}

func (s *SolutionClean) moveCart(cart *CartClean, grid []string) {
	// Calculate new position without moving yet
	newY := cart.Pos.Y
	newX := cart.Pos.X
	
	switch cart.Dir {
	case NORTH:
		newY--
	case SOUTH:
		newY++
	case WEST:
		newX--
	case EAST:
		newX++
	}
	
	// Check bounds before moving
	if newY < 0 || newY >= len(grid) || newX < 0 || newX >= len(grid[newY]) {
		// Cart would go off the track - mark as crashed
		cart.Alive = false
		return
	}
	
	// Move to new position
	cart.Pos.Y = newY
	cart.Pos.X = newX
	
	track := grid[cart.Pos.Y][cart.Pos.X]
	
	// Handle track-based direction changes
	switch track {
	case '/':
		// Forward slash curve
		switch cart.Dir {
		case NORTH: cart.Dir = EAST
		case EAST: cart.Dir = NORTH
		case SOUTH: cart.Dir = WEST
		case WEST: cart.Dir = SOUTH
		}
	case '\\':
		// Backslash curve
		switch cart.Dir {
		case NORTH: cart.Dir = WEST
		case WEST: cart.Dir = NORTH
		case SOUTH: cart.Dir = EAST
		case EAST: cart.Dir = SOUTH
		}
	case '+':
		// Intersection - apply turn based on turn state
		switch cart.TurnState {
		case TURN_LEFT:
			cart.Dir = (cart.Dir + 3) % 4 // Turn left (counterclockwise)
		case GO_STRAIGHT:
			// No change in direction
		case TURN_RIGHT:
			cart.Dir = (cart.Dir + 1) % 4 // Turn right (clockwise)
		}
		// Advance to next turn state
		cart.TurnState = (cart.TurnState + 1) % 3
	}
	// For '|' and '-', no direction change needed
}

func (s *SolutionClean) Part1Clean() (string, error) {
	grid, carts := s.parseInput()
	
	for tick := 0; tick < 100000; tick++ {
		// Sort carts by position (top to bottom, left to right)
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].Pos.Y == carts[j].Pos.Y {
				return carts[i].Pos.X < carts[j].Pos.X
			}
			return carts[i].Pos.Y < carts[j].Pos.Y
		})
		
		// Move each cart and check for collisions
		for i := range carts {
			s.moveCart(&carts[i], grid)
			
			// Check for collision with any other cart
			for j := range carts {
				if i != j && carts[i].Pos == carts[j].Pos {
					return fmt.Sprintf("%d,%d", carts[i].Pos.X, carts[i].Pos.Y), nil
				}
			}
		}
	}
	
	return "", fmt.Errorf("no collision found")
}

func (s *SolutionClean) Part2Clean() (string, error) {
	grid, carts := s.parseInput()
	
	for tick := 0; tick < 100000; tick++ {
		// Get alive carts and sort by position
		aliveCarts := []int{}
		for i := range carts {
			if carts[i].Alive {
				aliveCarts = append(aliveCarts, i)
			}
		}
		
		if len(aliveCarts) == 1 {
			cart := carts[aliveCarts[0]]
			return fmt.Sprintf("%d,%d", cart.Pos.X, cart.Pos.Y), nil
		}
		if len(aliveCarts) == 0 {
			return "", fmt.Errorf("no carts remaining")
		}
		
		// Sort alive carts by position
		sort.Slice(aliveCarts, func(i, j int) bool {
			ci, cj := aliveCarts[i], aliveCarts[j]
			if carts[ci].Pos.Y == carts[cj].Pos.Y {
				return carts[ci].Pos.X < carts[cj].Pos.X
			}
			return carts[ci].Pos.Y < carts[cj].Pos.Y
		})
		
		// Move each cart in order
		for _, cartIdx := range aliveCarts {
			if !carts[cartIdx].Alive {
				continue // Skip if crashed earlier this tick
			}
			
			s.moveCart(&carts[cartIdx], grid)
			if !carts[cartIdx].Alive {
				continue // Skip if went off track
			}
			
			// Check for collisions with all other alive carts
			for j := range carts {
				if j != cartIdx && carts[j].Alive && carts[cartIdx].Pos == carts[j].Pos {
					// Collision! Remove both carts
					carts[cartIdx].Alive = false
					carts[j].Alive = false
					break
				}
			}
		}
	}
	
	return "", fmt.Errorf("simulation timed out")
}