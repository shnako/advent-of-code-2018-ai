package main

import (
	"fmt"
	"log"
	"os"
	
	"github.com/shnako/advent-of-code-2018-ai/solutions/day13"
)

func main() {
	// Test example
	example := `/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/`
	
	fmt.Println("Testing example:")
	sol := day13.NewOpus(example)
	result2, err := sol.Part2()
	if err != nil {
		log.Printf("Example Part 2 failed: %v", err)
	} else {
		fmt.Printf("Example Part 2: %s (expected: 6,4)\n", result2)
	}
	
	// Test actual input
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	
	sol = day13.NewOpus(string(input))
	
	result1, err := sol.Part1()
	if err != nil {
		log.Printf("Part 1 failed: %v", err)
	} else {
		fmt.Printf("Part 1: %s (expected: 109,23)\n", result1)
	}
	
	result2, err = sol.Part2Final()
	if err != nil {
		log.Printf("Part 2 failed: %v", err)
	} else {
		fmt.Printf("Part 2: %s\n", result2)
	}
}