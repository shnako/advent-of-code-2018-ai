package day20

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
		{"example 1", "^WNE$", 3},
		{"example 2", "^ENWWW(NEEE|SSE(EE|N))$", 10},
		{"example 3", "^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$", 18},
		{"example 4", "^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$", 23},
		{"example 5", "^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$", 31},
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
	
	expected := 4025
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
	
	expected := 8186
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}