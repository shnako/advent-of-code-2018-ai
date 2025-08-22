package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/shnako/advent-of-code-2018-ai/solutions/day15"
)

func main() {
	// Search common locations to support `go run ./solutions/day15/cmd` and running the built binary.
	exePath, _ := os.Executable()
	candidates := []string{
		"input.txt",                         // solutions/day15/
		"solutions/day15/input.txt",         // repo root
		"../input.txt",                      // when cwd is solutions/day15/cmd
		filepath.Join(filepath.Dir(exePath), "input.txt"),
		filepath.Join(filepath.Dir(exePath), "..", "input.txt"),
	}
	var data []byte
	var err error
	for _, p := range candidates {
		if b, e := os.ReadFile(p); e == nil {
			data, err = b, nil
			break
		} else {
			err = e
		}
	}
	if data == nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	input := strings.TrimSpace(string(data))
	solution := day15.New(input)

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
