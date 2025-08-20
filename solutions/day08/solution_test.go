package day08

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
	input := `2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`

	solution := New(input)
	result, err := solution.Part1()

	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}

	expected := 138
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart1(t *testing.T) {
	solution := New(readInput(t))
	result, err := solution.Part1()

	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}

	expected := 36891
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	} else {
		t.Logf("Part1() = %v", result)
	}
}

func TestPart2Example(t *testing.T) {
	input := `2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`

	solution := New(input)
	result, err := solution.Part2()

	if err != nil {
		t.Errorf("Part2() error = %v", err)
		return
	}

	expected := 66
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	solution := New(readInput(t))
	result, err := solution.Part2()

	if err != nil {
		t.Errorf("Part2() error = %v", err)
		return
	}

	expected := 20083
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	} else {
		t.Logf("Part2() = %v", result)
	}
}
