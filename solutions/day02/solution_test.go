package day02

import (
	"os"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	input := `abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`
	
	solution := New(input)
	result, err := solution.Part1()
	
	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}
	
	expected := 12 // 4 * 3 = 12
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart2Examples(t *testing.T) {
	input := `abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`
	
	solution := New(input)
	result, err := solution.Part2()
	
	if err != nil {
		t.Errorf("Part2() error = %v", err)
		return
	}
	
	expected := "fgij"
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
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
	
	expected := 5390
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
	
	expected := "nvosmkcdtdbfhyxsphzgraljq"
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}