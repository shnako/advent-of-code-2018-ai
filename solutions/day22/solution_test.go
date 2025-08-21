package day22

import (
	"os"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input := `depth: 510
target: 10,10`

	solution := New(input)
	result, err := solution.Part1()

	if err != nil {
		t.Fatalf("Part1() error = %v", err)
	}

	expected := 114
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart2Example(t *testing.T) {
	input := `depth: 510
target: 10,10`

	solution := New(input)
	result, err := solution.Part2()

	if err != nil {
		t.Fatalf("Part2() error = %v", err)
	}

	expected := 45
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("Failed to read input: %v", err)
	}

	solution := New(string(input))
	result, err := solution.Part1()

	if err != nil {
		t.Fatalf("Part1() error = %v", err)
	}

	expected := 8575
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("Failed to read input: %v", err)
	}

	solution := New(string(input))
	result, err := solution.Part2()

	if err != nil {
		t.Fatalf("Part2() error = %v", err)
	}

	expected := 999
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}