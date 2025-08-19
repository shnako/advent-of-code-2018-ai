/*
 * Day 4: Repose Record
 * 
 * Part 1: Find the guard that has the most minutes asleep, then find what minute
 * that guard spends asleep the most. Return the product of guard ID and minute.
 * Parse timestamped guard records to track sleep patterns and find the sleepiest guard.
 * 
 * Part 2: Find the guard most frequently asleep on the same minute across all days.
 * Return the product of that guard's ID and the minute they're most often asleep.
 */

package day04

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type EventType int

const (
	BeginShift EventType = iota
	FallAsleep
	WakeUp
)

type Event struct {
	Time     time.Time
	Type     EventType
	GuardID  int
}

type GuardSleep struct {
	TotalMinutes int
	MinuteCounts [60]int // Index is minute, value is count of times asleep at that minute
}

type Solution struct {
	input string
}

func New(input string) *Solution {
	return &Solution{input: strings.TrimSpace(input)}
}

func (s *Solution) Part1() (int, error) {
	events, err := s.parseEvents()
	if err != nil {
		return 0, err
	}

	guards := s.analyzeGuardSleep(events)

	// Find guard with most total sleep minutes
	var sleepiestGuard int
	var maxSleep int
	for guardID, sleep := range guards {
		if sleep.TotalMinutes > maxSleep {
			maxSleep = sleep.TotalMinutes
			sleepiestGuard = guardID
		}
	}

	if sleepiestGuard == 0 {
		return 0, errors.New("no guards found")
	}

	// Find the minute this guard sleeps most often
	sleepData := guards[sleepiestGuard]
	var mostSleptMinute int
	var maxCount int
	for minute, count := range sleepData.MinuteCounts {
		if count > maxCount {
			maxCount = count
			mostSleptMinute = minute
		}
	}

	return sleepiestGuard * mostSleptMinute, nil
}

func (s *Solution) Part2() (int, error) {
	events, err := s.parseEvents()
	if err != nil {
		return 0, err
	}

	guards := s.analyzeGuardSleep(events)

	// Find the guard who is most frequently asleep on the same minute
	var bestGuard int
	var bestMinute int
	var maxFrequency int

	for guardID, sleep := range guards {
		for minute, count := range sleep.MinuteCounts {
			if count > maxFrequency {
				maxFrequency = count
				bestGuard = guardID
				bestMinute = minute
			}
		}
	}

	if bestGuard == 0 {
		return 0, errors.New("no guards found")
	}

	return bestGuard * bestMinute, nil
}

func (s *Solution) parseEvents() ([]Event, error) {
	lines := strings.Split(s.input, "\n")
	events := make([]Event, 0, len(lines))

	// Regex to parse: [1518-11-01 00:00] Guard #10 begins shift
	eventRe := regexp.MustCompile(`^\[(\d{4}-\d{2}-\d{2} \d{2}:\d{2})\] (.+)$`)
	guardRe := regexp.MustCompile(`^Guard #(\d+) begins shift$`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		matches := eventRe.FindStringSubmatch(line)
		if len(matches) != 3 {
			return nil, errors.New("invalid event format: " + line)
		}

		// Parse timestamp
		timestamp, err := time.Parse("2006-01-02 15:04", matches[1])
		if err != nil {
			return nil, err
		}

		// Parse event
		eventText := matches[2]
		var event Event
		event.Time = timestamp

		if guardMatches := guardRe.FindStringSubmatch(eventText); len(guardMatches) == 2 {
			// Guard begins shift
			guardID, err := strconv.Atoi(guardMatches[1])
			if err != nil {
				return nil, err
			}
			event.Type = BeginShift
			event.GuardID = guardID
		} else if eventText == "falls asleep" {
			event.Type = FallAsleep
		} else if eventText == "wakes up" {
			event.Type = WakeUp
		} else {
			return nil, errors.New("unknown event type: " + eventText)
		}

		events = append(events, event)
	}

	// Sort events by timestamp
	sort.Slice(events, func(i, j int) bool {
		return events[i].Time.Before(events[j].Time)
	})

	return events, nil
}

func (s *Solution) analyzeGuardSleep(events []Event) map[int]*GuardSleep {
	guards := make(map[int]*GuardSleep)
	var currentGuard int
	var sleepStart int

	for _, event := range events {
		switch event.Type {
		case BeginShift:
			currentGuard = event.GuardID
			if _, exists := guards[currentGuard]; !exists {
				guards[currentGuard] = &GuardSleep{}
			}

		case FallAsleep:
			sleepStart = event.Time.Minute()

		case WakeUp:
			if currentGuard == 0 {
				continue // Skip if no current guard
			}

			sleepEnd := event.Time.Minute()
			sleepDuration := sleepEnd - sleepStart

			// Update guard's sleep data
			guard := guards[currentGuard]
			guard.TotalMinutes += sleepDuration

			// Mark each minute the guard was asleep
			for minute := sleepStart; minute < sleepEnd; minute++ {
				guard.MinuteCounts[minute]++
			}
		}
	}

	return guards
}