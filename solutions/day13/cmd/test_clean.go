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
	solExample := day13.NewClean(example)
	result, err := solExample.Part2Clean()
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
	
	sol := day13.NewClean(string(input))
	
	fmt.Println("Testing Part 1 for verification:")
	result1, err := sol.Part1Clean()
	if err != nil {
		log.Printf("Part 1 failed: %v", err)
	} else {
		fmt.Printf("Part 1 result: %s (expected: 109,23)\n", result1)
	}
	
	fmt.Println("\nTesting Part 2:")
	result2, err := sol.Part2Clean()
	if err != nil {
		log.Printf("Part 2 failed: %v", err)
	} else {
		fmt.Printf("Part 2 result: %s\n", result2)
	}
}