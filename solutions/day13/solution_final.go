package day13

import (
	"fmt"
	"sort"
	"strings"
)

type FinalCart struct {
	X, Y    int
	DX, DY  int
	Turns   int
	ID      int
	Alive   bool
}

func NewFinal(input string) *OpusSolution {
	return &OpusSolution{input: input}
}

func (s *OpusSolution) Part2Final() (string, error) {
	lines := strings.Split(s.input, "\n")
	grid := make([][]byte, len(lines))
	carts := []*FinalCart{}
	
	cartID := 0
	for y := 0; y < len(lines); y++ {
		grid[y] = []byte(lines[y])
		for x := 0; x < len(grid[y]); x++ {
			var cart *FinalCart
			switch grid[y][x] {
			case '^':
				cart = &FinalCart{X: x, Y: y, DX: 0, DY: -1, ID: cartID, Alive: true}
				grid[y][x] = '|'
			case 'v':
				cart = &FinalCart{X: x, Y: y, DX: 0, DY: 1, ID: cartID, Alive: true}
				grid[y][x] = '|'
			case '<':
				cart = &FinalCart{X: x, Y: y, DX: -1, DY: 0, ID: cartID, Alive: true}
				grid[y][x] = '-'
			case '>':
				cart = &FinalCart{X: x, Y: y, DX: 1, DY: 0, ID: cartID, Alive: true}
				grid[y][x] = '-'
			}
			if cart != nil {
				carts = append(carts, cart)
				cartID++
			}
		}
	}
	
	fmt.Printf("Starting with %d carts\n", len(carts))
	
	for tick := 0; tick < 15000; tick++ {
		// Sort carts by Y then X (for those that are alive)
		sort.Slice(carts, func(i, j int) bool {
			// Dead carts go to the end
			if !carts[i].Alive && carts[j].Alive {
				return false
			}
			if carts[i].Alive && !carts[j].Alive {
				return true
			}
			// Both alive or both dead - sort by position
			if carts[i].Y != carts[j].Y {
				return carts[i].Y < carts[j].Y
			}
			return carts[i].X < carts[j].X
		})
		
		// Process each cart in order
		for i := 0; i < len(carts); i++ {
			cart := carts[i]
			if !cart.Alive {
				continue
			}
			
			// Check if movement would go out of bounds BEFORE moving
			newX := cart.X + cart.DX
			newY := cart.Y + cart.DY
			
			if newY < 0 || newY >= len(grid) || newX < 0 || newX >= len(grid[newY]) {
				fmt.Printf("Tick %d: Cart %d would go OOB from (%d,%d) to (%d,%d), removing it\n", 
					tick, cart.ID, cart.X, cart.Y, newX, newY)
				cart.Alive = false
				continue
			}
			
			// Actually move
			cart.X = newX
			cart.Y = newY
			
			// Check for collision
			for j := 0; j < len(carts); j++ {
				if i != j && carts[j].Alive && carts[j].X == cart.X && carts[j].Y == cart.Y {
					fmt.Printf("Tick %d: Collision! Carts %d and %d at (%d,%d)\n", 
						tick, cart.ID, carts[j].ID, cart.X, cart.Y)
					cart.Alive = false
					carts[j].Alive = false
					break
				}
			}
			
			if !cart.Alive {
				continue
			}
			
			// Handle track piece for direction change
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
		
		// Count alive
		alive := 0
		var lastCart *FinalCart
		for _, cart := range carts {
			if cart.Alive {
				alive++
				lastCart = cart
			}
		}
		
		if tick >= 10000 && tick <= 10010 {
			fmt.Printf("After tick %d: %d carts alive\n", tick, alive)
			for _, cart := range carts {
				if cart.Alive {
					fmt.Printf("  Cart %d at (%d,%d) facing (%d,%d)\n", 
						cart.ID, cart.X, cart.Y, cart.DX, cart.DY)
				}
			}
		}
		
		if alive == 1 {
			fmt.Printf("Tick %d complete: Only cart %d remains at (%d,%d)\n", 
				tick, lastCart.ID, lastCart.X, lastCart.Y)
			return fmt.Sprintf("%d,%d", lastCart.X, lastCart.Y), nil
		}
		
		if alive == 0 {
			fmt.Printf("Tick %d: All carts eliminated\n", tick)
			return "", fmt.Errorf("no carts remain")
		}
	}
	
	return "", fmt.Errorf("timeout")
}