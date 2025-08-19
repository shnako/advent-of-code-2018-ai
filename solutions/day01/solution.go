/*
 * Day 1: Chronal Calibration
 * 
 * Part 1: Calculate the final frequency after applying all frequency changes.
 * Starting from frequency 0, we parse each line as a signed integer and add it
 * to our running total.
 * 
 * Part 2: Find the first frequency reached twice when repeatedly cycling through
 * the frequency changes. This requires keeping track of all seen frequencies and
 * detecting when we encounter a duplicate.
 */

package day01

import (
	"errors"
	"strconv"
	"strings"
)

type Solution struct {
	input string
}

func New(input string) *Solution {
	return &Solution{input: strings.TrimSpace(input)}
}

func (s *Solution) Part1() (int, error) {
	changes, err := s.parseChanges()
	if err != nil {
		return 0, err
	}

	frequency := 0
	for _, change := range changes {
		frequency += change
	}

	return frequency, nil
}

func (s *Solution) Part2() (int, error) {
	changes, err := s.parseChanges()
	if err != nil {
		return 0, err
	}
	if len(changes) == 0 {
		return 0, errors.New("no frequency changes provided")
	}

	frequency := 0
	seen := make(map[int]bool)
	seen[0] = true

	for {
		for _, change := range changes {
			frequency += change
			if seen[frequency] {
				return frequency, nil
			}
			seen[frequency] = true
		}
	}
}

func (s *Solution) parseChanges() ([]int, error) {
	lines := strings.Split(s.input, "\n")
	changes := make([]int, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		change, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		changes = append(changes, change)
	}

	return changes, nil
}