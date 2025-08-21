package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/shnako/advent-of-code-2018-ai/solutions/day24"
)

func main() {
	// Search common locations to support `go run ./solutions/day24/cmd` and running the built binary.
	exePath, _ := os.Executable()
	candidates := []string{
		"input.txt",                         // solutions/day24/
		"solutions/day24/input.txt",         // repo root
		"../input.txt",                      // when cwd is solutions/day24/cmd
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
		fmt.Fprintf(os.Stderr, "Error reading input.txt: %v\n", err)
		os.Exit(1)
	}

	input := strings.TrimSpace(string(data))

	part1 := day24.Part1(input)
	fmt.Printf("Part 1: %s\n", part1)

	part2 := day24.Part2(input)
	fmt.Printf("Part 2: %s\n", part2)
}