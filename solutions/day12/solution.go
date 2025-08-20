/*
 * Day 12: Subterranean Sustainability
 * 
 * Part 1: Simulate cellular automata for plants in pots for 20 generations.
 * Each pot's next state depends on its 5-pot neighborhood (2 left + self + 2 right).
 * Sum the pot numbers containing plants after 20 generations.
 * 
 * Part 2: TBD after Part 1 is solved and Part 2 is revealed.
 */

package day12

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct {
	initialState string
	rules        map[string]rune
}

func New(input string) (*Solution, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) < 2 {
		return nil, fmt.Errorf("invalid input format")
	}

	// Parse initial state
	initialLine := lines[0]
	if !strings.HasPrefix(initialLine, "initial state: ") {
		return nil, fmt.Errorf("invalid initial state format")
	}
	initialState := strings.TrimPrefix(initialLine, "initial state: ")

	// Parse rules
	rules := make(map[string]rune)
	for i := 2; i < len(lines); i++ { // Skip empty line at index 1
		parts := strings.Split(lines[i], " => ")
		if len(parts) != 2 {
			continue
		}
		pattern := parts[0]
		result := rune(parts[1][0])
		rules[pattern] = result
	}

	return &Solution{
		initialState: initialState,
		rules:        rules,
	}, nil
}

func (s *Solution) Part1() (string, error) {
	result := s.simulate(20)
	return strconv.Itoa(result), nil
}

func (s *Solution) Part2() (string, error) {
	// For 50 billion generations, we need to detect when the pattern stabilizes
	// The sum will eventually grow by a constant amount each generation
	result := s.simulateUntilStable(50000000000)
	return strconv.Itoa(result), nil
}

func (s *Solution) simulate(generations int) int {
	// Use map to represent pots, where key is pot number and value is whether it has a plant
	pots := make(map[int]bool)
	
	// Initialize pots from initial state (pot 0 is at position 0)
	for i, char := range s.initialState {
		if char == '#' {
			pots[i] = true
		}
	}

	// Simulate generations
	for gen := 0; gen < generations; gen++ {
		pots = s.nextGeneration(pots)
	}

	// Sum pot numbers containing plants
	sum := 0
	for potNum, hasPlant := range pots {
		if hasPlant {
			sum += potNum
		}
	}
	
	return sum
}

func (s *Solution) nextGeneration(currentPots map[int]bool) map[int]bool {
	nextPots := make(map[int]bool)
	
	// Find the range of pots to check (expand by 2 on each side)
	minPot, maxPot := s.getPotRange(currentPots)
	minPot -= 2
	maxPot += 2
	
	// Check each pot in the expanded range
	for pot := minPot; pot <= maxPot; pot++ {
		pattern := s.getPattern(currentPots, pot)
		if result, exists := s.rules[pattern]; exists && result == '#' {
			nextPots[pot] = true
		}
	}
	
	return nextPots
}

func (s *Solution) getPotRange(pots map[int]bool) (int, int) {
	if len(pots) == 0 {
		return 0, 0
	}
	
	minPot := 1000000
	maxPot := -1000000
	
	for potNum, hasPlant := range pots {
		if hasPlant {
			if potNum < minPot {
				minPot = potNum
			}
			if potNum > maxPot {
				maxPot = potNum
			}
		}
	}
	
	return minPot, maxPot
}

func (s *Solution) getPattern(pots map[int]bool, centerPot int) string {
	var pattern strings.Builder
	
	for offset := -2; offset <= 2; offset++ {
		pot := centerPot + offset
		if pots[pot] {
			pattern.WriteRune('#')
		} else {
			pattern.WriteRune('.')
		}
	}
	
	return pattern.String()
}

func (s *Solution) simulateUntilStable(targetGenerations int) int {
	// Use map to represent pots
	pots := make(map[int]bool)
	
	// Initialize pots from initial state
	for i, char := range s.initialState {
		if char == '#' {
			pots[i] = true
		}
	}

	// Track sums to detect stability
	prevSum := s.calculateSum(pots)
	prevDiff := 0
	stableDiffCount := 0
	
	for gen := 1; gen <= targetGenerations; gen++ {
		pots = s.nextGeneration(pots)
		currentSum := s.calculateSum(pots)
		
		diff := currentSum - prevSum
		
		// Check if the difference has stabilized
		if diff == prevDiff {
			stableDiffCount++
		} else {
			stableDiffCount = 0
		}
		
		// If stable for several generations, we can extrapolate
		if stableDiffCount >= 10 {
			remainingGens := targetGenerations - gen
			return currentSum + (diff * remainingGens)
		}
		
		prevSum = currentSum
		prevDiff = diff
	}
	
	return prevSum
}

func (s *Solution) calculateSum(pots map[int]bool) int {
	sum := 0
	for potNum, hasPlant := range pots {
		if hasPlant {
			sum += potNum
		}
	}
	return sum
}