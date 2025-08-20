package main

import (
	"fmt"
	"log"
	"os"
	
	"github.com/shnako/advent-of-code-2018-ai/solutions/day13"
)

func main() {
	// Test actual input with "before tick" interpretation
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	
	fmt.Println("Testing 'before tick' interpretation:")
	sol := day13.NewBeforeTick(string(input))
	
	result, err := sol.Part2()
	if err != nil {
		log.Printf("Part 2 failed: %v", err)
	} else {
		fmt.Printf("Part 2 result (position before final tick): %s\n", result)
		
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
			fmt.Println("This would be cart 0's position at the START of tick 10004")
		}
	}
}