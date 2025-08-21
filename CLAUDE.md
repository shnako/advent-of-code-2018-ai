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
    - **CRITICAL**: Do NOT implement Part 2 before submitting Part 1 successfully 
    - **WAIT** for Part 1 submission confirmation before reading Part 2 problem
    - Only implement what is explicitly stated in the current puzzle.txt
10. **Implement Part 2**: Update solution.go with Part 2 solution
    - **ONLY AFTER** Part 1 is confirmed correct and Part 2 is available
11. **Test Part 2**: Run with any Part 2 examples
12. **Run Part 2**: Execute with the full input
13. **Submit Part 2**: Immediately submit the Part 2 answer
14. **Run final tests**: Ensure all tests pass with `go test`
15. **Update progress**: Mark day as complete in progress tracking
16. **Create Pull Request**: Create PR with comprehensive description of the solution
17. **CodeRabbit Review Cycle (Automated)**: 
    - Wait 3 minutes initially for CodeRabbit to start its review
    - Check every minute for new comments using `gh pr view PR_NUMBER --comments`
    - **DISTINGUISH between summary and actual review**:
      - **SUMMARY ONLY**: If CodeRabbit only provides a walkthrough/summary without specific code feedback, WAIT for the actual review
      - **ACTUAL REVIEW**: Look for specific code review comments with suggestions, improvements, or issues
    - **RATE LIMIT HANDLING**: If CodeRabbit indicates rate limit exceeded:
      - **NEVER MERGE during rate limit** — this bypasses the review process entirely
      - Wait for the FULL specified time plus a small buffer (e.g., "6 minutes and 56 seconds" → wait 7 minutes + 10 seconds) to avoid early re-triggers
      - Post `@coderabbitai review` ONLY ONCE per rate-limit event after the wait time (do not spam repeated triggers)
      - Add a PR comment noting the rate-limit message and the next ETA before sleeping (for auditability)
      - Continue checking every 60 seconds until CodeRabbit provides actual, line-level review feedback
      - If the same rate-limit recurs 3 times consecutively or total wait exceeds 30 minutes, follow "Handling Rate Limits" below (treat as a daily cap) and pause/resume instead of looping indefinitely
      - **CRITICAL**: Rate limit does NOT excuse skipping the review — follow the bounded retry policy above, then pause according to "Handling Rate Limits"
    - **ONLY PROCEED when CodeRabbit provides ACTUAL CODE REVIEW**:
      - Look for specific line-by-line comments with actionable feedback
      - Ignore initial summary/walkthrough - this is NOT the review
      - CodeRabbit will provide specific suggestions, optimizations, or issues to address
      - **NEVER MERGE while CodeRabbit shows "Currently processing new changes" or similar processing messages**
      - **WAIT until CodeRabbit fully completes its review** - processing messages indicate incomplete review
    - **CRITICAL: MANDATORY ADDRESS ALL FEEDBACK**: Once CodeRabbit's ACTUAL review comments are detected, you MUST address ALL comments that make sense by making necessary code changes
      - **ADDRESS ALL COMMENT TYPES**: This includes ALL comments marked as:
        - **Regular comments**: Standard code review feedback
        - **Nitpick comments**: Small improvements that should still be addressed
        - **Duplicate comments**: Similar issues in multiple places - fix all occurrences
        - **Additional comments**: Extra suggestions that improve code quality
      - **FIRST STEP - REPLY TO COMMENTS**: Before making any code changes, REPLY to each CodeRabbit review comment acknowledging it and explaining how you will address it
      - **SECOND STEP - MAKE CODE CHANGES**: After replying to all comments, make the necessary code changes to fix all issues, suggestions, and improvements
      - **NEVER SKIP FEEDBACK**: Every single piece of actionable feedback must be addressed before merging, including nitpicks and duplicates
      - **NO EXCEPTIONS**: CodeRabbit feedback is not optional - all comments that make sense must be implemented
      - **DO NOT PROCEED TO BUILD/MERGE without addressing feedback first**
    - Push changes to the feature branch after addressing all feedback
    - **CRITICAL**: After EVERY push, CodeRabbit will automatically re-review and potentially add NEW comments
    - **ABSOLUTE REQUIREMENT**: After EVERY SINGLE COMMIT you make to the PR, you MUST wait for CodeRabbit to provide NEW review comments and address ALL of them before proceeding
    - **NEVER SKIP WAITING**: Even if you think your fix is minor, CodeRabbit WILL re-review and may have additional feedback
    - **MANDATORY WAIT CYCLE**: After each commit: 1) Wait for CodeRabbit processing to complete, 2) Wait for actual review comments, 3) Address ALL feedback, 4) Only then proceed to next step
    - **ONLY REQUEST REVIEW AFTER RATE LIMITS**: Post `@coderabbitai review` ONLY when rate limits have been encountered - normal pushes trigger automatic review
    - **AFTER RATE LIMIT RECOVERY**: When rate limit expires, MUST post `@coderabbitai review` to actually get the review - waiting without requesting is pointless
    - Continue checking every minute for NEW CodeRabbit comments after each push
    - **USE POWERSHELL FOR WAITING**: Use `powershell -command "Start-Sleep -Seconds X"` for non-interactive waits
    - **ALWAYS SET MAXIMUM TIMEOUT**: When using sleep commands, ALWAYS set timeout=600000 (10 minutes max) regardless of actual sleep duration to prevent premature timeouts
    - **DO NOT assume review is complete** after addressing initial comments - wait for re-review
    - **CRITICAL VIOLATION**: Attempting to merge while CodeRabbit has pending feedback or is still processing is a MAJOR WORKFLOW VIOLATION
    - Repeat this automated cycle until CodeRabbit stops adding new substantive review comments
    - Only proceed to merge when CodeRabbit has completed its actual review AND has no more feedback
    - **DO NOT MERGE after just a summary** - wait for the complete review process
    - **DO NOT MERGE while CodeRabbit is still processing** - wait for processing to complete fully
    - **DO NOT STOP**: This process should be fully automated without user intervention
