/*
 * Day 7: The Sum of Its Parts
 * 
 * Part 1: Topological sorting of assembly steps with alphabetical tiebreaking.
 * Parse dependency requirements and determine the correct order to complete all steps.
 * When multiple steps are available, choose the one that comes first alphabetically.
 * 
 * Part 2: Simulate parallel execution with 5 workers. Each step takes 60 + letter_value seconds.
 * Multiple workers can work simultaneously on different steps. Track the total time to complete
 * all steps with proper scheduling and timing constraints.
 */

package day07

import (
	"sort"
	"strings"
)

type Solution struct {
	dependencies map[string][]string // step -> list of steps that depend on it
	prerequisites map[string][]string // step -> list of prerequisites
	allSteps      []string
}

func New(input string) *Solution {
	// Handle both Unix (\n) and Windows (\r\n) line endings
	input = strings.ReplaceAll(strings.TrimSpace(input), "\r\n", "\n")
	lines := strings.Split(input, "\n")
	
	dependencies := make(map[string][]string)
	prerequisites := make(map[string][]string)
	stepSet := make(map[string]bool)
	
	for _, line := range lines {
		// Parse "Step X must be finished before step Y can begin."
		parts := strings.Fields(line)
		prerequisite := parts[1]
		dependent := parts[7]
		
		dependencies[prerequisite] = append(dependencies[prerequisite], dependent)
		prerequisites[dependent] = append(prerequisites[dependent], prerequisite)
		
		stepSet[prerequisite] = true
		stepSet[dependent] = true
	}
	
	// Get all steps and sort them alphabetically
	var allSteps []string
	for step := range stepSet {
		allSteps = append(allSteps, step)
	}
	sort.Strings(allSteps)
	
	return &Solution{
		dependencies:  dependencies,
		prerequisites: prerequisites,
		allSteps:      allSteps,
	}
}

func (s *Solution) Part1() (string, error) {
	return s.topologicalSort(), nil
}

func (s *Solution) Part2() (int, error) {
	return s.simulateWithWorkers(5, 60), nil
}

// topologicalSort performs Kahn's algorithm with alphabetical tiebreaking
func (s *Solution) topologicalSort() string {
	// Copy prerequisites map to avoid modifying original
	remainingPrereqs := make(map[string][]string)
	for step, prereqs := range s.prerequisites {
		remainingPrereqs[step] = make([]string, len(prereqs))
		copy(remainingPrereqs[step], prereqs)
	}
	
	var result []string
	var available []string
	
	// Find initial steps with no prerequisites
	for _, step := range s.allSteps {
		if len(remainingPrereqs[step]) == 0 {
			available = append(available, step)
		}
	}
	
	for len(available) > 0 {
		// Sort available steps alphabetically and pick the first one
		sort.Strings(available)
		current := available[0]
		available = available[1:]
		
		result = append(result, current)
		
		// Remove current step from prerequisites of dependent steps
		for _, dependent := range s.dependencies[current] {
			// Remove current from dependent's prerequisites
			newPrereqs := make([]string, 0)
			for _, prereq := range remainingPrereqs[dependent] {
				if prereq != current {
					newPrereqs = append(newPrereqs, prereq)
				}
			}
			remainingPrereqs[dependent] = newPrereqs
			
			// If dependent has no more prerequisites, add it to available
			if len(remainingPrereqs[dependent]) == 0 {
				available = append(available, dependent)
			}
		}
	}
	
	return strings.Join(result, "")
}

// Worker represents a worker with their current task and completion time
type Worker struct {
	currentStep string
	finishTime  int
}

// simulateWithWorkers simulates parallel execution with the given number of workers
func (s *Solution) simulateWithWorkers(numWorkers, baseTime int) int {
	// Copy prerequisites map to avoid modifying original
	remainingPrereqs := make(map[string][]string)
	for step, prereqs := range s.prerequisites {
		remainingPrereqs[step] = make([]string, len(prereqs))
		copy(remainingPrereqs[step], prereqs)
	}
	
	workers := make([]Worker, numWorkers)
	var available []string
	completed := make(map[string]bool)
	currentTime := 0
	
	// Find initial steps with no prerequisites
	for _, step := range s.allSteps {
		if len(remainingPrereqs[step]) == 0 {
			available = append(available, step)
		}
	}
	
	for len(completed) < len(s.allSteps) {
		// Check for completed work and free up workers
		for i := range workers {
			if workers[i].currentStep != "" && currentTime >= workers[i].finishTime {
				completedStep := workers[i].currentStep
				completed[completedStep] = true
				workers[i].currentStep = ""
				workers[i].finishTime = 0
				
				// Remove completed step from prerequisites of dependent steps
				for _, dependent := range s.dependencies[completedStep] {
					if !completed[dependent] {
						newPrereqs := make([]string, 0)
						for _, prereq := range remainingPrereqs[dependent] {
							if prereq != completedStep {
								newPrereqs = append(newPrereqs, prereq)
							}
						}
						remainingPrereqs[dependent] = newPrereqs
						
						// If dependent has no more prerequisites, add it to available
						if len(remainingPrereqs[dependent]) == 0 {
							available = append(available, dependent)
						}
					}
				}
			}
		}
		
		// Sort available steps alphabetically
		sort.Strings(available)
		
		// Assign work to available workers
		for i := range workers {
			if workers[i].currentStep == "" && len(available) > 0 {
				step := available[0]
				available = available[1:]
				
				workers[i].currentStep = step
				workers[i].finishTime = currentTime + baseTime + int(step[0]-'A'+1)
			}
		}
		
		// Find the next time when a worker will finish
		nextFinishTime := -1
		for i := range workers {
			if workers[i].currentStep != "" {
				if nextFinishTime == -1 || workers[i].finishTime < nextFinishTime {
					nextFinishTime = workers[i].finishTime
				}
			}
		}
		
		if nextFinishTime != -1 {
			currentTime = nextFinishTime
		} else {
			// No one is working, but if there are still available steps, assign them
			if len(available) > 0 {
				currentTime++
			} else {
				break
			}
		}
	}
	
	return currentTime
}