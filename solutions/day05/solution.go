/*
 * Day 5: Alchemical Reduction
 *
 * Part 1: Simulate polymer reactions where adjacent units of same type but opposite
 * polarity destroy each other. Use a stack to efficiently process reactions until
 * no more reactions are possible. Return the final polymer length.
 *
 * Part 2: Find the polymer unit type that, when completely removed, results in the
 * shortest final polymer after reactions. Try removing each letter type (a/A, b/B, etc.)
 * and find which gives the minimum result.
 */

package day05

import (
	"strings"
	"unicode"
)

type Solution struct {
	input string
}

func New(input string) *Solution {
	return &Solution{input: strings.TrimSpace(input)}
}

func (s *Solution) Part1() (int, error) {
	return s.reactPolymer(s.input), nil
}

func (s *Solution) Part2() (int, error) {
	minLength := len(s.input)

	// Try removing each unit type (a/A, b/B, c/C, etc.)
	for r := 'a'; r <= 'z'; r++ {
		// Skip if neither lower nor upper case exists to save work
		if !strings.ContainsRune(s.input, r) && !strings.ContainsRune(s.input, unicode.ToUpper(r)) {
			continue
		}

		// Create polymer without this unit type and react
		filtered := s.removeUnitType(s.input, r)
		length := s.reactPolymer(filtered)

		if length < minLength {
			minLength = length
			// Can't beat zero; short-circuit
			if minLength == 0 {
				break
			}
		}
	}

	return minLength, nil
}

// reactPolymer simulates polymer reactions using a stack approach
func (s *Solution) reactPolymer(polymer string) int {
	if len(polymer) == 0 {
		return 0
	}

	stack := make([]rune, 0, len(polymer))

	for _, unit := range polymer {
		if len(stack) > 0 && s.canReact(stack[len(stack)-1], unit) {
			// Remove the last unit from stack (reaction destroys both units)
			stack = stack[:len(stack)-1]
		} else {
			// Add unit to stack
			stack = append(stack, unit)
		}
	}

	return len(stack)
}

// canReact checks if two units can react (same type, opposite polarity)
func (s *Solution) canReact(a, b rune) bool {
	// Same type but opposite polarity means one is uppercase and other is lowercase
	// and they represent the same letter
	return unicode.ToLower(a) == unicode.ToLower(b) && a != b
}

// removeUnitType removes all instances of a specific unit type (both polarities)
func (s *Solution) removeUnitType(polymer string, unitType rune) string {
	var result strings.Builder
	result.Grow(len(polymer)) // Pre-allocate capacity

	for _, unit := range polymer {
		if unicode.ToLower(unit) == unitType {
			continue
		}
		result.WriteRune(unit)
	}

	return result.String()
}
