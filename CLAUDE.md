# Advent of Code 2018 Go Assistant Instructions

You are helping solve Advent of Code 2018 puzzles using Go. This is a learning experience as the user has no Go experience. Follow these guidelines:

## Sequential Solving Workflow

When the user says "Let's start solving" or similar, and the AOC_SESSION_COOKIE environment variable is set, solve ALL 25 days sequentially:

### For Each Day (1-25):
1. **Fetch the puzzle**: Run `go run cmd/fetch/main.go -day=X`
   - This fetches both the puzzle description AND your personal input from:
   - Puzzle: `https://adventofcode.com/2018/day/X`
   - Input: `https://adventofcode.com/2018/day/X/input`
2. **Read and understand**: Read the puzzle.txt to understand the problem
3. **Create feature branch**: Create a new branch for the day's work (e.g., `feature/day-X`)
4. **Implement Part 1**: Write the solution in solution.go
5. **Test Part 1**: Run with the example from puzzle.txt using `go test`
6. **Run Part 1**: Execute with the full input.txt using `go run .`
7. **Submit Part 1**: Immediately submit using `go run cmd/submit/main.go -day=X -part=1 -answer=Y`
8. **Verify submission**: Check if the answer was correct
9. **If correct, fetch Part 2**: The puzzle.txt will be updated automatically when Part 1 is solved
10. **Implement Part 2**: Update solution.go with Part 2 solution
11. **Test Part 2**: Run with any Part 2 examples
12. **Run Part 2**: Execute with the full input
13. **Submit Part 2**: Immediately submit the Part 2 answer
14. **Run final tests**: Ensure all tests pass with `go test`
15. **Update progress**: Mark day as complete in progress tracking
16. **Create Pull Request**: Create PR with comprehensive description of the solution
17. **CodeRabbit Review Cycle**: 
    - Wait for CodeRabbit to provide automated PR comments
    - Address ALL CodeRabbit comments by making necessary code changes
    - Push changes to the feature branch
    - Reply to CodeRabbit conversations explaining the changes made
    - Resolve conversations that have been adequately addressed
    - Wait for CodeRabbit to re-review and provide additional comments
    - Repeat this cycle until CodeRabbit is satisfied (no more substantive comments)
18. **Merge PR**: Once CodeRabbit review is complete and approved, merge the PR
19. **VERIFY COMPLIANCE**: Double-check that ALL steps in this CLAUDE.md workflow have been followed:
    - Did I fetch both puzzle and input?
    - Did I implement and test both parts?
    - Did I submit both answers immediately?
    - Did I update the README progress table?
    - Did I run all tests successfully?
    - Is the code clean, idiomatic Go with proper error handling?
    - Have I followed all Go best practices outlined in this document?
    - Did I complete the full CodeRabbit review cycle?
    - Did I merge the PR successfully?
20. **COMPACT CONTEXT**: After completing each day, use the `/compact` command to reduce context size
21. **Continue to next day**: Automatically proceed to the next day without stopping

