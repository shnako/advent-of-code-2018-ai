/*
 * Day 16: Chronal Classification
 * 
 * Part 1: Analyze instruction samples to determine how many behave like 3+ opcodes.
 * Each sample shows register state before/after an instruction execution.
 * Test each of 16 possible opcodes against the sample to see which ones match.
 * 
 * Part 2: Use the samples to deduce which opcode number corresponds to which operation,
 * then execute the test program using the decoded opcodes.
 */

package day16

import (
	"strconv"
	"strings"
)

type Sample struct {
	Before      [4]int
	Instruction [4]int // opcode, A, B, C
	After       [4]int
}

type Solution struct {
	input string
}

func New(input string) *Solution {
	return &Solution{input: strings.TrimSpace(input)}
}

func (s *Solution) Part1() (int, error) {
	samples, _ := s.parseInput()
	
	count := 0
	for _, sample := range samples {
		matches := s.countMatchingOpcodes(sample)
		if matches >= 3 {
			count++
		}
	}
	
	return count, nil
}

func (s *Solution) Part2() (int, error) {
	samples, program := s.parseInput()
	
	// Determine which opcode number corresponds to which operation
	opcodeMapping := s.deduceOpcodes(samples)
	
	// Execute the test program
	registers := [4]int{0, 0, 0, 0}
	for _, instruction := range program {
		opcode := opcodeMapping[instruction[0]]
		s.executeOpcode(opcode, instruction[1], instruction[2], instruction[3], &registers)
	}
	
	return registers[0], nil
}

func (s *Solution) parseInput() ([]Sample, [][4]int) {
	// Normalize line endings first
	normalizedInput := strings.ReplaceAll(s.input, "\r\n", "\n")
	
	parts := strings.Split(normalizedInput, "\n\n\n")
	if len(parts) < 2 {
		// Try different separators in case of formatting issues
		parts = strings.Split(normalizedInput, "\n\n\n\n")
	}
	
	sampleLines := strings.Split(strings.TrimSpace(parts[0]), "\n")
	var samples []Sample
	
	for i := 0; i < len(sampleLines); i += 4 {
		if i+2 >= len(sampleLines) {
			break
		}
		
		sample := Sample{}
		
		// Parse Before line: "Before: [3, 1, 0, 1]"
		beforeStr := strings.TrimPrefix(sampleLines[i], "Before: [")
		beforeStr = strings.TrimSuffix(beforeStr, "]")
		beforeParts := strings.Split(beforeStr, ", ")
		for j, part := range beforeParts {
			sample.Before[j], _ = strconv.Atoi(strings.TrimSpace(part))
		}
		
		// Parse instruction line: "9 3 3 2"
		instParts := strings.Fields(sampleLines[i+1])
		for j, part := range instParts {
			sample.Instruction[j], _ = strconv.Atoi(part)
		}
		
		// Parse After line: "After:  [3, 1, 0, 1]"
		afterStr := strings.TrimPrefix(sampleLines[i+2], "After:  [")
		afterStr = strings.TrimSuffix(afterStr, "]")
		afterParts := strings.Split(afterStr, ", ")
		for j, part := range afterParts {
			sample.After[j], _ = strconv.Atoi(strings.TrimSpace(part))
		}
		
		samples = append(samples, sample)
	}
	
	// Parse program (second part)
	var program [][4]int
	if len(parts) >= 2 {
		programLines := strings.Split(strings.TrimSpace(parts[1]), "\n")
		for _, line := range programLines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			parts := strings.Fields(line)
			if len(parts) == 4 {
				instruction := [4]int{}
				for i, part := range parts {
					instruction[i], _ = strconv.Atoi(part)
				}
				program = append(program, instruction)
			}
		}
	}
	
	return samples, program
}

func (s *Solution) countMatchingOpcodes(sample Sample) int {
	opcodes := []string{
		"addr", "addi", "mulr", "muli",
		"banr", "bani", "borr", "bori",
		"setr", "seti", "gtir", "gtri",
		"gtrr", "eqir", "eqri", "eqrr",
	}
	
	count := 0
	for _, opcode := range opcodes {
		registers := sample.Before
		s.executeOpcode(opcode, sample.Instruction[1], sample.Instruction[2], sample.Instruction[3], &registers)
		if registers == sample.After {
			count++
		}
	}
	
	return count
}

func (s *Solution) deduceOpcodes(samples []Sample) map[int]string {
	opcodes := []string{
		"addr", "addi", "mulr", "muli",
		"banr", "bani", "borr", "bori",
		"setr", "seti", "gtir", "gtri",
		"gtrr", "eqir", "eqri", "eqrr",
	}
	
	// For each opcode number, track which operations are still possible
	possible := make(map[int]map[string]bool)
	for i := 0; i < 16; i++ {
		possible[i] = make(map[string]bool)
		for _, op := range opcodes {
			possible[i][op] = true
		}
	}
	
	// Eliminate impossible combinations based on samples
	for _, sample := range samples {
		opcodeNum := sample.Instruction[0]
		for _, opcode := range opcodes {
			registers := sample.Before
			s.executeOpcode(opcode, sample.Instruction[1], sample.Instruction[2], sample.Instruction[3], &registers)
			if registers != sample.After {
				possible[opcodeNum][opcode] = false
			}
		}
	}
	
	// Use process of elimination to find unique mappings
	mapping := make(map[int]string)
	used := make(map[string]bool)
	
	// Safety counter to prevent infinite loops
	maxIterations := 100
	iterations := 0
	
	for len(mapping) < 16 && iterations < maxIterations {
		iterations++
		progress := false
		
		// Find opcode numbers with only one possible operation
		for opcodeNum := 0; opcodeNum < 16; opcodeNum++ {
			if _, found := mapping[opcodeNum]; found {
				continue
			}
			
			var candidates []string
			for op, stillPossible := range possible[opcodeNum] {
				if stillPossible && !used[op] {
					candidates = append(candidates, op)
				}
			}
			
			if len(candidates) == 1 {
				mapping[opcodeNum] = candidates[0]
				used[candidates[0]] = true
				progress = true
			}
		}
		
		// If no progress was made, we might be stuck
		if !progress {
			break
		}
	}
	
	return mapping
}

func (s *Solution) executeOpcode(opcode string, A, B, C int, registers *[4]int) {
	switch opcode {
	case "addr":
		registers[C] = registers[A] + registers[B]
	case "addi":
		registers[C] = registers[A] + B
	case "mulr":
		registers[C] = registers[A] * registers[B]
	case "muli":
		registers[C] = registers[A] * B
	case "banr":
		registers[C] = registers[A] & registers[B]
	case "bani":
		registers[C] = registers[A] & B
	case "borr":
		registers[C] = registers[A] | registers[B]
	case "bori":
		registers[C] = registers[A] | B
	case "setr":
		registers[C] = registers[A]
	case "seti":
		registers[C] = A
	case "gtir":
		if A > registers[B] {
			registers[C] = 1
		} else {
			registers[C] = 0
		}
	case "gtri":
		if registers[A] > B {
			registers[C] = 1
		} else {
			registers[C] = 0
		}
	case "gtrr":
		if registers[A] > registers[B] {
			registers[C] = 1
		} else {
			registers[C] = 0
		}
	case "eqir":
		if A == registers[B] {
			registers[C] = 1
		} else {
			registers[C] = 0
		}
	case "eqri":
		if registers[A] == B {
			registers[C] = 1
		} else {
			registers[C] = 0
		}
	case "eqrr":
		if registers[A] == registers[B] {
			registers[C] = 1
		} else {
			registers[C] = 0
		}
	}
}