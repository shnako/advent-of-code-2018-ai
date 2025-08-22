// Package utils provides common utility functions for Advent of Code solutions.
package utils

import (
	"strconv"
	"strings"
)

// ParseInts parses lines of integers from input string.
// Empty lines are skipped. Returns an error if any non-empty line cannot be parsed.
func ParseInts(input string) ([]int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	nums := make([]int, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

// ParseIntsFromString parses integers from a string with a custom separator.
// Empty parts are skipped. Returns an error if any non-empty part cannot be parsed.
func ParseIntsFromString(s string, sep string) ([]int, error) {
	parts := strings.Split(s, sep)
	nums := make([]int, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		n, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

// SplitLines splits input into lines, trimming surrounding whitespace.
func SplitLines(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

// SplitByEmptyLines splits input by double newlines (empty lines).
func SplitByEmptyLines(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n\n")
}

// ParseGrid parses input into a 2D grid of runes.
func ParseGrid(input string) [][]rune {
	lines := SplitLines(input)
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

// MustAtoi converts a string to int, panicking on error.
// Use only when you're certain the input is valid.
func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
