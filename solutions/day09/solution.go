/*
 * Day 9: Marble Mania
 *
 * Part 1: Simulate a marble game with circular placement and special scoring rules.
 * Players take turns placing marbles in a circle. Normal marbles are placed between
 * marbles 1 and 2 positions clockwise. Marbles divisible by 23 score points and
 * remove a marble 7 positions counter-clockwise.
 *
 * Part 2: Same game but with 100 times more marbles, requiring efficient data structure.
 * Uses circular doubly-linked list for O(1) insertion and removal operations.
 */

package day09

import (
	"strconv"
	"strings"
)

// Marble represents a node in our circular doubly-linked list
type Marble struct {
	value int
	next  *Marble
	prev  *Marble
}

type Solution struct {
	players    int
	lastMarble int
}

func New(input string) *Solution {
	input = strings.ReplaceAll(strings.TrimSpace(input), "\r\n", "\n")
	parts := strings.Fields(input)

	players, err := strconv.Atoi(parts[0])
	if err != nil {
		players = 0 // Default value if parsing fails
	}
	lastMarble, err := strconv.Atoi(parts[6])
	if err != nil {
		lastMarble = 0 // Default value if parsing fails
	}

	return &Solution{
		players:    players,
		lastMarble: lastMarble,
	}
}

func (s *Solution) Part1() (int, error) {
	return s.playGame(s.lastMarble), nil
}

func (s *Solution) Part2() (int, error) {
	return s.playGame(s.lastMarble * 100), nil
}

// playGame simulates the marble game and returns the highest score
func (s *Solution) playGame(maxMarble int) int {
	scores := make([]int, s.players)

	// Create initial marble 0
	current := &Marble{value: 0}
	current.next = current
	current.prev = current

	for marble := 1; marble <= maxMarble; marble++ {
		player := (marble - 1) % s.players

		if marble%23 == 0 {
			// Special case: marble divisible by 23
			scores[player] += marble

			// Move 7 positions counter-clockwise
			for i := 0; i < 7; i++ {
				current = current.prev
			}

			// Add the marble being removed to score
			scores[player] += current.value

			// Remove the marble and update current
			current.prev.next = current.next
			current.next.prev = current.prev
			current = current.next
		} else {
			// Normal case: place marble between 1 and 2 clockwise
			current = current.next
			newMarble := &Marble{
				value: marble,
				next:  current.next,
				prev:  current,
			}
			current.next.prev = newMarble
			current.next = newMarble
			current = newMarble
		}
	}

	// Find maximum score
	maxScore := 0
	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}
