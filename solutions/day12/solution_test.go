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

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name: "example",
			input: `initial state: #..#.#..##......###...###

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
####. => #`,
			want: "325",
		},
		{
			name:  "input",
			input: readInput(t),
			want:  "3217",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			solution, err := New(tt.input)
			if err != nil {
				t.Fatalf("Failed to create solution: %v", err)
			}
			got, err := solution.Part1()
			if err != nil {
				t.Fatalf("Part1() error = %v", err)
			}
			if got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()
	solution, err := New(readInput(t))
	if err != nil {
		t.Fatalf("Failed to create solution: %v", err)
	}
	result, err := solution.Part2()

	if err != nil {
		t.Fatalf("Part2() error = %v", err)
	}

	expected := "4000000000866"
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	} else {
		t.Logf("Part2() = %v", result)
	}
}

func TestSimulateGenerations(t *testing.T) {
	t.Parallel()
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
