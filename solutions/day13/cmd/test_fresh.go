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
	
	fmt.Println("Testing example with fresh implementation:")
	sol := day13.NewFresh(example)
	result, err := sol.Part2()
	if err != nil {
		log.Printf("Example Part 2 failed: %v", err)
	} else {
		fmt.Printf("Example Part 2: %s (expected: 6,4)\n", result)
		if result != "6,4" {
			fmt.Println("❌ Example test failed!")
		} else {
			fmt.Println("✅ Example test passed!")
		}
	}
	
	// Test actual input
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	
	fmt.Println("\nTesting actual input with fresh implementation:")
	sol = day13.NewFresh(string(input))
	
	// Try standard approach
	fmt.Println("\nStandard approach:")
	result, err = sol.Part2()
	if err != nil {
		log.Printf("Part 2 failed: %v", err)
	} else {
		fmt.Printf("Part 2 result: %s\n", result)
		
		// Check against known wrong answers
		wrongAnswers := []string{
			"73,122", "74,122", "73,121", "72,122", "73,123", 
			"73,124", "122,73", "127,0", "74,121",
		}
		
		isWrong := false
		for _, wrong := range wrongAnswers {
			if result == wrong {
				fmt.Printf("❌ This is a known wrong answer: %s\n", wrong)
				isWrong = true
				break
			}
		}
		
		if !isWrong {
			fmt.Printf("✨ This is a NEW answer: %s\n", result)
		}
	}
	
	// Try alternative approach
	fmt.Println("\nAlternative approach (checking position after tick):")
	sol = day13.NewFresh(string(input))
	result, err = sol.Part2Alternative()
	if err != nil {
		log.Printf("Part 2 Alternative failed: %v", err)
	} else {
		fmt.Printf("Part 2 Alternative result: %s\n", result)
		
		// Check against known wrong answers
		wrongAnswers := []string{
			"73,122", "74,122", "73,121", "72,122", "73,123", 
			"73,124", "122,73", "127,0", "74,121",
		}
		
		isWrong := false
		for _, wrong := range wrongAnswers {
			if result == wrong {
				fmt.Printf("❌ This is a known wrong answer: %s\n", wrong)
				isWrong = true
				break
			}
		}
		
		if !isWrong {
			fmt.Printf("✨ This is a NEW answer: %s\n", result)
		}
	}
}