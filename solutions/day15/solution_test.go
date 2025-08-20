package day15

import (
	"os"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	// Example from puzzle description
	input1 := `#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`

	solution := New(input1)
	result, err := solution.Part1()

	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}

	expected := 27730 // 47 rounds × 590 HP
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart1AdditionalExamples(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			"Example 2",
			`#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`,
			36334, // 37 rounds × 982 HP
		},
		{
			"Example 3",
			`#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`,
			39514, // 46 rounds × 859 HP
		},
		{
			"Example 4",
			`#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`,
			27755, // 35 rounds × 793 HP
		},
		{
			"Example 5",
			`#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######`,
			28944, // 54 rounds × 536 HP
		},
		{
			"Example 6",
			`#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########`,
			18740, // 20 rounds × 937 HP
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
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
		name     string
		input    string
		expected int
	}{
		{
			"Example 1",
			`#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`,
			4988, // 15 attack power, 29 rounds × 172 HP
		},
		{
			"Example 2",
			`#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`,
			31284, // 4 attack power, 33 rounds × 948 HP
		},
		{
			"Example 3",
			`#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`,
			3478, // 15 attack power, 37 rounds × 94 HP
		},
		{
			"Example 4",
			`#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######`,
			6474, // 12 attack power, 39 rounds × 166 HP
		},
		{
			"Example 5",
			`#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########`,
			1140, // 34 attack power, 30 rounds × 38 HP
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
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

	expected := 224370 // Confirmed correct
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

	expected := 45539 // Confirmed correct
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}
