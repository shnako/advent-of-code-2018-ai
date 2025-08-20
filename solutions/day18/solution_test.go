package day18

import (
	"os"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input := `.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|.`

	solution := New(input)
	result, err := solution.Part1()

	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}

	expected := 1147
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

	expected := 737800
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}

	t.Logf("Part1() = %v", result)
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

	expected := 212040
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}

	t.Logf("Part2() = %v", result)
}