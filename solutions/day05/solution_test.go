package day05

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
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"empty", "", 0},
		{"no reaction", "abAB", 4},
		{"simple reaction", "aA", 0},
		{"chain reaction", "abBA", 0},
		{"complex example", "dabAcCaCBAcCcaDA", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solution := New(tt.input)
			result, err := solution.Part1()

			if err != nil {
				t.Errorf("Part1() error = %v", err)
				return
			}

			if result != tt.expected {
				t.Errorf("Part1() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPart2Examples(t *testing.T) {
	input := "dabAcCaCBAcCcaDA"
	
	solution := New(input)
	result, err := solution.Part2()

	if err != nil {
		t.Errorf("Part2() error = %v", err)
		return
	}

	expected := 4 // After removing 'c/C', result is "daDA" which reacts to "da" = 4 units
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

	expected := 10878
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

	expected := 6874
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}