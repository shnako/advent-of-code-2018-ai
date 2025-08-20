package utils

import (
	"strconv"
	"strings"
)

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

func SplitLines(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func SplitByEmptyLines(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n\n")
}

func ParseGrid(input string) [][]rune {
	lines := SplitLines(input)
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
