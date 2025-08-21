package day22

import (
	"fmt"
	"log"
	"os"

)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	solution := New(string(input))

	part1, err := solution.Part1()
	if err != nil {
		log.Fatalf("Part 1 failed: %v", err)
	}
	fmt.Printf("Part 1: %d\n", part1)

	part2, err := solution.Part2()
	if err != nil {
		log.Fatalf("Part 2 failed: %v", err)
	}
	fmt.Printf("Part 2: %d\n", part2)
}