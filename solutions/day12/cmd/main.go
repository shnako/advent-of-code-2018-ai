package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shnako/advent-of-code-2018-ai/solutions/day12"
)

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	solution, err := day12.New(string(input))
	if err != nil {
		log.Fatalf("Init failed: %v", err)
	}

	part1, err := solution.Part1()
	if err != nil {
		log.Fatalf("Part 1 failed: %v", err)
	}
	fmt.Printf("Part 1: %s\n", part1)

	part2, err := solution.Part2()
	if err != nil {
		log.Fatalf("Part 2 failed: %v", err)
	}
	fmt.Printf("Part 2: %s\n", part2)
}
