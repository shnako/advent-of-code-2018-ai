package day13

import (
	"fmt"
	"sort"
	"strings"
)

// FreshCart represents a cart with clean implementation
type FreshCart struct {
	x, y      int
	dx, dy    int
	turnCount int
	id        int
	crashed   bool
}

// FreshSolution is a completely fresh implementation
type FreshSolution struct {
	input string
}

func NewFresh(input string) *FreshSolution {
	return &FreshSolution{input: input}
}

func (s *FreshSolution) Part2() (string, error) {
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
		
		// Count remaining carts
		remaining := []*FreshCart{}
		for _, cart := range carts {
			if !cart.crashed {
				remaining = append(remaining, cart)
			}
		}
		
		if len(remaining) == 1 {
			fmt.Printf("Final cart at tick %d: (%d,%d)\n", tick, remaining[0].x, remaining[0].y)
			return fmt.Sprintf("%d,%d", remaining[0].x, remaining[0].y), nil
		}
		
		if len(remaining) == 0 {
			return "", fmt.Errorf("all carts crashed")
		}
		
		// Debug output for specific ticks
		if tick >= 10003 && tick <= 10005 {
			fmt.Printf("Tick %d: %d carts remaining\n", tick, len(remaining))
			for _, cart := range remaining {
				fmt.Printf("  Cart %d at (%d,%d) heading (%d,%d)\n", 
					cart.id, cart.x, cart.y, cart.dx, cart.dy)
			}
		}
	}
	
	return "", fmt.Errorf("timeout")
}

// Part2Alternative tries a different approach - checking position AFTER the tick completes
func (s *FreshSolution) Part2Alternative() (string, error) {
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
		
		// Track positions after each move for immediate collision detection
		positions := make(map[string]*FreshCart)
		for _, cart := range carts {
			if !cart.crashed {
				key := fmt.Sprintf("%d,%d", cart.x, cart.y)
				positions[key] = cart
			}
		}
		
		// Move each cart
		for _, cart := range carts {
			if cart.crashed {
				continue
			}
			
			// Remove from old position
			oldKey := fmt.Sprintf("%d,%d", cart.x, cart.y)
			delete(positions, oldKey)
			
			// Move the cart
			cart.x += cart.dx
			cart.y += cart.dy
			
			// Check bounds
			if cart.y < 0 || cart.y >= len(grid) || cart.x < 0 || cart.x >= len(grid[cart.y]) {
				fmt.Printf("Tick %d: Cart %d out of bounds at (%d,%d)\n", tick, cart.id, cart.x, cart.y)
				cart.crashed = true
				continue
			}
			
			// Check for collision at new position
			newKey := fmt.Sprintf("%d,%d", cart.x, cart.y)
			if other, exists := positions[newKey]; exists {
				fmt.Printf("Tick %d: Collision at (%d,%d) between carts %d and %d\n", 
					tick, cart.x, cart.y, cart.id, other.id)
				cart.crashed = true
				other.crashed = true
				delete(positions, newKey)
				continue
			}
			
			// Add to new position if not crashed
			if !cart.crashed {
				positions[newKey] = cart
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
		
		// Count remaining carts
		remaining := []*FreshCart{}
		for _, cart := range carts {
			if !cart.crashed {
				remaining = append(remaining, cart)
			}
		}
		
		if len(remaining) == 1 {
			fmt.Printf("Final cart at END of tick %d: (%d,%d)\n", tick, remaining[0].x, remaining[0].y)
			return fmt.Sprintf("%d,%d", remaining[0].x, remaining[0].y), nil
		}
		
		if len(remaining) == 0 {
			return "", fmt.Errorf("all carts crashed")
		}
		
		// Debug output for specific ticks
		if tick >= 10003 && tick <= 10005 {
			fmt.Printf("End of tick %d: %d carts remaining\n", tick, len(remaining))
			for _, cart := range remaining {
				fmt.Printf("  Cart %d at (%d,%d) heading (%d,%d)\n", 
					cart.id, cart.x, cart.y, cart.dx, cart.dy)
			}
		}
	}
	
	return "", fmt.Errorf("timeout")
}