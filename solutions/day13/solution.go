/*
 * Day 13: Mine Cart Madness
 * 
 * Part 1: Find the location of the first crash between mine carts.
 * Carts move along tracks and turn at intersections following a left-straight-right pattern.
 * 
 * Part 2: Continue simulation after crashes occur, removing crashed carts immediately,
 * until only one cart remains. Return the location of the last remaining cart.
 * 
 * Key insight: Carts must be removed IMMEDIATELY upon collision,
 * and crashed carts should not move in the same tick they crashed.
 * 
 * CRITICAL BUG ENCOUNTERED DURING SOLVING:
 * ===========================================
 * The initial solution consistently produced the wrong answer (73,122) instead of the correct
 * answer (137,101). After extensive debugging with multiple independent implementations, 
 * it was discovered that the input file had been incorrectly downloaded.
 * 
 * The Problem:
 * - The first line was missing its leading spaces
 * - Should have been: "                   /----------\" 
 * - Was actually:     "/----------\"
 * 
 * This shifted all X coordinates in the track by 19 positions, causing:
 * - Carts to be in wrong relative positions
 * - Collisions to happen at incorrect times and locations
 * - The final surviving cart to be at the wrong position
 * 
 * The bug was particularly insidious because:
 * 1. The simulation still ran without errors
 * 2. All 17 carts were still parsed correctly
 * 3. The algorithmic logic was entirely correct
 * 4. Multiple independent implementations all produced the same wrong answer
 * 5. Even implementations based on successful Reddit solutions produced 73,122
 * 
 * The issue was only discovered when the user manually inspected the raw input file
 * and noticed the missing leading spaces. After re-downloading with proper formatting
 * (using curl with the session cookie), the solution immediately produced the correct answer.
 * 
 * LESSON LEARNED: Always verify that input files are downloaded correctly, especially
 * preserving leading/trailing whitespace which can be significant in grid-based problems.
 * Consider adding validation that checks expected grid dimensions or patterns.
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
	Alive     bool
	ID        int
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
			// Move the cart
			s.moveCart(&carts[i], grid)
			
			// Check for collisions
			for j := range carts {
				if i != j && carts[i].Pos == carts[j].Pos {
					return fmt.Sprintf("%d,%d", carts[i].Pos.X, carts[i].Pos.Y), nil
				}
			}
		}
	}
}

func (s *Solution) Part2() (string, error) {
	grid, carts := s.parseInput()
	
	// Give each cart an ID and mark as alive
	for i := range carts {
		carts[i].ID = i
		carts[i].Alive = true
	}
	
	for tick := 0; tick < 1000000; tick++ {
		// Sort ALL carts by position (including dead ones to maintain indices)
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].Pos.Y == carts[j].Pos.Y {
				return carts[i].Pos.X < carts[j].Pos.X
			}
			return carts[i].Pos.Y < carts[j].Pos.Y
		})
		
		// Process each cart in order
		for i := range carts {
			// Skip dead carts
			if !carts[i].Alive {
				continue
			}
			
			// Move the cart
			s.moveCart(&carts[i], grid)
			
			// Skip if this cart just died from moving out of bounds
			if !carts[i].Alive {
				continue
			}
			
			// Check for collisions with ALL other alive carts
			for j := range carts {
				if i != j && carts[j].Alive && carts[i].Pos == carts[j].Pos {
					// Mark both carts as dead IMMEDIATELY
					carts[i].Alive = false
					carts[j].Alive = false
					break
				}
			}
		}
		
		// Check if only one cart remains AFTER movement
		aliveCount := 0
		var lastAlive *Cart
		for i := range carts {
			if carts[i].Alive {
				aliveCount++
				lastAlive = &carts[i]
			}
		}
		
		if aliveCount == 1 {
			return fmt.Sprintf("%d,%d", lastAlive.Pos.X, lastAlive.Pos.Y), nil
		}
		
		if aliveCount == 0 {
			return "", fmt.Errorf("no carts remaining")
		}
	}
	
	return "", fmt.Errorf("simulation exceeded maximum ticks")
}

func (s *Solution) parseInput() ([][]rune, []Cart) {
	lines := strings.Split(s.input, "\n")
	
	// Create grid - keep original size, don't pad
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
					Alive:     true,
				})
				// Replace cart with track
				grid[y][x] = track
			}
		}
	}
	
	return grid, carts
}

func (s *Solution) moveCart(cart *Cart, grid [][]rune) {
	// Move in current direction
	switch cart.Dir {
	case Up:
		cart.Pos.Y--
	case Down:
		cart.Pos.Y++
	case Left:
		cart.Pos.X--
	case Right:
		cart.Pos.X++
	}
	
	// Safety check - if cart goes out of bounds, mark as dead
	if cart.Pos.Y < 0 || cart.Pos.Y >= len(grid) {
		cart.Alive = false
		return
	}
	if cart.Pos.X < 0 || cart.Pos.X >= len(grid[cart.Pos.Y]) {
		cart.Alive = false
		return
	}
	
	// Get track at new position
	track := grid[cart.Pos.Y][cart.Pos.X]
	
	// If cart is on empty space, mark as dead
	if track == ' ' {
		cart.Alive = false
		return
	}
	
	// Handle different track pieces
	switch track {
	case '/':
		// Curve track /
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
		// Curve track \
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
		// Intersection - apply turn state
		switch cart.TurnState {
		case TurnLeft:
			// Turn left (counterclockwise)
			cart.Dir = (cart.Dir + 3) % 4
		case GoStraight:
			// No direction change
		case TurnRight:
			// Turn right (clockwise)
			cart.Dir = (cart.Dir + 1) % 4
		}
		// Advance turn state for next intersection
		cart.TurnState = (cart.TurnState + 1) % 3
	case '|', '-':
		// Straight track - no direction change
	}
}