18. **Verify ALL CI checks pass**: **CRITICAL - ALL THREE CHECKS MUST PASS**: After addressing all CodeRabbit feedback, ensure ALL CI checks pass
    - **MANDATORY**: Run `go build ./...` to verify the entire project builds successfully
    - **MANDATORY**: Run `go test ./...` to verify all tests pass across the project
    - **MANDATORY**: Ensure validation check passes - fix ANY puzzle.txt formatting issues across ALL days
    - **THREE REQUIRED CHECKS**: 1) Tests must pass, 2) Validation must pass, 3) CodeRabbit must complete
    - **STOP IF ANY FAILURES**: Fix any build errors, test failures, OR validation failures before proceeding to merge
    - **NO MERGE until 100% success**: All three CI checks (test, validate, CodeRabbit) must be green/passing
19. **Merge PR**: **ONLY** after ALL THREE CI checks pass (test=SUCCESS, validate=SUCCESS, CodeRabbit=SUCCESS), merge the PR using "Squash and merge"
20. **VERIFY COMPLIANCE**: Double-check that ALL steps in this CLAUDE.md workflow have been followed:
    - Did I fetch both puzzle and input?
    - Did I implement and test both parts?
    - Did I submit both answers immediately?
    - Did I update the README progress table?
    - Did I run all tests successfully?
    - Is the code clean, idiomatic Go with proper error handling?
    - Have I followed all Go best practices outlined in this document?
    - Did I complete the full CodeRabbit review cycle (actual review, not just summary)?
    - Did I merge the PR successfully?
21. **Continue to next day**: Automatically proceed to the next day without stopping

### Important for Automation:
- **Never stop during the entire process**: Continue solving until all 25 days are complete
- **The fetch command downloads BOTH puzzle and input** using the session cookie
- **input.txt contains your personal puzzle input** (different for each user)
- **Always submit immediately** when you have an answer (don't wait for user confirmation)
- **Continue to Part 2 automatically** after Part 1 is confirmed correct
- **Handle wrong answers**: If an answer is wrong, debug and try again
- **Automated PR Review**: Wait and check for CodeRabbit comments automatically, don't stop for user input
- **Complete all 25 days**: Don't stop until all puzzles are solved and all PRs are merged
- **Quality Assurance**: All steps 1-20 MUST be completed - verify full compliance with this document before moving on

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
- [ ] All steps 1-20 from the workflow have been executed

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