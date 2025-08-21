package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/shnako/advent-of-code-2018-ai/solutions/day24"
)

func main() {
	// Try to find input.txt in the parent directory first (when run from cmd/)
	inputPath := "../input.txt"
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		// Try current directory (when run from solutions/day24/)
		inputPath = "input.txt"
	}

	data, err := os.ReadFile(inputPath)
	if err != nil {
		// Try to find it relative to the executable
		exePath, _ := os.Executable()
		inputPath = filepath.Join(filepath.Dir(exePath), "..", "input.txt")
		data, err = os.ReadFile(inputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input.txt: %v\n", err)
			os.Exit(1)
		}
	}

	input := strings.TrimSpace(string(data))

	part1 := day24.Part1(input)
	fmt.Printf("Part 1: %s\n", part1)

	part2 := day24.Part2(input)
	fmt.Printf("Part 2: %s\n", part2)
}