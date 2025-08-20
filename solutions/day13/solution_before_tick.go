package day13

import (
	"fmt"
	"sort"
	"strings"
)

// BeforeTickSolution checks position before final movement
type BeforeTickSolution struct {
	input string
}

func NewBeforeTick(input string) *BeforeTickSolution {
	return &BeforeTickSolution{input: input}
}

func (s *BeforeTickSolution) Part2() (string, error) {
	lines := strings.Split(s.input, "\n")
	
	// Build grid and find carts
	grid := make([][]rune, len(lines))
	carts := []*FreshCart{}
	cartID := 0
	
	for y, line := range lines {
		grid[y] = []rune(line)
		for x, ch := range grid[y] {
			var cart *FreshCart
			switch ch {
			case '^':
				cart = &FreshCart{x: x, y: y, dx: 0, dy: -1, id: cartID}
				grid[y][x] = '|'
			case 'v':
				cart = &FreshCart{x: x, y: y, dx: 0, dy: 1, id: cartID}
				grid[y][x] = '|'
			case '<':
				cart = &FreshCart{x: x, y: y, dx: -1, dy: 0, id: cartID}
				grid[y][x] = '-'
			case '>':
				cart = &FreshCart{x: x, y: y, dx: 1, dy: 0, id: cartID}
				grid[y][x] = '-'
			}
			if cart != nil {
				carts = append(carts, cart)
				cartID++
			}
		}
	}
	
	fmt.Printf("Found %d carts\n", len(carts))
	
	// Run simulation
	for tick := 0; tick < 100000; tick++ {
		// Sort carts by position (Y first, then X)
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y != carts[j].y {
				return carts[i].y < carts[j].y
			}
			return carts[i].x < carts[j].x
		})
		
		// Count carts BEFORE moving
		activeCarts := 0
		var lastActiveCart *FreshCart
		for _, cart := range carts {
			if !cart.crashed {
				activeCarts++
				lastActiveCart = cart
			}
		}
		
		// Check if only one cart remains BEFORE this tick's movements
		if activeCarts == 1 {
			fmt.Printf("Only one cart remains at START of tick %d: Cart %d at (%d,%d)\n", 
				tick, lastActiveCart.id, lastActiveCart.x, lastActiveCart.y)
			// Return position BEFORE movement
			return fmt.Sprintf("%d,%d", lastActiveCart.x, lastActiveCart.y), nil
		}
		
		// Move each cart
		for _, cart := range carts {
			if cart.crashed {
				continue
			}
			
			// Move the cart
			cart.x += cart.dx
			cart.y += cart.dy
			
			// Check bounds
			if cart.y < 0 || cart.y >= len(grid) || cart.x < 0 || cart.x >= len(grid[cart.y]) {
				fmt.Printf("Tick %d: Cart %d out of bounds at (%d,%d)\n", tick, cart.id, cart.x, cart.y)
				cart.crashed = true
				continue
			}
			
			// Check for collisions with other carts
			for _, other := range carts {
				if other.id != cart.id && !other.crashed && 
				   other.x == cart.x && other.y == cart.y {
					fmt.Printf("Tick %d: Collision at (%d,%d) between carts %d and %d\n", 
						tick, cart.x, cart.y, cart.id, other.id)
					cart.crashed = true
					other.crashed = true
				}
			}
			
			if cart.crashed {
				continue
			}
			
			// Update direction based on track
			track := grid[cart.y][cart.x]
			switch track {
			case '/':
				// Swap and negate
				cart.dx, cart.dy = -cart.dy, -cart.dx
			case '\\':
				// Just swap
				cart.dx, cart.dy = cart.dy, cart.dx
			case '+':
				// Handle intersection
				turn := cart.turnCount % 3
				cart.turnCount++
				
				if turn == 0 {
					// Turn left: (dx,dy) -> (dy,-dx)
					cart.dx, cart.dy = cart.dy, -cart.dx
				} else if turn == 2 {
					// Turn right: (dx,dy) -> (-dy,dx)
					cart.dx, cart.dy = -cart.dy, cart.dx
				}
				// turn == 1 means go straight (no change)
			}
		}
		
		// Debug output for specific ticks
		if tick >= 10002 && tick <= 10005 {
			fmt.Printf("After tick %d movements:\n", tick)
			for _, cart := range carts {
				if !cart.crashed {
					fmt.Printf("  Cart %d at (%d,%d) heading (%d,%d)\n", 
						cart.id, cart.x, cart.y, cart.dx, cart.dy)
				}
			}
		}
		
		// Check if all carts crashed
		allCrashed := true
		for _, cart := range carts {
			if !cart.crashed {
				allCrashed = false
				break
			}
		}
		
		if allCrashed {
			return "", fmt.Errorf("all carts crashed")
		}
	}
	
	return "", fmt.Errorf("timeout")
}