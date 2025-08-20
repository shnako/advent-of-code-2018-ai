package day13

import (
	"fmt"
	"sort"
	"strings"
)

type DebugCart struct {
	X, Y    int
	DX, DY  int
	Turns   int
	ID      int
	Crashed bool
}

func NewDebug(input string) *OpusSolution {
	return &OpusSolution{input: input}
}

func (s *OpusSolution) Part2Debug() (string, error) {
	lines := strings.Split(s.input, "\n")
	grid := make([][]byte, len(lines))
	var carts []*DebugCart
	
	cartID := 0
	for y := 0; y < len(lines); y++ {
		grid[y] = []byte(lines[y])
		for x := 0; x < len(grid[y]); x++ {
			var cart *DebugCart
			switch grid[y][x] {
			case '^':
				cart = &DebugCart{X: x, Y: y, DX: 0, DY: -1, Turns: 0, ID: cartID}
				grid[y][x] = '|'
			case 'v':
				cart = &DebugCart{X: x, Y: y, DX: 0, DY: 1, Turns: 0, ID: cartID}
				grid[y][x] = '|'
			case '<':
				cart = &DebugCart{X: x, Y: y, DX: -1, DY: 0, Turns: 0, ID: cartID}
				grid[y][x] = '-'
			case '>':
				cart = &DebugCart{X: x, Y: y, DX: 1, DY: 0, Turns: 0, ID: cartID}
				grid[y][x] = '-'
			}
			if cart != nil {
				carts = append(carts, cart)
				cartID++
			}
		}
	}
	
	fmt.Printf("Starting with %d carts\n", len(carts))
	
	for tick := 0; tick < 100000; tick++ {
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
			
			// Check bounds
			if cart.Y < 0 || cart.Y >= len(grid) || cart.X < 0 || cart.X >= len(grid[cart.Y]) {
				cart.Crashed = true
				if tick > 9000 && tick < 11000 {
					fmt.Printf("Tick %d: Cart %d crashed out of bounds at (%d,%d)\n", tick, cart.ID, cart.X, cart.Y)
				}
				continue
			}
			
			// Check for collision
			for j := 0; j < len(carts); j++ {
				if i != j && !carts[j].Crashed && carts[j].X == cart.X && carts[j].Y == cart.Y {
					cart.Crashed = true
					carts[j].Crashed = true
					if tick > 9000 && tick < 11000 {
						fmt.Printf("Tick %d: Carts %d and %d collided at (%d,%d)\n", tick, cart.ID, carts[j].ID, cart.X, cart.Y)
					}
					break
				}
			}
			
			if cart.Crashed {
				continue
			}
			
			// Handle track piece
			track := grid[cart.Y][cart.X]
			switch track {
			case '/':
				cart.DX, cart.DY = -cart.DY, -cart.DX
			case '\\':
				cart.DX, cart.DY = cart.DY, cart.DX
			case '+':
				turn := cart.Turns % 3
				cart.Turns++
				if turn == 0 {
					// Turn left
					cart.DX, cart.DY = cart.DY, -cart.DX
				} else if turn == 2 {
					// Turn right
					cart.DX, cart.DY = -cart.DY, cart.DX
				}
			}
		}
		
		// Count alive carts
		alive := 0
		var lastCart *DebugCart
		for _, cart := range carts {
			if !cart.Crashed {
				alive++
				lastCart = cart
			}
		}
		
		if tick > 9000 && tick < 10100 {
			fmt.Printf("After tick %d: %d carts alive\n", tick, alive)
			if alive <= 5 {
				for _, cart := range carts {
					if !cart.Crashed {
						fmt.Printf("  Cart %d at (%d,%d)\n", cart.ID, cart.X, cart.Y)
					}
				}
			}
		}
		
		if alive == 1 {
			fmt.Printf("Final: After tick %d, only cart %d remains at (%d,%d)\n", tick, lastCart.ID, lastCart.X, lastCart.Y)
			return fmt.Sprintf("%d,%d", lastCart.X, lastCart.Y), nil
		}
		if alive == 0 {
			return "", fmt.Errorf("all carts crashed")
		}
	}
	
	return "", fmt.Errorf("timeout")
}