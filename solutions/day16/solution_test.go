package day16

import (
	"os"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input := `Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]`

	solution := New(input)

	// Parse the sample
	samples, _, err := solution.parseInput()
	if err != nil {
		t.Fatalf("Failed to parse input: %v", err)
	}
	if len(samples) != 1 {
		t.Fatalf("Expected 1 sample, got %d", len(samples))
	}

	sample := samples[0]
	matches := solution.countMatchingOpcodes(sample)

	// According to the problem, this sample behaves like 3 opcodes
	expected := 3
	if matches != expected {
		t.Errorf("Expected %d matching opcodes, got %d", expected, matches)
	}
}

func TestExecuteOpcodes(t *testing.T) {
	solution := New("")

	tests := []struct {
		name      string
		opcode    string
		registers [4]int
		A, B, C   int
		expected  [4]int
	}{
		{
			name:      "addr",
			opcode:    "addr",
			registers: [4]int{3, 2, 1, 1},
			A:         0, B: 1, C: 2,
			expected: [4]int{3, 2, 5, 1}, // 3 + 2 = 5
		},
		{
			name:      "addi",
			opcode:    "addi",
			registers: [4]int{3, 2, 1, 1},
			A:         0, B: 7, C: 3,
			expected: [4]int{3, 2, 1, 10}, // 3 + 7 = 10
		},
		{
			name:      "mulr",
			opcode:    "mulr",
			registers: [4]int{3, 2, 1, 1},
			A:         2, B: 1, C: 2,
			expected: [4]int{3, 2, 2, 1}, // 1 * 2 = 2
		},
		{
			name:      "seti",
			opcode:    "seti",
			registers: [4]int{3, 2, 1, 1},
			A:         2, B: 1, C: 2,
			expected: [4]int{3, 2, 2, 1}, // set 2 to register 2
		},
		{
			name:      "gtir",
			opcode:    "gtir",
			registers: [4]int{1, 2, 3, 0},
			A:         4, B: 2, C: 3,
			expected: [4]int{1, 2, 3, 1}, // 4 > 3 = true (1)
		},
		{
			name:      "eqrr",
			opcode:    "eqrr",
			registers: [4]int{1, 2, 2, 0},
			A:         1, B: 2, C: 3,
			expected: [4]int{1, 2, 2, 1}, // 2 == 2 = true (1)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registers := tt.registers
			solution.executeOpcode(tt.opcode, tt.A, tt.B, tt.C, &registers)

			if registers != tt.expected {
				t.Errorf("Expected registers %v, got %v", tt.expected, registers)
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

	expected := 636
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

	expected := 674
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}

	t.Logf("Part2() = %v", result)
}
