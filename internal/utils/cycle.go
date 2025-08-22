// Package utils provides common utility functions for Advent of Code solutions.
package utils

// CycleDetector helps detect cycles in sequences.
// It tracks seen states and identifies when a cycle occurs.
type CycleDetector[T comparable] struct {
	seen  map[T]int
	index int
}

// NewCycleDetector creates a new cycle detector.
func NewCycleDetector[T comparable]() *CycleDetector[T] {
	return &CycleDetector[T]{
		seen: make(map[T]int),
	}
}

// Add adds a state and returns true if a cycle is detected,
// along with the cycle start index and cycle length.
func (cd *CycleDetector[T]) Add(state T) (bool, int, int) {
	if startIdx, exists := cd.seen[state]; exists {
		cycleLen := cd.index - startIdx
		return true, startIdx, cycleLen
	}
	cd.seen[state] = cd.index
	cd.index++
	return false, 0, 0
}

// Reset clears the cycle detector.
func (cd *CycleDetector[T]) Reset() {
	cd.seen = make(map[T]int)
	cd.index = 0
}