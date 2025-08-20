package day19

import (
	"os"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input := `#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5`
	
	solution := New(input)
	result, err := solution.Part1()
	
	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}
	
	expected := 6
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
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
	
	expected := 888
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
	
	expected := 10708992
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}