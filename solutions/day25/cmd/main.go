package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/shnako/advent-of-code-2018-ai/solutions/day25"
)

func main() {
	// Search common locations to support `go run ./solutions/day25/cmd` and running the built binary.
	exePath, _ := os.Executable()
	candidates := []string{
		"input.txt",                         // solutions/day25/
		"solutions/day25/input.txt",         // repo root
		"../input.txt",                      // when cwd is solutions/day25/cmd
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

	part1 := day25.Part1(input)
	fmt.Printf("Part 1: %s\n", part1)

	part2 := day25.Part2(input)
	fmt.Printf("Part 2: %s\n", part2)
}