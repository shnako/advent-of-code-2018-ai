package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/shnako/advent-of-code-2018-ai/solutions/day22"
)

func main() {
	var inputPath string
	flag.StringVar(&inputPath, "input", "", "path to input.txt (default: auto-detect)")
	flag.Parse()

	candidates := []string{}
	if inputPath != "" {
		candidates = []string{inputPath}
	} else {
		// Try common run modes: from day dir, from cmd dir, from repo root
		candidates = []string{"input.txt", "../input.txt", "solutions/day22/input.txt"}
	}

	var input []byte
	var err error
	for _, p := range candidates {
		if b, e := os.ReadFile(p); e == nil {
			input, err = b, nil
			break
		} else {
			err = e
		}
	}
	if err != nil {
		log.Fatalf("Failed to read input. Tried: %v. Last error: %v", candidates, err)
	}

	solution := day22.New(string(input))

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