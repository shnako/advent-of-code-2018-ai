package day03

import (
	"os"
	"testing"
)

func readInput(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("Failed to read input: %v", err)
	}
	return string(b)
}

func TestPart1Examples(t *testing.T) {
	input := `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`

	solution := New(input)
	result, err := solution.Part1()

	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}

	expected := 4 // The four square inches marked with X
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart2Examples(t *testing.T) {
	input := `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`

	solution := New(input)
	result, err := solution.Part2()

	if err != nil {
		t.Errorf("Part2() error = %v", err)
		return
	}

	expected := 3 // Claim 3 doesn't overlap
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}

func TestPart1(t *testing.T) {
	solution := New(readInput(t))
	result, err := solution.Part1()

	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}

	expected := 113966
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	solution := New(readInput(t))
	result, err := solution.Part2()

	if err != nil {
		t.Errorf("Part2() error = %v", err)
		return
	}

	expected := 235
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}
