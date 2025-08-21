package day23

import (
	"os"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input := `pos=<0,0,0>, r=4
pos=<1,0,0>, r=1
pos=<4,0,0>, r=3
pos=<0,2,0>, r=1
pos=<0,5,0>, r=3
pos=<0,0,3>, r=1
pos=<1,1,1>, r=1
pos=<1,1,2>, r=1
pos=<1,3,1>, r=1`

	result, err := Part1(input)
	if err != nil {
		t.Fatalf("Part1 failed: %v", err)
	}

	expected := "7"
	if result != expected {
		t.Errorf("Part1 Example = %s; expected %s", result, expected)
	}
}

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("Failed to read input: %v", err)
	}

	result, err := Part1(string(input))
	if err != nil {
		t.Fatalf("Part1 failed: %v", err)
	}

	t.Logf("Part1 Result: %s", result)
}

func TestPart2Example(t *testing.T) {
	input := `pos=<10,12,12>, r=2
pos=<12,14,12>, r=2
pos=<16,12,12>, r=4
pos=<14,14,14>, r=6
pos=<50,50,50>, r=200
pos=<10,10,10>, r=5`

	result, err := Part2(input)
	if err != nil {
		t.Fatalf("Part2 failed: %v", err)
	}

	expected := "36"
	if result != expected {
		t.Errorf("Part2 Example = %s; expected %s", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("Failed to read input: %v", err)
	}

	result, err := Part2(string(input))
	if err != nil {
		t.Fatalf("Part2 failed: %v", err)
	}

	t.Logf("Part2 Result: %s", result)
}