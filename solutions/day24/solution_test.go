package day24

import (
	"os"
	"strings"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input := `Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4`

	expected := "5216"
	result := Part1(input)
	if result != expected {
		t.Errorf("Part1 with example input = %s; want %s", result, expected)
	}
}

func TestPart1(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("Failed to read input.txt: %v", err)
	}
	input := strings.TrimSpace(string(data))
	
	result := Part1(input)
	t.Logf("Part 1 Result: %s", result)
}

func TestPart2Example(t *testing.T) {
	input := `Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4`

	expected := "51"
	result := Part2(input)
	if result != expected {
		t.Errorf("Part2 with example input = %s; want %s", result, expected)
	}
}

func TestPart2(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("Failed to read input.txt: %v", err)
	}
	input := strings.TrimSpace(string(data))
	
	result := Part2(input)
	t.Logf("Part 2 Result: %s", result)
}