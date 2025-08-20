package day09

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
		{"10 players, 1618 marbles", "10 players; last marble is worth 1618 points", 8317},
		{"13 players, 7999 marbles", "13 players; last marble is worth 7999 points", 146373},
		{"17 players, 1104 marbles", "17 players; last marble is worth 1104 points", 2764},
		{"21 players, 6111 marbles", "21 players; last marble is worth 6111 points", 54718},
		{"30 players, 5807 marbles", "30 players; last marble is worth 5807 points", 37305},
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
	solution := New(readInput(t))
	result, err := solution.Part1()

	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}

	expected := 418237
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	} else {
		t.Logf("Part1() = %v", result)
	}
}

func TestPart2(t *testing.T) {
	solution := New(readInput(t))
	result, err := solution.Part2()

	if err != nil {
		t.Errorf("Part2() error = %v", err)
		return
	}

	expected := 3505711612
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	} else {
		t.Logf("Part2() = %v", result)
	}
}
