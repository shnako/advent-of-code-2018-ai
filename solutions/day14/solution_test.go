package day14

import (
	"os"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"9", "5158916779"},
		{"5", "0124515891"},
		{"18", "9251071085"},
		{"2018", "5941429882"},
	}

	for _, test := range tests {
		t.Run("input_"+test.input, func(t *testing.T) {
			solution := New(test.input)
			result, err := solution.Part1()

			if err != nil {
				t.Errorf("Part1() error = %v", err)
				return
			}

			if result != test.expected {
				t.Errorf("Part1() = %v, want %v", result, test.expected)
			}
		})
	}
}

func TestPart2Examples(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"51589", 9},
		{"01245", 5},
		{"92510", 18},
		{"59414", 2018},
	}

	for _, test := range tests {
		t.Run("input_"+test.input, func(t *testing.T) {
			solution := New(test.input)
			result, err := solution.Part2()

			if err != nil {
				t.Errorf("Part2() error = %v", err)
				return
			}

			if result != test.expected {
				t.Errorf("Part2() = %v, want %v", result, test.expected)
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

	expected := "1776718175" // Confirmed correct
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

	expected := 20220949 // Confirmed correct
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}
