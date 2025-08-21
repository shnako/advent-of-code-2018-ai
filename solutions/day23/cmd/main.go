package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/shnako/advent-of-code-2018-ai/solutions/day23"
)

func main() {
	var inputPath string
	flag.StringVar(&inputPath, "input", "", "path to input.txt (default: auto-detect)")
	flag.Parse()

	candidates := []string{}
	if inputPath != "" {
		candidates = []string{inputPath}
	} else {
		candidates = []string{"input.txt", "../input.txt", "solutions/day23/input.txt"}
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

	part1, err := day23.Part1(string(input))
	if err != nil {
		log.Fatalf("Part1 failed: %v", err)
	}
	fmt.Printf("Part 1: %s\n", part1)

	part2, err := day23.Part2(string(input))
	if err != nil {
		log.Fatalf("Part2 failed: %v", err)
	}
	fmt.Printf("Part 2: %s\n", part2)
}