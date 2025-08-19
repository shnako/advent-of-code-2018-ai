/*
 * Day 3: No Matter How You Slice It
 * 
 * Part 1: Count square inches of fabric that are within two or more claims.
 * Parse fabric claims in format "#ID @ x,y: wxh" and track overlapping areas
 * on a 2D grid to count squares claimed by multiple elves.
 * 
 * Part 2: Find the ID of the claim that doesn't overlap with any other claim.
 * This requires tracking which specific claims occupy each square inch.
 */

package day03

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Claim struct {
	ID     int
	X, Y   int
	Width  int
	Height int
}

type Solution struct {
	input string
}

func New(input string) *Solution {
	return &Solution{input: strings.TrimSpace(input)}
}

func (s *Solution) Part1() (int, error) {
	claims, err := s.parseClaims()
	if err != nil {
		return 0, err
	}

	// Track how many claims each square inch has
	fabric := make(map[string]int)

	for _, claim := range claims {
		for x := claim.X; x < claim.X+claim.Width; x++ {
			for y := claim.Y; y < claim.Y+claim.Height; y++ {
				key := strconv.Itoa(x) + "," + strconv.Itoa(y)
				fabric[key]++
			}
		}
	}

	// Count squares with 2+ claims
	overlapping := 0
	for _, count := range fabric {
		if count >= 2 {
			overlapping++
		}
	}

	return overlapping, nil
}

func (s *Solution) Part2() (int, error) {
	claims, err := s.parseClaims()
	if err != nil {
		return 0, err
	}

	// Track which claims occupy each square inch
	fabric := make(map[string][]int)

	for _, claim := range claims {
		for x := claim.X; x < claim.X+claim.Width; x++ {
			for y := claim.Y; y < claim.Y+claim.Height; y++ {
				key := strconv.Itoa(x) + "," + strconv.Itoa(y)
				fabric[key] = append(fabric[key], claim.ID)
			}
		}
	}

	// Find claims that never overlap with others
	overlapping := make(map[int]bool)
	for _, claimIDs := range fabric {
		if len(claimIDs) > 1 {
			for _, id := range claimIDs {
				overlapping[id] = true
			}
		}
	}

	// Return the claim that's not overlapping
	for _, claim := range claims {
		if !overlapping[claim.ID] {
			return claim.ID, nil
		}
	}

	return 0, errors.New("no non-overlapping claim found")
}

func (s *Solution) parseClaims() ([]Claim, error) {
	lines := strings.Split(s.input, "\n")
	claims := make([]Claim, 0, len(lines))

	// Regular expression to parse format: #123 @ 3,2: 5x4
	re := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		matches := re.FindStringSubmatch(line)
		if len(matches) != 6 {
			return nil, errors.New("invalid claim format: " + line)
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, err
		}

		x, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, err
		}

		y, err := strconv.Atoi(matches[3])
		if err != nil {
			return nil, err
		}

		width, err := strconv.Atoi(matches[4])
		if err != nil {
			return nil, err
		}

		height, err := strconv.Atoi(matches[5])
		if err != nil {
			return nil, err
		}

		claims = append(claims, Claim{
			ID:     id,
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
		})
	}

	return claims, nil
}