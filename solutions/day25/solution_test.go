package day25

import (
	"os"
	"strings"
	"testing"
)

func TestPart1Example1(t *testing.T) {
	input := ` 0,0,0,0
 3,0,0,0
 0,3,0,0
 0,0,3,0
 0,0,0,3
 0,0,0,6
 9,0,0,0
12,0,0,0`

	expected := "2"
	result := Part1(input)
	if result != expected {
		t.Errorf("Part1 with example 1 = %s; want %s", result, expected)
	}
}

func TestPart1Example2(t *testing.T) {
	input := `-1,2,2,0
0,0,2,-2
0,0,0,-2
-1,2,0,0
-2,-2,-2,2
3,0,2,-1
-1,3,2,2
-1,0,-1,0
0,2,1,-2
3,0,0,0`

	expected := "4"
	result := Part1(input)
	if result != expected {
		t.Errorf("Part1 with example 2 = %s; want %s", result, expected)
	}
}

func TestPart1Example3(t *testing.T) {
	input := `1,-1,0,1
2,0,-1,0
3,2,-1,0
0,0,3,1
0,0,-1,-1
2,3,-2,0
-2,2,0,0
2,-2,0,-1
1,-1,0,-1
3,2,0,2`

	expected := "3"
	result := Part1(input)
	if result != expected {
		t.Errorf("Part1 with example 3 = %s; want %s", result, expected)
	}
}

func TestPart1Example4(t *testing.T) {
	input := `1,-1,-1,-2
-2,-2,0,1
0,2,1,3
-2,3,-2,1
0,2,3,-2
-1,-1,1,-2
0,-2,-1,0
-2,2,3,-1
1,2,2,0
-1,-2,0,-2`

	expected := "8"
	result := Part1(input)
	if result != expected {
		t.Errorf("Part1 with example 4 = %s; want %s", result, expected)
	}
}

func TestPart1(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("Failed to read input.txt: %v", err)
	}
	input := strings.TrimSpace(string(data))
	
	got := Part1(input)
	want := "396"
	if got != want {
		t.Fatalf("Part1(input.txt) = %s; want %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	// Day 25 has no Part 2
	result := Part2("")
	expected := "Merry Christmas!"
	if result != expected {
		t.Errorf("Part2 = %s; want %s", result, expected)
	}
}