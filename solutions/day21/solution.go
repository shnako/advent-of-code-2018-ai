/*
 * Day 21: Chronal Conversion
 * 
 * Part 1: Find the lowest non-negative integer for register 0 that causes the program to halt.
 * The program halts when instruction 30 (eqrr 3 0 5) is true, meaning register 3 equals register 0.
 * We simulate the program to find the first value that register 3 takes when reaching instruction 30.
 * 
 * Part 2: Find the value for register 0 that causes the program to halt after executing the most instructions.
 * We track all values that register 3 takes at instruction 30 until we find a cycle.
 * The last unique value before the cycle starts is the answer.
 */

package day21

import (
	"strconv"
	"strings"
)

type Solution struct {
	input string
}

func New(input string) *Solution {
	return &Solution{input: strings.TrimSpace(input)}
}

type Instruction struct {
	op string
	a  int
	b  int
	c  int
}

func (s *Solution) Part1() (int, error) {
	lines := strings.Split(s.input, "\n")
	ipReg := s.parseIPRegister(lines[0])
	instructions := s.parseInstructions(lines[1:])
	
	// Find the first value that r3 takes when reaching instruction 30
	registers := make([]int, 6)
	ip := 0
	
	for ip >= 0 && ip < len(instructions) {
		if ip == 30 {
			// This is where the program checks if r3 == r0
			// The first value r3 takes here is our answer for Part 1
			return registers[3], nil
		}
		
		registers[ipReg] = ip
		s.execute(instructions[ip], registers)
		ip = registers[ipReg]
		ip++
	}
	
	return 0, nil
}

func (s *Solution) Part2() (int, error) {
	// Instead of simulating the entire program, we can optimize by understanding what it does
	// The program generates a sequence of values in r3 and checks if r3 == r0
	// We need to find the last unique value before the sequence repeats
	
	seen := make(map[int]bool)
	lastUnique := 0
	r3 := 0
	
	for {
		r4 := r3 | 65536
		r3 = 7041048
		
		for {
			r3 = (((r3 + (r4 & 255)) & 16777215) * 65899) & 16777215
			
			if r4 < 256 {
				break
			}
			r4 = r4 / 256
		}
		
		if seen[r3] {
			// We've seen this value before - cycle detected
			return lastUnique, nil
		}
		seen[r3] = true
		lastUnique = r3
	}
}

func (s *Solution) parseIPRegister(line string) int {
	parts := strings.Split(line, " ")
	ip, _ := strconv.Atoi(parts[1])
	return ip
}

func (s *Solution) parseInstructions(lines []string) []Instruction {
	instructions := make([]Instruction, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		a, _ := strconv.Atoi(parts[1])
		b, _ := strconv.Atoi(parts[2])
		c, _ := strconv.Atoi(parts[3])
		instructions = append(instructions, Instruction{
			op: parts[0],
			a:  a,
			b:  b,
			c:  c,
		})
	}
	return instructions
}

func (s *Solution) execute(inst Instruction, registers []int) {
	switch inst.op {
	case "addr":
		registers[inst.c] = registers[inst.a] + registers[inst.b]
	case "addi":
		registers[inst.c] = registers[inst.a] + inst.b
	case "mulr":
		registers[inst.c] = registers[inst.a] * registers[inst.b]
	case "muli":
		registers[inst.c] = registers[inst.a] * inst.b
	case "banr":
		registers[inst.c] = registers[inst.a] & registers[inst.b]
	case "bani":
		registers[inst.c] = registers[inst.a] & inst.b
	case "borr":
		registers[inst.c] = registers[inst.a] | registers[inst.b]
	case "bori":
		registers[inst.c] = registers[inst.a] | inst.b
	case "setr":
		registers[inst.c] = registers[inst.a]
	case "seti":
		registers[inst.c] = inst.a
	case "gtir":
		if inst.a > registers[inst.b] {
			registers[inst.c] = 1
		} else {
			registers[inst.c] = 0
		}
	case "gtri":
		if registers[inst.a] > inst.b {
			registers[inst.c] = 1
		} else {
			registers[inst.c] = 0
		}
	case "gtrr":
		if registers[inst.a] > registers[inst.b] {
			registers[inst.c] = 1
		} else {
			registers[inst.c] = 0
		}
	case "eqir":
		if inst.a == registers[inst.b] {
			registers[inst.c] = 1
		} else {
			registers[inst.c] = 0
		}
	case "eqri":
		if registers[inst.a] == inst.b {
			registers[inst.c] = 1
		} else {
			registers[inst.c] = 0
		}
	case "eqrr":
		if registers[inst.a] == registers[inst.b] {
			registers[inst.c] = 1
		} else {
			registers[inst.c] = 0
		}
	}
}