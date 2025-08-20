/*
 * Day 13: Mine Cart Madness - Opus Solution
 * 
 * Very careful implementation paying attention to:
 * - Exact cart movement order
 * - Immediate collision removal
 * - Proper intersection cycling
 */

package day13

import (
	"fmt"
	"sort"
	"strings"
)

type OpusCart struct {
	X, Y      int
	DX, DY    int // Direction as delta
	Turns     int // Number of turns taken at intersections
	Crashed   bool
}

type OpusSolution struct {
	input string
}

func NewOpus(input string) *OpusSolution {
	return &OpusSolution{input: input}
}

func (s *OpusSolution) parse() ([][]byte, []*OpusCart) {
	lines := strings.Split(s.input, "\n")
	grid := make([][]byte, len(lines))
	var carts []*OpusCart
	
	for y := 0; y < len(lines); y++ {
		grid[y] = []byte(lines[y])
		for x := 0; x < len(grid[y]); x++ {
			var cart *OpusCart
			switch grid[y][x] {
			case '^':
				cart = &OpusCart{X: x, Y: y, DX: 0, DY: -1, Turns: 0}
				grid[y][x] = '|'
			case 'v':
				cart = &OpusCart{X: x, Y: y, DX: 0, DY: 1, Turns: 0}
				grid[y][x] = '|'
			case '<':
				cart = &OpusCart{X: x, Y: y, DX: -1, DY: 0, Turns: 0}
				grid[y][x] = '-'
			case '>':
				cart = &OpusCart{X: x, Y: y, DX: 1, DY: 0, Turns: 0}
				grid[y][x] = '-'
			}
			if cart != nil {
				carts = append(carts, cart)
			}
		}
	}
	
	return grid, carts
}

func (s *OpusSolution) tick(grid [][]byte, carts []*OpusCart) (int, int, bool) {
	// Sort carts by Y then X
	sort.Slice(carts, func(i, j int) bool {
		if carts[i].Y != carts[j].Y {
			return carts[i].Y < carts[j].Y
		}
		return carts[i].X < carts[j].X
	})
	
	for i := 0; i < len(carts); i++ {
		cart := carts[i]
		if cart.Crashed {
			continue
		}
		
		// Move the cart
		cart.X += cart.DX
		cart.Y += cart.DY
		
		// Check for immediate collision
		for j := 0; j < len(carts); j++ {
			if i != j && !carts[j].Crashed && carts[j].X == cart.X && carts[j].Y == cart.Y {
				// Collision! Return the position
				return cart.X, cart.Y, true
			}
		}
		
		// Handle track piece
		if cart.Y >= 0 && cart.Y < len(grid) && cart.X >= 0 && cart.X < len(grid[cart.Y]) {
			track := grid[cart.Y][cart.X]
			switch track {
			case '/':
				// Swap and negate: (dx,dy) -> (-dy,-dx)
				cart.DX, cart.DY = -cart.DY, -cart.DX
			case '\\':
				// Swap: (dx,dy) -> (dy,dx)
				cart.DX, cart.DY = cart.DY, cart.DX
			case '+':
				// Intersection - turn based on count
				turn := cart.Turns % 3
				cart.Turns++
				if turn == 0 {
					// Turn left: (dx,dy) -> (dy,-dx)
					cart.DX, cart.DY = cart.DY, -cart.DX
				} else if turn == 2 {
					// Turn right: (dx,dy) -> (-dy,dx)
					cart.DX, cart.DY = -cart.DY, cart.DX
				}
				// turn == 1 means go straight, no change
			}
		}
	}
	
	return 0, 0, false
}

func (s *OpusSolution) tick2(grid [][]byte, carts []*OpusCart) {
	// Sort carts by Y then X
	sort.Slice(carts, func(i, j int) bool {
		if carts[i].Y != carts[j].Y {
			return carts[i].Y < carts[j].Y
		}
		return carts[i].X < carts[j].X
	})
	
	for i := 0; i < len(carts); i++ {
		cart := carts[i]
		if cart.Crashed {
			continue
		}
		
		// Check if movement would go out of bounds
		newX := cart.X + cart.DX
		newY := cart.Y + cart.DY
		
		if newY < 0 || newY >= len(grid) || newX < 0 || newX >= len(grid[newY]) {
			// Don't move if it would go out of bounds - this cart is stuck
			cart.Crashed = true
			continue
		}
		
		// Move the cart
		cart.X = newX
		cart.Y = newY
		
		// Check for collision with any other non-crashed cart
		for j := 0; j < len(carts); j++ {
			if i != j && !carts[j].Crashed && carts[j].X == cart.X && carts[j].Y == cart.Y {
				// Collision! Mark both as crashed
				cart.Crashed = true
				carts[j].Crashed = true
				break
			}
		}
		
		if cart.Crashed {
			continue
		}
		
		// Handle track piece
		if cart.Y >= 0 && cart.Y < len(grid) && cart.X >= 0 && cart.X < len(grid[cart.Y]) {
			track := grid[cart.Y][cart.X]
			switch track {
			case '/':
				// Swap and negate: (dx,dy) -> (-dy,-dx)
				cart.DX, cart.DY = -cart.DY, -cart.DX
			case '\\':
				// Swap: (dx,dy) -> (dy,dx)
				cart.DX, cart.DY = cart.DY, cart.DX
			case '+':
				// Intersection - turn based on count
				turn := cart.Turns % 3
				cart.Turns++
				if turn == 0 {
					// Turn left: (dx,dy) -> (dy,-dx)
					cart.DX, cart.DY = cart.DY, -cart.DX
				} else if turn == 2 {
					// Turn right: (dx,dy) -> (-dy,dx)
					cart.DX, cart.DY = -cart.DY, cart.DX
				}
				// turn == 1 means go straight, no change
			}
		}
	}
}

func (s *OpusSolution) Part1() (string, error) {
	grid, carts := s.parse()
	
	for i := 0; i < 100000; i++ {
		x, y, collision := s.tick(grid, carts)
		if collision {
			return fmt.Sprintf("%d,%d", x, y), nil
		}
	}
	
	return "", fmt.Errorf("no collision found")
}

func (s *OpusSolution) Part2() (string, error) {
	grid, carts := s.parse()
	
	for i := 0; i < 100000; i++ {
		// Count non-crashed carts BEFORE the tick
		aliveBefore := 0
		for _, cart := range carts {
			if !cart.Crashed {
				aliveBefore++
			}
		}
		
		s.tick2(grid, carts)
		
		// Count non-crashed carts AFTER the tick
		alive := 0
		var lastCart *OpusCart
		for _, cart := range carts {
			if !cart.Crashed {
				alive++
				lastCart = cart
			}
		}
		
		if i % 1000 == 0 && i > 0 {
			fmt.Printf("Tick %d: %d carts alive (was %d)\n", i, alive, aliveBefore)
		}
		
		if alive == 1 {
			// The problem asks for position at END of tick where only one remains
			return fmt.Sprintf("%d,%d", lastCart.X, lastCart.Y), nil
		}
		if alive == 0 {
			return "", fmt.Errorf("all carts crashed")
		}
	}
	
	return "", fmt.Errorf("timeout")
}