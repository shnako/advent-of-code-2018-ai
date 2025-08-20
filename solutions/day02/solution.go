/*
 * Day 2: Inventory Management System
 *
 * Part 1: Calculate checksum by counting box IDs that have exactly two of any letter
 * and separately counting those with exactly three of any letter, then multiply.
 *
 * Part 2: Find the two correct box IDs that differ by exactly one character at the
 * same position, then return the common letters between them.
 */

package day02

import (
	"errors"
	"strings"
)

type Solution struct {
	input string
}

func New(input string) *Solution {
	return &Solution{input: strings.TrimSpace(input)}
}

func (s *Solution) Part1() (int, error) {
	lines := s.parseLines()

	twos := 0
	threes := 0

	for _, line := range lines {

		var counts [26]int
		for i := 0; i < len(line); i++ {
			c := line[i]
			idx := int(c - 'a')
			if idx >= 0 && idx < 26 {
				counts[idx]++
			}
		}

		hasTwo := false
		hasThree := false
		for _, count := range counts {
			if count == 2 {
				hasTwo = true
			}
			if count == 3 {
				hasThree = true
			}
			if hasTwo && hasThree {
				break
			}
		}

		if hasTwo {
			twos++
		}
		if hasThree {
			threes++
		}
	}

	return twos * threes, nil
}

func (s *Solution) Part2() (string, error) {
	lines := s.parseLines()

	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if diffByOne(lines[i], lines[j]) {
				return commonLetters(lines[i], lines[j]), nil
			}
		}
	}

	return "", errors.New("no pair of IDs differing by exactly one character was found")
}

func (s *Solution) parseLines() []string {
	lines := strings.Split(s.input, "\n")
	result := make([]string, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			result = append(result, line)
		}
	}

	return result
}

func diffByOne(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	differences := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			differences++
			if differences > 1 {
				return false
			}
		}
	}

	return differences == 1
}

func commonLetters(a, b string) string {
	var result strings.Builder
	result.Grow(len(a))

	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] == b[i] {
			result.WriteByte(a[i])
		}
	}

	return result.String()
}
