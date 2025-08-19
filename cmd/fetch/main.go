package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	day := flag.Int("day", 0, "Day to fetch (1-25)")
	flag.Parse()

	if *day < 1 || *day > 25 {
		fmt.Fprintf(os.Stderr, "Day must be between 1 and 25\n")
		os.Exit(1)
	}

	sessionCookie := os.Getenv("AOC_SESSION_COOKIE")
	if sessionCookie == "" {
		fmt.Fprintf(os.Stderr, "AOC_SESSION_COOKIE environment variable not set\n")
		os.Exit(1)
	}

	dayStr := fmt.Sprintf("%02d", *day)
	solutionDir := filepath.Join("solutions", fmt.Sprintf("day%s", dayStr))

	// Create solution directory
	if err := os.MkdirAll(solutionDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create directory: %v\n", err)
		os.Exit(1)
	}

	// Fetch puzzle description
	puzzleURL := fmt.Sprintf("https://adventofcode.com/2018/day/%d", *day)
	puzzleContent, err := fetchContent(puzzleURL, sessionCookie, true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to fetch puzzle: %v\n", err)
		os.Exit(1)
	}

	puzzlePath := filepath.Join(solutionDir, "puzzle.txt")
	if err := os.WriteFile(puzzlePath, []byte(puzzleContent), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write puzzle: %v\n", err)
		os.Exit(1)
	}

	// Fetch input
	inputURL := fmt.Sprintf("https://adventofcode.com/2018/day/%d/input", *day)
	inputContent, err := fetchContent(inputURL, sessionCookie, false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to fetch input: %v\n", err)
		os.Exit(1)
	}

	inputPath := filepath.Join(solutionDir, "input.txt")
	if err := os.WriteFile(inputPath, []byte(inputContent), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully fetched Day %d puzzle and input to %s\n", *day, solutionDir)
}

func fetchContent(url, sessionCookie string, isHTML bool) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Cookie", fmt.Sprintf("session=%s", sessionCookie))
	req.Header.Set("User-Agent", "github.com/shnak/advent-of-code-2018-go")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	content := string(body)

	if isHTML {
		// Extract the main article content
		start := strings.Index(content, "<article")
		end := strings.LastIndex(content, "</article>")
		if start != -1 && end != -1 {
			end += len("</article>")
			content = content[start:end]
		}

		// Convert HTML to plain text (basic conversion)
		content = stripHTML(content)
	}

	return strings.TrimSpace(content), nil
}

func stripHTML(html string) string {
	// Basic HTML to text conversion
	text := html

	// Remove style and script tags with their content
	for strings.Contains(text, "<style") {
		start := strings.Index(text, "<style")
		end := strings.Index(text[start:], "</style>")
		if end != -1 {
			text = text[:start] + text[start+end+8:]
		} else {
			break
		}
	}

	// Replace common tags
	replacements := map[string]string{
		"<br>":  "\n",
		"<br/>": "\n",
		"<br />": "\n",
		"</p>":  "\n\n",
		"</div>": "\n",
		"<li>":  "- ",
		"</li>": "\n",
		"&lt;":  "<",
		"&gt;":  ">",
		"&amp;": "&",
		"&nbsp;": " ",
		"<em>": "*",
		"</em>": "*",
		"<code>": "`",
		"</code>": "`",
		"<pre>": "\n```\n",
		"</pre>": "\n```\n",
	}

	for old, new := range replacements {
		text = strings.ReplaceAll(text, old, new)
	}

	// Remove remaining HTML tags
	for {
		start := strings.Index(text, "<")
		if start == -1 {
			break
		}
		end := strings.Index(text[start:], ">")
		if end == -1 {
			break
		}
		text = text[:start] + text[start+end+1:]
	}

	// Clean up multiple newlines
	for strings.Contains(text, "\n\n\n") {
		text = strings.ReplaceAll(text, "\n\n\n", "\n\n")
	}

	return strings.TrimSpace(text)
}