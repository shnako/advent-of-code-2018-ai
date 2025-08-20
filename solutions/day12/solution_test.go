package day12

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func readInput(t *testing.T) string {
	t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("runtime.Caller failed")
	}
	path := filepath.Join(filepath.Dir(filename), "input.txt")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read input at %s: %v", path, err)
	}
	return string(b)
}

func TestPart1Example(t *testing.T) {
	input := `initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`

	solution, err := New(input)
	if err != nil {
		t.Fatalf("Failed to create solution: %v", err)
	}

	result, err := solution.Part1()
	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}

	expected := "325"
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart1(t *testing.T) {
	solution, err := New(readInput(t))
	if err != nil {
		t.Fatalf("Failed to create solution: %v", err)
	}
	result, err := solution.Part1()

	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}

	expected := "3217"
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	} else {
		t.Logf("Part1() = %v", result)
	}
}

func TestPart2(t *testing.T) {
	solution, err := New(readInput(t))
	if err != nil {
		t.Fatalf("Failed to create solution: %v", err)
	}
	result, err := solution.Part2()

	if err != nil {
		t.Errorf("Part2() error = %v", err)
		return
	}

	expected := "4000000000866"
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	} else {
		t.Logf("Part2() = %v", result)
	}
}

func TestSimulateGenerations(t *testing.T) {
	input := `initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`

	solution, err := New(input)
	if err != nil {
		t.Fatalf("Failed to create solution: %v", err)
	}

	// Test only the final result that we know is correct
	result := solution.simulate(20)
	expected := 325
	if result != expected {
		t.Errorf("simulate(20) = %v, want %v", result, expected)
	}
}