### Important for Automation:
- **Never stop between days**: Continue solving until all 25 days are complete
- **The fetch command downloads BOTH puzzle and input** using the session cookie
- **input.txt contains your personal puzzle input** (different for each user)
- **Always submit immediately** when you have an answer (don't wait for user confirmation)
- **Continue to Part 2 automatically** after Part 1 is confirmed correct
- **Handle wrong answers**: If an answer is wrong, debug and try again
- **Complete all 25 days**: Don't stop until all puzzles are solved
- **Context Management**: After completing each day, use `/compact` to minimize context before proceeding
- **Quality Assurance**: Step 15 MUST be completed - verify full compliance with this document before moving on

### Handling Rate Limits:
If the user's daily allowance is exceeded:
1. **Save Progress**: Note which day/part was last completed in the conversation
2. **User Action**: User should close the session and return when limits reset
3. **Resume**: When returning, user can use `--resume` flag or say "Continue from Day X Part Y"
4. **State Recovery**: Check the README.md progress table and solutions/ directory to determine where to continue
5. **Continue Solving**: Resume from the exact point where we stopped

## Go Project Structure
```
advent-of-code-2018-go/
├── go.mod                              # Go module file
├── go.sum                              # Go module checksums
├── README.md                           # Main README with progress table
├── CLAUDE-GO.md                        # This file
├── cmd/                                # Command-line tools
│   ├── fetch/
│   │   └── main.go                    # Fetches puzzle and input
│   └── submit/
│       └── main.go                    # Submits answers
├── internal/                           # Internal packages
│   └── utils/
│       ├── input.go                   # Input parsing utilities
│       ├── math.go                    # Math utilities
│       ├── grid.go                    # 2D grid utilities
│       └── graph.go                   # Graph algorithms
└── solutions/
    ├── day01/
    │   ├── solution.go                # Solution implementation
    │   ├── solution_test.go           # Tests with examples and answers
    │   ├── input.txt                  # Puzzle input
    │   └── puzzle.txt                 # Problem description
    ├── day02/
    │   └── ...
    └── day25/
        └── ...
```

## Go Module Setup
Initialize the project as a Go module:
```bash
go mod init github.com/username/advent-of-code-2018-go
```

## Solution Template (solution.go)
```go
/*
 * Day X: [Problem Title]
 * 
 * Part 1: [Clear description of what part 1 does and how it works]
 * [Additional details about the approach, algorithms, or key concepts]
 * 
 * Part 2: [Clear description of what part 2 does and how it differs from part 1]
 * [Additional details about new logic, modifications, or different approaches]
 */

package day01

import (
    "fmt"
    "strings"
)

type Solution struct {
    input string
}

func New(input string) *Solution {
    return &Solution{input: strings.TrimSpace(input)}
}

func (s *Solution) Part1() (int, error) {
    lines := strings.Split(s.input, "\n")
    
    // Implementation here
    
    return 0, nil
}

func (s *Solution) Part2() (int, error) {
    lines := strings.Split(s.input, "\n")
    
    // Implementation here
    
    return 0, nil
}

// Helper functions as needed
func (s *Solution) parseInput() []string {
    return strings.Split(s.input, "\n")
}
```

## Test Template (solution_test.go)
```go
package day01

import (
    "os"
    "testing"
)

func TestPart1Example(t *testing.T) {
    input := `example input here`
    
    solution := New(input)
    result, err := solution.Part1()
    
    if err != nil {
        t.Errorf("Part1() error = %v", err)
        return
    }
    
    expected := 42
    if result != expected {
        t.Errorf("Part1() = %v, want %v", result, expected)
    }
}

func TestPart1(t *testing.T) {
    input, err := os.ReadFile("input.txt")
    if err != nil {
        t.Fatalf("Failed to read input: %v", err)
    }
    
    solution := New(string(input))
    result, err := solution.Part1()
    
    if err != nil {
        t.Errorf("Part1() error = %v", err)
        return
    }
    
    expected := 12345 // Update after solving
    if result != expected {
        t.Errorf("Part1() = %v, want %v", result, expected)
    }
}

func TestPart2(t *testing.T) {
    input, err := os.ReadFile("input.txt")
    if err != nil {
        t.Fatalf("Failed to read input: %v", err)
    }
    
    solution := New(string(input))
    result, err := solution.Part2()
    
    if err != nil {
        t.Errorf("Part2() error = %v", err)
        return
    }
    
    expected := 67890 // Update after solving
    if result != expected {
        t.Errorf("Part2() = %v, want %v", result, expected)
    }
}
```

## Main Runner Template
Create a main.go in each day's folder:
```go
package main

import (
    "fmt"
    "log"
    "os"
    
    "github.com/username/advent-of-code-2018-go/solutions/day01"
)

func main() {
    input, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatalf("Failed to read input: %v", err)
    }
    
    solution := day01.New(string(input))
    
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
```

## Utility Functions (internal/utils/)

### input.go
```go
package utils

import (
    "strconv"
    "strings"
)

func ParseInts(input string) ([]int, error) {
    lines := strings.Split(strings.TrimSpace(input), "\n")
    nums := make([]int, 0, len(lines))
    
    for _, line := range lines {
        n, err := strconv.Atoi(strings.TrimSpace(line))
        if err != nil {
            return nil, err
        }
        nums = append(nums, n)
    }
    
    return nums, nil
}

func SplitLines(input string) []string {
    return strings.Split(strings.TrimSpace(input), "\n")
}
```

### grid.go
```go
package utils

type Point struct {
    X, Y int
}

type Grid struct {
    Width, Height int
    Data          [][]rune
}

func NewGrid(input string) *Grid {
    lines := strings.Split(strings.TrimSpace(input), "\n")
    height := len(lines)
    width := len(lines[0])
    
    data := make([][]rune, height)
    for i, line := range lines {
        data[i] = []rune(line)
    }
    
    return &Grid{
        Width:  width,
        Height: height,
        Data:   data,
    }
}

func (g *Grid) Get(p Point) rune {
    if p.X < 0 || p.X >= g.Width || p.Y < 0 || p.Y >= g.Height {
        return 0
    }
    return g.Data[p.Y][p.X]
}

func (g *Grid) Set(p Point, val rune) {
    if p.X >= 0 && p.X < g.Width && p.Y >= 0 && p.Y < g.Height {
        g.Data[p.Y][p.X] = val
    }
}

// Common directions
var (
    North = Point{0, -1}
    South = Point{0, 1}
    East  = Point{1, 0}
    West  = Point{-1, 0}
    
    Cardinals = []Point{North, South, East, West}
)
```

## Go Best Practices

### Code Style
- Use `gofmt` to format code (automatically done by most editors)
- Follow Go naming conventions:
  - Exported names start with capital letters
  - Use camelCase for local variables
  - Use MixedCaps for exported functions/types
- Keep functions small and focused
- Return errors explicitly, don't panic unless critical
- Use meaningful variable names that match the problem domain

### Performance Considerations
- Use `strings.Builder` for string concatenation in loops
- Prefer `[]byte` over `string` for heavy manipulation
- Pre-allocate slices when size is known: `make([]int, 0, capacity)`
- Use pointers for large structs to avoid copying
- Consider using `sync.Map` for concurrent access
- Profile with `go test -bench` for performance-critical code

### Error Handling
```go
if err != nil {
    return nil, fmt.Errorf("failed to parse input: %w", err)
}
```

### Testing
- Write table-driven tests for multiple cases:
```go
tests := []struct {
    name     string
    input    string
    expected int
}{
    {"example 1", "input1", 42},
    {"example 2", "input2", 84},
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        result := solve(tt.input)
        if result != tt.expected {
            t.Errorf("got %d, want %d", result, tt.expected)
        }
    })
}
```

### Common Patterns

#### BFS (Breadth-First Search)
```go
type State struct {
    pos   Point
    steps int
}

func bfs(start Point, target Point) int {
    queue := []State{{start, 0}}
    visited := make(map[Point]bool)
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        if current.pos == target {
            return current.steps
        }
        
        if visited[current.pos] {
            continue
        }
        visited[current.pos] = true
        
        for _, dir := range Cardinals {
            next := Point{current.pos.X + dir.X, current.pos.Y + dir.Y}
            if !visited[next] {
                queue = append(queue, State{next, current.steps + 1})
            }
        }
    }
    
    return -1
}
```

#### DFS (Depth-First Search)
```go
func dfs(node int, graph map[int][]int, visited map[int]bool) {
    if visited[node] {
        return
    }
    visited[node] = true
    
    for _, neighbor := range graph[node] {
        dfs(neighbor, graph, visited)
    }
}
```

#### Memoization
```go
type memoKey struct {
    param1 int
    param2 string
}

var memo = make(map[memoKey]int)

func solve(param1 int, param2 string) int {
    key := memoKey{param1, param2}
    if val, ok := memo[key]; ok {
        return val
    }
    
    // Calculate result
    result := computeExpensiveOperation(param1, param2)
    
    memo[key] = result
    return result
}
```

## Main README Updates
After solving each day, update the main README.md with:

| Day | Puzzle                                               | Solution                                       | Results                                         | Input                                     | Text                                        |
|-----|------------------------------------------------------|------------------------------------------------|-------------------------------------------------|-------------------------------------------|---------------------------------------------|
| 01  | [Day 01: Title](https://adventofcode.com/2018/day/1) | [Day 01 solution](solutions/day01/solution.go) | [Day 01 test](solutions/day01/solution_test.go) | [Day 01 input](solutions/day01/input.txt) | [Day 01 puzzle](solutions/day01/puzzle.txt) |

Extract the day title from the puzzle.txt file (it's in the format "--- Day X: Title ---").

## Debugging Strategies
- Use `fmt.Printf` for debugging (remove before committing)
- Use Go's debugger: `dlv debug`
- Write small test cases to isolate issues
- Check edge cases: empty input, single element, boundaries
- Verify parsing logic first - many bugs come from input parsing

## Common Problem Patterns (2018 Specific)
Based on Advent of Code 2018:
- **Frequency changes**: Running sums, cycle detection
- **Box IDs**: String manipulation, character counting
- **Fabric claims**: 2D grid overlaps, coordinate systems
- **Guard schedules**: Time parsing, data aggregation
- **Polymer reactions**: Stack-based processing, string transformations
- **Coordinate systems**: Manhattan distance, Voronoi diagrams
- **Tree traversal**: Recursive parsing, tree construction
- **License processing**: Nested structures, metadata handling
- **Marble games**: Circular linked lists, efficient insertion/deletion
- **Power grids**: 2D convolution, dynamic programming
- **Cart tracks**: State machines, collision detection
- **Plant evolution**: Cellular automata, pattern recognition
- **Water flow**: Flood fill, boundary conditions
- **Combat simulation**: Turn-based systems, pathfinding
- **Assembly programs**: Instruction parsing, register manipulation

## Daily Completion Checklist
**MANDATORY**: Before marking any day as complete, verify:
- [ ] Both parts solved and submitted
- [ ] All tests pass with `go test`
- [ ] README.md updated with correct links and status
- [ ] Solution follows Go best practices (error handling, naming conventions)
- [ ] Code is clean without debug statements
- [ ] Solution runs in reasonable time (< 5 seconds)
- [ ] Feature branch created for the day's work
- [ ] Pull Request created with comprehensive description
- [ ] CodeRabbit review cycle completed successfully
- [ ] All CodeRabbit comments addressed and resolved
- [ ] Pull Request merged to main branch
- [ ] All steps 1-21 from the workflow have been executed

## Completion Criteria
The full 2018 challenge is considered complete when:
1. All 25 days are solved (both parts each)
2. All tests pass for every day
3. Code is clean, idiomatic Go with proper error handling
4. README is updated with all solutions and results
5. All helper utilities are properly documented
6. Performance is reasonable (each solution runs in < 5 seconds)
7. Daily Completion Checklist verified for each day

## Tool Commands Summary
```bash
# Initialize project
go mod init github.com/username/advent-of-code-2018-go

# Fetch puzzle and input for a day
go run cmd/fetch/main.go -day=1

# Run solution for current directory
go run .

# Run tests
go test
go test -v  # verbose output

# Submit answer
go run cmd/submit/main.go -day=1 -part=1 -answer=42

# Run all tests in project
go test ./...

# Format code
go fmt ./...

# Check for issues
go vet ./...
```

## Sequential Execution Instructions
**CRITICAL**: When user says "start solving", you must:
1. Begin with Day 1 and continue through Day 25 without stopping
2. For each day, complete both parts before moving to the next
3. Submit all answers immediately without asking for confirmation
4. Handle any errors and retry as needed
5. Update progress tracking after each day
6. Only stop when all 25 days are complete or if explicitly told to stop

**RESUMING**: When user says "continue", "resume", or "continue from Day X":
1. Check the solutions/ directory to see which days are already complete
2. Read the README.md to verify progress
3. Identify the next unsolved day/part
4. Continue solving from that point forward
5. Follow the same workflow as initial solving

Remember: The user has NO Go experience, so:
- Explain Go-specific concepts when they appear
- Write clear, readable code over clever optimizations
- Use standard library functions where possible
- Add helpful comments for learning purposes