package day11

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

func TestGetPowerLevel(t *testing.T) {
	tests := []struct {
		x, y, serial, expected int
	}{
		{3, 5, 8, 4},
		{122, 79, 57, -5},
		{217, 196, 39, 0},
		{101, 153, 71, 4},
	}
	
	for _, tt := range tests {
		solution := &Solution{serialNumber: tt.serial}
		result := solution.getPowerLevel(tt.x, tt.y)
		if result != tt.expected {
			t.Errorf("getPowerLevel(%d,%d) with serial %d = %d, want %d", 
				tt.x, tt.y, tt.serial, result, tt.expected)
		}
	}
}

func TestPart1Examples(t *testing.T) {
	tests := []struct {
		serial   int
		expected string
	}{
		{18, "33,45"},
		{42, "21,61"},
	}
	
	for _, tt := range tests {
		solution := &Solution{serialNumber: tt.serial}
		solution.calculateGrid()
		solution.buildSummedAreaTable()
		
		result, err := solution.Part1()
		if err != nil {
			t.Errorf("Part1() error = %v", err)
			continue
		}
		
		if result != tt.expected {
			t.Errorf("Part1() with serial %d = %s, want %s", tt.serial, result, tt.expected)
		}
	}
}

func TestPart2Examples(t *testing.T) {
	tests := []struct {
		serial   int
		expected string
	}{
		{18, "90,269,16"},
		{42, "232,251,12"},
	}
	
	for _, tt := range tests {
		solution := &Solution{serialNumber: tt.serial}
		solution.calculateGrid()
		solution.buildSummedAreaTable()
		
		result, err := solution.Part2()
		if err != nil {
			t.Errorf("Part2() error = %v", err)
			continue
		}
		
		if result != tt.expected {
			t.Errorf("Part2() with serial %d = %s, want %s", tt.serial, result, tt.expected)
		}
	}
}

func TestPart1(t *testing.T) {
	solution := New(readInput(t))
	result, err := solution.Part1()
	
	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}
	
	expected := "243,68"
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
	
	expected := "236,252,12"
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	} else {
		t.Logf("Part2() = %v", result)
	}
}