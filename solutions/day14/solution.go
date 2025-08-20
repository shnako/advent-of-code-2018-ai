/*
 * Day 14: Chocolate Charts
 * 
 * Part 1: Simulate the recipe generation process and find the scores of the ten recipes
 * immediately after the number of recipes specified in the puzzle input.
 * Two elves start with recipes scoring 3 and 7, and create new recipes by combining their
 * current recipes' scores and adding the digits of the sum to the scoreboard.
 * 
 * Part 2: Find how many recipes appear on the scoreboard to the left of the sequence
 * that matches the puzzle input digits.
 */

package day14

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct {
	input string
}

func New(input string) *Solution {
	return &Solution{input: strings.TrimSpace(input)}
}

func (s *Solution) Part1() (string, error) {
	target, err := strconv.Atoi(s.input)
	if err != nil {
		return "", err
	}
	
	// Start with recipes 3 and 7; preallocate to target+10 to reduce reallocations
	capacity := target + 12 // a small buffer above target+10
	if capacity < 2 {
		capacity = 2
	}
	recipes := make([]int, 0, capacity)
	recipes = append(recipes, 3, 7)
	elf1, elf2 := 0, 1
	
	// Continue until we have at least target + 10 recipes
	for len(recipes) < target+10 {
		// Create new recipes from sum of current recipes
		sum := recipes[elf1] + recipes[elf2]
		
		// Add digits of sum to recipes
		if sum >= 10 {
			recipes = append(recipes, sum/10, sum%10)
		} else {
			recipes = append(recipes, sum)
		}
		
		// Move elves to new positions
		elf1 = (elf1 + 1 + recipes[elf1]) % len(recipes)
		elf2 = (elf2 + 1 + recipes[elf2]) % len(recipes)
	}
	
	// Get the ten recipes after target recipes
	var b strings.Builder
	b.Grow(10)
	for i := target; i < target+10; i++ {
		b.WriteByte(byte('0' + recipes[i]))
	}
	return b.String(), nil
}

func (s *Solution) Part2() (int, error) {
	targetStr := s.input
	targetLen := len(targetStr)
	if targetLen == 0 {
		return 0, fmt.Errorf("empty input")
	}
	
	// Convert target string to slice of ints for comparison
	target := make([]int, targetLen)
	for i, char := range targetStr {
		if char < '0' || char > '9' {
			return 0, fmt.Errorf("input contains non-digit at position %d: %q", i, char)
		}
		target[i] = int(char - '0')
	}
	
	// Start with recipes 3 and 7
	recipes := []int{3, 7}
	elf1, elf2 := 0, 1
	
	// Continue until we find the target sequence
	for {
		// Create new recipes from sum of current recipes
		sum := recipes[elf1] + recipes[elf2]
		
		// Add digits of sum to recipes and check after each addition
		if sum >= 10 {
			recipes = append(recipes, sum/10)
			// Check if we found the target sequence
			if len(recipes) >= targetLen && s.matchesTarget(recipes, target, len(recipes)-targetLen) {
				return len(recipes) - targetLen, nil
			}
			
			recipes = append(recipes, sum%10)
			// Check again after adding the second digit
			if len(recipes) >= targetLen && s.matchesTarget(recipes, target, len(recipes)-targetLen) {
				return len(recipes) - targetLen, nil
			}
		} else {
			recipes = append(recipes, sum)
			// Check if we found the target sequence
			if len(recipes) >= targetLen && s.matchesTarget(recipes, target, len(recipes)-targetLen) {
				return len(recipes) - targetLen, nil
			}
		}
		
		// Move elves to new positions
		elf1 = (elf1 + 1 + recipes[elf1]) % len(recipes)
		elf2 = (elf2 + 1 + recipes[elf2]) % len(recipes)
	}
}

// Helper function to check if the target sequence matches at the given position
func (s *Solution) matchesTarget(recipes []int, target []int, pos int) bool {
	if pos < 0 || pos+len(target) > len(recipes) {
		return false
	}
	for i := 0; i < len(target); i++ {
		if recipes[pos+i] != target[i] {
			return false
		}
	}
	return true
}