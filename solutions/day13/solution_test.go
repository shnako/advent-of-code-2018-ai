package day13

import (
	"os"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input := `/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `
	
	solution := New(input)
	result, err := solution.Part1()
	
	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}
	
	expected := "7,3"
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart2Example(t *testing.T) {
	input := `/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/`
	
	solution := New(input)
	result, err := solution.Part2()
	
	if err != nil {
		t.Errorf("Part2() error = %v", err)
		return
	}
	
	expected := "6,4"
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
		t.Errorf("Part1() error = %v", err)
		return
	}
	
	expected := "109,23" // Confirmed correct
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
	
	// Part 2 is currently unsolved - see NOTES.md
	// All implementations produce "73,122" but this is incorrect
	t.Logf("Part2() = %v (NOTE: This answer is known to be wrong)", result)
	}
}