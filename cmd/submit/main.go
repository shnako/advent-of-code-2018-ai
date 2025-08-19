package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	day := flag.Int("day", 0, "Day to submit (1-25)")
	part := flag.Int("part", 0, "Part to submit (1 or 2)")
	answer := flag.String("answer", "", "Answer to submit")
	flag.Parse()

	if *day < 1 || *day > 25 {
		fmt.Fprintf(os.Stderr, "Day must be between 1 and 25\n")
		os.Exit(1)
	}

	if *part != 1 && *part != 2 {
		fmt.Fprintf(os.Stderr, "Part must be 1 or 2\n")
		os.Exit(1)
	}

	if *answer == "" {
		fmt.Fprintf(os.Stderr, "Answer must be provided\n")
		os.Exit(1)
	}

	sessionCookie := os.Getenv("AOC_SESSION_COOKIE")
	if sessionCookie == "" {
		fmt.Fprintf(os.Stderr, "AOC_SESSION_COOKIE environment variable not set\n")
		os.Exit(1)
	}

	submitURL := fmt.Sprintf("https://adventofcode.com/2018/day/%d/answer", *day)

	// Prepare form data
	formData := url.Values{}
	formData.Set("level", fmt.Sprintf("%d", *part))
	formData.Set("answer", *answer)

	// Create request
	client := &http.Client{}
	req, err := http.NewRequest("POST", submitURL, strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create request: %v\n", err)
		os.Exit(1)
	}

	req.Header.Set("Cookie", fmt.Sprintf("session=%s", sessionCookie))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "github.com/shnako/advent-of-code-2018-ai")

	// Submit answer
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to submit answer: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "HTTP %d: %s\n", resp.StatusCode, resp.Status)
		os.Exit(1)
	}

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read response: %v\n", err)
		os.Exit(1)
	}

	responseText := string(body)

	// Check for common responses
	if strings.Contains(responseText, "That's the right answer") {
		fmt.Printf("✓ Day %d Part %d: Correct! Answer: %s\n", *day, *part, *answer)
		
		// Check if we got a new puzzle part
		if strings.Contains(responseText, "You've unlocked") || strings.Contains(responseText, "second half") {
			fmt.Println("Part 2 unlocked! Run fetch command again to get the updated puzzle.")
		}
	} else if strings.Contains(responseText, "That's not the right answer") {
		fmt.Printf("✗ Day %d Part %d: Incorrect answer: %s\n", *day, *part, *answer)
		
		// Extract wait time if present
		if strings.Contains(responseText, "Please wait") {
			start := strings.Index(responseText, "Please wait")
			end := strings.Index(responseText[start:], "before trying again")
			if end != -1 {
				fmt.Println(responseText[start : start+end+20])
			}
		}
		
		// Check if too high or too low
		if strings.Contains(responseText, "too high") {
			fmt.Println("Hint: Your answer is too high")
		} else if strings.Contains(responseText, "too low") {
			fmt.Println("Hint: Your answer is too low")
		}
	} else if strings.Contains(responseText, "You gave an answer too recently") {
		fmt.Printf("⏳ Day %d Part %d: Rate limited. Please wait before submitting again.\n", *day, *part)
	} else if strings.Contains(responseText, "You don't seem to be solving the right level") {
		fmt.Printf("⚠ Day %d Part %d: Already solved or wrong part\n", *day, *part)
	} else {
		fmt.Printf("? Day %d Part %d: Unknown response for answer: %s\n", *day, *part, *answer)
		fmt.Println("Response snippet:", getResponseSnippet(responseText))
	}
}

func getResponseSnippet(html string) string {
	// Extract the main article content for debugging
	start := strings.Index(html, "<article")
	if start == -1 {
		return "Could not parse response"
	}
	
	end := strings.Index(html[start:], "</article>")
	if end == -1 {
		return "Could not parse response"
	}
	
	article := html[start : start+end+10]
	
	// Remove HTML tags for readability
	text := article
	for {
		tagStart := strings.Index(text, "<")
		if tagStart == -1 {
			break
		}
		tagEnd := strings.Index(text[tagStart:], ">")
		if tagEnd == -1 {
			break
		}
		text = text[:tagStart] + " " + text[tagStart+tagEnd+1:]
	}
	
	// Clean up whitespace
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", " ")
	for strings.Contains(text, "  ") {
		text = strings.ReplaceAll(text, "  ", " ")
	}
	
	text = strings.TrimSpace(text)
	if len(text) > 200 {
		text = text[:200] + "..."
	}
	
	return text
}