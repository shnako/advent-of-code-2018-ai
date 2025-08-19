# Advent of Code Automation Setup

This guide explains how to automate fetching puzzles and submitting solutions for Advent of Code.

## Prerequisites

1. Java/Kotlin (already installed with your project)
2. Active Advent of Code account

## Getting Your Session Cookie

1. Log into [adventofcode.com](https://adventofcode.com)
2. Open browser Developer Tools (F12)
3. Go to Application/Storage → Cookies → adventofcode.com
4. Find the cookie named `session` and copy its value
5. Set it as an environment variable:

### Windows (PowerShell)
```powershell
$env:AOC_SESSION_COOKIE = "your_session_cookie_here"
```

### Windows (Command Prompt)
```cmd
set AOC_SESSION_COOKIE=your_session_cookie_here
```

### Linux/Mac
```bash
export AOC_SESSION_COOKIE="your_session_cookie_here"
```

### Permanent Setup (recommended)
Add to your `.bashrc`, `.zshrc`, or Windows environment variables.

## Usage

### Method 1: Using Gradle Tasks (Recommended)

#### Fetch a puzzle and input
```bash
./gradlew fetchDay -Pday=20
```

#### Submit a solution
```bash
# Submit Part 1
./gradlew submitAnswer -Pday=20 -Ppart=1 -Panswer=42

# Submit Part 2  
./gradlew submitAnswer -Pday=20 -Ppart=2 -Panswer=12345
```

### Method 2: Direct Kotlin Execution

#### Fetch a puzzle
```bash
./gradlew aoc -Pargs="20 fetch"
# or just
./gradlew aoc -Pargs="20"
```

#### Submit a solution
```bash
./gradlew aoc -Pargs="20 submit 1 42"
```

### What Gets Created

When fetching a day, these files are created:
- `src/main/kotlin/solutions/day20/Solution.kt` - Skeleton solution class
- `src/main/resources/solutions/day20/puzzle.txt` - Puzzle description
- `src/main/resources/solutions/day20/input.txt` - Your personal input
- `src/test/kotlin/solutions/day20/SolutionTest.kt` - Test skeleton

## Integration with Claude

Once set up, you can tell Claude:
1. "Fetch and solve day 20" - Claude will run the fetch command and solve the puzzle
2. "Submit the answer X for part 1" - Claude will submit the solution

## Workflow with Claude

1. **Setup session cookie** (one-time):
   ```powershell
   # Windows PowerShell
   $env:AOC_SESSION_COOKIE = "your_cookie"
   ```
   ```bash
   # Linux/Mac
   export AOC_SESSION_COOKIE="your_cookie"
   ```

2. **Tell Claude the day number**:
   "Let's do day 20 of Advent of Code"

3. **Claude will automatically**:
   - Run `./gradlew fetchDay -Pday=20` to fetch puzzle and input
   - Read and understand the puzzle
   - Implement Part 1 solution
   - Test with the example
   - Run with your input
   - **Submit Part 1 immediately** (no confirmation needed)
   - If correct, fetch Part 2 (automatic)
   - Implement Part 2 solution
   - Test Part 2
   - **Submit Part 2 immediately**
   - Update test file with both answers
   - Update README
   - Create branch and pull request

### Fully Automated Process
Claude will complete the entire day without stopping:
- ✅ Auto-submits answers as soon as they're calculated
- ✅ Continues to Part 2 after Part 1 succeeds
- ✅ Handles wrong answers by debugging and retrying
- ✅ Creates PR when both parts are complete

## Security Notes

- **Never commit your session cookie** to git
- Consider using a `.env` file (add to `.gitignore`)
- Session cookies expire after ~30 days
- Be respectful of AoC servers - don't make excessive requests

## Troubleshooting

### "Error: AOC_SESSION_COOKIE environment variable not set"
- Make sure you've set the environment variable in your current shell session

### "HTTP 400" or "HTTP 500" errors
- Your session cookie may be expired. Get a new one from the browser.

### "You gave an answer too recently"
- AoC has rate limiting. Wait 1-5 minutes between incorrect submissions.

### Build errors
- Make sure you're in the project root directory
- Run `./gradlew build` to ensure everything compiles

## Quick Start Example

```powershell
# 1. Set your session cookie (get from browser)
$env:AOC_SESSION_COOKIE = "53616c7465645f5f..."

# 2. Fetch day 20
./gradlew fetchDay -Pday=20

# 3. Solve the puzzle (manually or with Claude)

# 4. Submit part 1 answer
./gradlew submitAnswer -Pday=20 -Ppart=1 -Panswer=12345

# 5. Submit part 2 answer
./gradlew submitAnswer -Pday=20 -Ppart=2 -Panswer=67890
```