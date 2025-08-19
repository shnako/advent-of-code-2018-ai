# Advent of Code 2018 - Go Solutions

The solutions to this are fully implemented by AI as an experiment.

This repository contains solutions for the [Advent of Code 2018](https://adventofcode.com/2018) challenges, implemented in Go.

## Progress

| Day | Puzzle | Solution | Tests | Input | Description |
|-----|--------|----------|-------|-------|-------------|
| 01  | [Day 01](https://adventofcode.com/2018/day/1) | - | - | - | - |
| 02  | [Day 02](https://adventofcode.com/2018/day/2) | - | - | - | - |
| 03  | [Day 03](https://adventofcode.com/2018/day/3) | - | - | - | - |
| 04  | [Day 04](https://adventofcode.com/2018/day/4) | - | - | - | - |
| 05  | [Day 05](https://adventofcode.com/2018/day/5) | - | - | - | - |
| 06  | [Day 06](https://adventofcode.com/2018/day/6) | - | - | - | - |
| 07  | [Day 07](https://adventofcode.com/2018/day/7) | - | - | - | - |
| 08  | [Day 08](https://adventofcode.com/2018/day/8) | - | - | - | - |
| 09  | [Day 09](https://adventofcode.com/2018/day/9) | - | - | - | - |
| 10  | [Day 10](https://adventofcode.com/2018/day/10) | - | - | - | - |
| 11  | [Day 11](https://adventofcode.com/2018/day/11) | - | - | - | - |
| 12  | [Day 12](https://adventofcode.com/2018/day/12) | - | - | - | - |
| 13  | [Day 13](https://adventofcode.com/2018/day/13) | - | - | - | - |
| 14  | [Day 14](https://adventofcode.com/2018/day/14) | - | - | - | - |
| 15  | [Day 15](https://adventofcode.com/2018/day/15) | - | - | - | - |
| 16  | [Day 16](https://adventofcode.com/2018/day/16) | - | - | - | - |
| 17  | [Day 17](https://adventofcode.com/2018/day/17) | - | - | - | - |
| 18  | [Day 18](https://adventofcode.com/2018/day/18) | - | - | - | - |
| 19  | [Day 19](https://adventofcode.com/2018/day/19) | - | - | - | - |
| 20  | [Day 20](https://adventofcode.com/2018/day/20) | - | - | - | - |
| 21  | [Day 21](https://adventofcode.com/2018/day/21) | - | - | - | - |
| 22  | [Day 22](https://adventofcode.com/2018/day/22) | - | - | - | - |
| 23  | [Day 23](https://adventofcode.com/2018/day/23) | - | - | - | - |
| 24  | [Day 24](https://adventofcode.com/2018/day/24) | - | - | - | - |
| 25  | [Day 25](https://adventofcode.com/2018/day/25) | - | - | - | - |

## Setup

1. Set your AOC session cookie as an environment variable:
   ```bash
   export AOC_SESSION_COOKIE="your_session_cookie_here"
   ```

2. Fetch a puzzle and input:
   ```bash
   go run cmd/fetch/main.go -day=1
   ```

3. Implement the solution in `solutions/dayXX/solution.go`

4. Run the solution:
   ```bash
   cd solutions/day01
   go run .
   ```

5. Submit your answer:
   ```bash
   go run cmd/submit/main.go -day=1 -part=1 -answer=YOUR_ANSWER
   ```

## Project Structure

```
.
├── cmd/
│   ├── fetch/      # Fetches puzzle descriptions and inputs
│   └── submit/     # Submits answers to Advent of Code
├── internal/
│   └── utils/      # Shared utility functions
│       ├── input.go   # Input parsing utilities
│       ├── math.go    # Math utilities
│       ├── grid.go    # 2D grid utilities
│       └── graph.go   # Graph algorithms
└── solutions/
    └── dayXX/      # Solutions for each day
        ├── solution.go      # Implementation
        ├── solution_test.go # Tests
        ├── main.go          # Runner
        ├── input.txt        # Puzzle input
        └── puzzle.txt       # Problem description
```

## Running Tests

Run tests for a specific day:
```bash
cd solutions/day01
go test
```

Run all tests:
```bash
go test ./...
```
