package day06

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

func TestPart1Example(t *testing.T) {
	input := `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

	solution := New(input)
	result, err := solution.Part1()

	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}

	expected := 17
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart2Example(t *testing.T) {
	input := `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

	solution := New(input)

	// Test with the example threshold of 32 (should give 16)
	result := solution.countRegionSize(32)
	expected := 16
	if result != expected {
		t.Errorf("countRegionSize(32) = %v, want %v", result, expected)
	}
}

func TestPart1(t *testing.T) {
	solution := New(readInput(t))
	result, err := solution.Part1()

	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}

	expected := 4284
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

	expected := 35490
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}
