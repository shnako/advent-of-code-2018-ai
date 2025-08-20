/*
 * Day 13: Mine Cart Madness
 * 
 * Part 1: Track carts moving on a mine track system and find the location of the first crash.
 * Carts move in order (top to bottom, left to right) and follow specific turning rules at intersections.
 * At intersections, carts alternate between turning left, going straight, and turning right.
 * 
 * Part 2: Continue simulation after crashes occur, removing crashed carts, until only one cart remains.
 * Return the location of the last remaining cart.
 */

package day13

import (
	"fmt"
	"sort"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type TurnState int

const (
	TurnLeft TurnState = iota
	GoStraight
	TurnRight
)

type Point struct {
	X, Y int
}

type Cart struct {
	Pos       Point
	Dir       Direction
	TurnState TurnState
	Crashed   bool
}

type Solution struct {
	input string
}

func New(input string) *Solution {
	return &Solution{input: strings.TrimSpace(input)}
}

func (s *Solution) Part1() (string, error) {
	grid, carts := s.parseInput()
	
	for {
		// Sort carts by position (top to bottom, left to right)
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].Pos.Y == carts[j].Pos.Y {
				return carts[i].Pos.X < carts[j].Pos.X
			}
			return carts[i].Pos.Y < carts[j].Pos.Y
		})
		
		for i := range carts {
			if carts[i].Crashed {
				continue
			}
			
			// Move the cart
			s.moveCart(&carts[i], grid)
			
			// Check for collisions
			for j := range carts {
				if i != j && !carts[j].Crashed && carts[i].Pos == carts[j].Pos {
					return fmt.Sprintf("%d,%d", carts[i].Pos.X, carts[i].Pos.Y), nil
				}
			}
		}
	}
}

func (s *Solution) Part2() (string, error) {
	grid, carts := s.parseInput()
	
	for tick := 0; tick < 1000000; tick++ { // Add safety limit
		// Sort carts by position (top to bottom, left to right)
		// Only consider non-crashed carts for the sort
		activeCarts := make([]int, 0)
		for i := range carts {
			if !carts[i].Crashed {
				activeCarts = append(activeCarts, i)
			}
		}
		
		sort.Slice(activeCarts, func(i, j int) bool {
			ci, cj := activeCarts[i], activeCarts[j]
			if carts[ci].Pos.Y == carts[cj].Pos.Y {
				return carts[ci].Pos.X < carts[cj].Pos.X
			}
			return carts[ci].Pos.Y < carts[cj].Pos.Y
		})
		
		for _, cartIdx := range activeCarts {
			if carts[cartIdx].Crashed {
				continue // Skip if this cart crashed earlier in this tick
			}
			
			// Move the cart
			s.moveCart(&carts[cartIdx], grid)
			if carts[cartIdx].Crashed {
				continue // Cart went out of bounds
			}
			
			// Check for collisions with all other carts (including ones that haven't moved yet this tick)
			for j := range carts {
				if j != cartIdx && !carts[j].Crashed && carts[cartIdx].Pos == carts[j].Pos {
					carts[cartIdx].Crashed = true
					carts[j].Crashed = true
					break
				}
			}
		}
		
		// Count remaining active carts
		active := 0
		var lastCart *Cart
		for i := range carts {
			if !carts[i].Crashed {
				active++
				lastCart = &carts[i]
			}
		}
		
		if active == 1 {
			return fmt.Sprintf("%d,%d", lastCart.Pos.X, lastCart.Pos.Y), nil
		}
		if active == 0 {
			return "", fmt.Errorf("no carts remaining")
		}
	}
	
	return "", fmt.Errorf("simulation exceeded maximum ticks")
}

func (s *Solution) parseInput() ([][]rune, []Cart) {
	lines := strings.Split(s.input, "\n")
	grid := make([][]rune, len(lines))
	var carts []Cart
	
	for y, line := range lines {
		grid[y] = []rune(line)
		for x, char := range line {
			var dir Direction
			var isCart bool
			var track rune
			
			switch char {
			case '^':
				dir = Up
				isCart = true
				track = '|'
			case 'v':
				dir = Down
				isCart = true
				track = '|'
			case '<':
				dir = Left
				isCart = true
				track = '-'
			case '>':
				dir = Right
				isCart = true
				track = '-'
			}
			
			if isCart {
				carts = append(carts, Cart{
					Pos:       Point{X: x, Y: y},
					Dir:       dir,
					TurnState: TurnLeft,
				})
				grid[y][x] = track
			}
		}
	}
	
	return grid, carts
}

func (s *Solution) moveCart(cart *Cart, grid [][]rune) {
	// Calculate potential new position
	newY := cart.Pos.Y
	newX := cart.Pos.X
	switch cart.Dir {
	case Up:
		newY--
	case Down:
		newY++
	case Left:
		newX--
	case Right:
		newX++
	}
	
	// Check bounds before moving  
	if newY < 0 || newY >= len(grid) {
		cart.Crashed = true
		return
	}
	if newX < 0 || newX >= len(grid[newY]) {
		cart.Crashed = true
		return
	}
	
	// Move to new position
	cart.Pos.Y = newY
	cart.Pos.X = newX
	
	// Handle turns based on track piece
	track := grid[cart.Pos.Y][cart.Pos.X]
	
	switch track {
	case '/':
		switch cart.Dir {
		case Up:
			cart.Dir = Right
		case Down:
			cart.Dir = Left
		case Left:
			cart.Dir = Down
		case Right:
			cart.Dir = Up
		}
	case '\\':
		switch cart.Dir {
		case Up:
			cart.Dir = Left
		case Down:
			cart.Dir = Right
		case Left:
			cart.Dir = Up
		case Right:
			cart.Dir = Down
		}
	case '+':
		// Intersection - turn based on turn state
		switch cart.TurnState {
		case TurnLeft:
			cart.Dir = (cart.Dir + 3) % 4 // Turn left
		case GoStraight:
			// No change
		case TurnRight:
			cart.Dir = (cart.Dir + 1) % 4 // Turn right
		}
		cart.TurnState = (cart.TurnState + 1) % 3
	}
}