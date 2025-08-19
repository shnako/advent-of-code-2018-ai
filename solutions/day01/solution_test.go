package day01

import (
	"os"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"example 1", "+1\n-2\n+3\n+1", 3},
		{"example 2", "+1\n+1\n+1", 3},
		{"example 3", "+1\n+1\n-2", 0},
		{"example 4", "-1\n-2\n-3", -6},
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

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("Failed to read input: %v", err)
	}

	solution := New(string(input))
	result, err := solution.Part1()

	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}

	expected := 459
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
		t.Errorf("Part2() error = %v", err)
		return
	}

	expected := 65474
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}