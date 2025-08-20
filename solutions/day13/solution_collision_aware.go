package day13

import (
	"fmt"
	"sort"
	"strings"
)

// CollisionAwareSolution handles collisions differently
type CollisionAwareSolution struct {
	input string
}

func NewCollisionAware(input string) *CollisionAwareSolution {
	return &CollisionAwareSolution{input: input}
}

func (s *CollisionAwareSolution) Part2() (string, error) {
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
		// Sort carts by position (Y first, then X) at the START of the tick
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y != carts[j].y {
				return carts[i].y < carts[j].y
			}
			return carts[i].x < carts[j].x
		})
		
		// Track which carts moved this tick
		movedThisTick := make(map[int]bool)
		
		// Move each cart in the sorted order
		for i := 0; i < len(carts); i++ {
			cart := carts[i]
			if cart.crashed {
				continue
			}
			
			// Save old position for debugging
			oldX, oldY := cart.x, cart.y
			
			// Move the cart
			cart.x += cart.dx
			cart.y += cart.dy
			movedThisTick[cart.id] = true
			
			// Check bounds
			if cart.y < 0 || cart.y >= len(grid) || cart.x < 0 || cart.x >= len(grid[cart.y]) {
				fmt.Printf("Tick %d: Cart %d moves from (%d,%d) out of bounds to (%d,%d)\n", 
					tick, cart.id, oldX, oldY, cart.x, cart.y)
				cart.crashed = true
				
				// Check if this leaves only one cart
				remaining := 0
				var lastCart *FreshCart
				for _, c := range carts {
					if !c.crashed {
						remaining++
						lastCart = c
					}
				}
				
				if remaining == 1 {
					// The last cart might not have moved yet this tick
					if !movedThisTick[lastCart.id] {
						fmt.Printf("Cart %d hasn't moved yet in tick %d, at (%d,%d)\n",
							lastCart.id, tick, lastCart.x, lastCart.y)
						// It will move, so calculate its final position
						finalX := lastCart.x + lastCart.dx
						finalY := lastCart.y + lastCart.dy
						fmt.Printf("After its move, cart %d will be at (%d,%d)\n",
							lastCart.id, finalX, finalY)
						return fmt.Sprintf("%d,%d", finalX, finalY), nil
					} else {
						fmt.Printf("Last remaining cart %d already moved to (%d,%d)\n",
							lastCart.id, lastCart.x, lastCart.y)
						return fmt.Sprintf("%d,%d", lastCart.x, lastCart.y), nil
					}
				}
				continue
			}
			
			// Check for collisions with all other carts (including those that haven't moved yet)
			for j := 0; j < len(carts); j++ {
				other := carts[j]
				if other.id != cart.id && !other.crashed && 
				   other.x == cart.x && other.y == cart.y {
					fmt.Printf("Tick %d: Collision at (%d,%d) between carts %d and %d\n", 
						tick, cart.x, cart.y, cart.id, other.id)
					cart.crashed = true
					other.crashed = true
					break
				}
			}
			
			if cart.crashed {
				continue
			}
			
			// Update direction based on track
			track := grid[cart.y][cart.x]
			switch track {
			case '/':
				cart.dx, cart.dy = -cart.dy, -cart.dx
			case '\\':
				cart.dx, cart.dy = cart.dy, cart.dx
			case '+':
				turn := cart.turnCount % 3
				cart.turnCount++
				
				if turn == 0 {
					// Turn left
					cart.dx, cart.dy = cart.dy, -cart.dx
				} else if turn == 2 {
					// Turn right
					cart.dx, cart.dy = -cart.dy, cart.dx
				}
			}
		}
		
		// Count remaining after all moves
		remaining := 0
		var lastCart *FreshCart
		for _, cart := range carts {
			if !cart.crashed {
				remaining++
				lastCart = cart
			}
		}
		
		// Debug for specific ticks
		if tick >= 10002 && tick <= 10005 {
			fmt.Printf("End of tick %d: %d carts remaining\n", tick, remaining)
			for _, cart := range carts {
				if !cart.crashed {
					fmt.Printf("  Cart %d at (%d,%d) heading (%d,%d)\n", 
						cart.id, cart.x, cart.y, cart.dx, cart.dy)
				}
			}
		}
		
		if remaining == 1 {
			fmt.Printf("Final: Only cart %d remains at (%d,%d) after tick %d\n", 
				lastCart.id, lastCart.x, lastCart.y, tick)
			return fmt.Sprintf("%d,%d", lastCart.x, lastCart.y), nil
		}
		
		if remaining == 0 {
			return "", fmt.Errorf("all carts crashed")
		}
	}
	
	return "", fmt.Errorf("timeout")
}