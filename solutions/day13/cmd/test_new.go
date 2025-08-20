package main

import (
	"fmt"
	"log"
	"os"
	
	"github.com/shnako/advent-of-code-2018-ai/solutions/day13"
)

func main() {
	// Test with example first
	example := `/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/`
	
	fmt.Println("Testing Part 2 example:")
	sol := day13.NewSolutionFromScratch(example)
	result, err := sol.Part2Debug()
	if err != nil {
		log.Printf("Part 2 example failed: %v", err)
	} else {
		fmt.Printf("Part 2 example result: %s (expected: 6,4)\n", result)
	}
	
	// Test with actual input
	fmt.Println("\nTesting with actual input:")
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	
	sol2 := day13.NewSolutionFromScratch(string(input))
	
	fmt.Println("Testing Part 1 for verification:")
	result1, err := sol2.Part1Debug()
	if err != nil {
		log.Printf("Part 1 failed: %v", err)
	} else {
		fmt.Printf("Part 1 result: %s (expected: 109,23)\n", result1)
	}
	
	fmt.Println("\nChecking grid around problem area:")
	grid, _ := sol2.ParseGrid()
	fmt.Printf("Grid at (127,0): '%c'\n", grid.Get(127, 0))
	fmt.Printf("Grid at (127,1): '%c'\n", grid.Get(127, 1))
	fmt.Printf("Grid at (126,0): '%c'\n", grid.Get(126, 0))
	fmt.Printf("Grid at (128,0): '%c'\n", grid.Get(128, 0))
	
	fmt.Println("\nTesting Part 2:")
	result2, err := sol2.Part2Debug()
	if err != nil {
		log.Printf("Part 2 failed: %v", err)
	} else {
		fmt.Printf("Part 2 result: %s\n", result2)
	}
}