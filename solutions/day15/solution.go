/*
 * Day 15: Beverage Bandits
 * 
 * Part 1: Simulate combat between Elves and Goblins on a grid.
 * Units take turns in reading order, move toward enemies using shortest path,
 * and attack adjacent enemies with lowest HP. Combat ends when one side has no targets.
 * Calculate outcome as (full rounds completed) × (sum of remaining HP).
 * 
 * Part 2: Find the minimum attack power for Elves such that no Elf dies during combat.
 * Binary search on Elf attack power until all Elves survive.
 */

package day15

import (
	"fmt"
	"sort"
	"strings"
)

type Point struct {
	X, Y int
}

type Unit struct {
	Pos        Point
	Type       rune  // 'E' or 'G'
	HP         int
	AttackPower int
	Alive      bool
}

type Game struct {
	Grid  [][]rune
	Units []*Unit
	Width int
	Height int
}

type Solution struct {
	input string
}

func New(input string) *Solution {
	return &Solution{input: strings.TrimSpace(input)}
}

func (s *Solution) Part1() (int, error) {
	game := s.parseInput()
	
	rounds := 0
	for {
		// Check if combat should end
		elves, goblins := game.countLiving()
		if elves == 0 || goblins == 0 {
			break
		}
		
		// Get units in reading order
		units := game.getLivingUnitsInOrder()
		
		roundCompleted := true
		for _, unit := range units {
			if !unit.Alive {
				continue
			}
			
			// Check if combat should end during this round
			elves, goblins := game.countLiving()
			if elves == 0 || goblins == 0 {
				roundCompleted = false
				break
			}
			
			// Take unit's turn
			game.takeTurn(unit)
		}
		
		if roundCompleted {
			rounds++
		}
	}
	
	// Calculate outcome
	totalHP := 0
	for _, unit := range game.Units {
		if unit.Alive {
			totalHP += unit.HP
		}
	}
	
	return rounds * totalHP, nil
}

func (s *Solution) Part2() (int, error) {
	// Find upper bound using exponential search
	low := 4
	high := 4
	for {
		game := s.parseInput()
		initialElves := game.countInitialElves()
		
		// Set elf attack power to test
		for _, unit := range game.Units {
			if unit.Type == 'E' {
				unit.AttackPower = high
			}
		}
		
		// Run combat simulation
		rounds := 0
		for {
			elves, goblins := game.countLiving()
			if elves == 0 || goblins == 0 {
				break
			}
			
			units := game.getLivingUnitsInOrder()
			
			roundCompleted := true
			for _, unit := range units {
				if !unit.Alive {
					continue
				}
				
				elves, goblins := game.countLiving()
				if elves == 0 || goblins == 0 {
					roundCompleted = false
					break
				}
				
				game.takeTurn(unit)
			}
			
			if roundCompleted {
				rounds++
			}
		}
		
		// Check if all elves survived
		finalElves, _ := game.countLiving()
		if finalElves == initialElves {
			// All elves survived, we found a valid upper bound
			break
		}
		
		high *= 2
		if high > 1000 { // Safety cap
			return 0, fmt.Errorf("attack power exceeded reasonable bounds")
		}
	}
	
	// Binary search for minimum elf attack power where no elf dies
	for low < high {
		mid := (low + high) / 2
		
		game := s.parseInput()
		initialElves := game.countInitialElves()
		
		// Set elf attack power
		for _, unit := range game.Units {
			if unit.Type == 'E' {
				unit.AttackPower = mid
			}
		}
		
		rounds := 0
		for {
			elves, goblins := game.countLiving()
			if elves == 0 || goblins == 0 {
				break
			}
			
			units := game.getLivingUnitsInOrder()
			
			roundCompleted := true
			for _, unit := range units {
				if !unit.Alive {
					continue
				}
				
				elves, goblins := game.countLiving()
				if elves == 0 || goblins == 0 {
					roundCompleted = false
					break
				}
				
				game.takeTurn(unit)
			}
			
			if roundCompleted {
				rounds++
			}
		}
		
		// Check if all elves survived
		finalElves, _ := game.countLiving()
		if finalElves == initialElves {
			// All elves survived, try lower attack power
			high = mid
		} else {
			// Some elves died, need higher attack power
			low = mid + 1
		}
	}
	
	// Calculate outcome with the minimum attack power
	game := s.parseInput()
	for _, unit := range game.Units {
		if unit.Type == 'E' {
			unit.AttackPower = low
		}
	}
	
	rounds := 0
	for {
		elves, goblins := game.countLiving()
		if elves == 0 || goblins == 0 {
			break
		}
		
		units := game.getLivingUnitsInOrder()
		
		roundCompleted := true
		for _, unit := range units {
			if !unit.Alive {
				continue
			}
			
			elves, goblins := game.countLiving()
			if elves == 0 || goblins == 0 {
				roundCompleted = false
				break
			}
			
			game.takeTurn(unit)
		}
		
		if roundCompleted {
			rounds++
		}
	}
	
	totalHP := 0
	for _, unit := range game.Units {
		if unit.Alive {
			totalHP += unit.HP
		}
	}
	
	return rounds * totalHP, nil
}

func (s *Solution) parseInput() *Game {
	lines := strings.Split(s.input, "\n")
	height := len(lines)
	width := len(lines[0])
	
	grid := make([][]rune, height)
	var units []*Unit
	
	for y, raw := range lines {
		line := strings.TrimRight(raw, "\r")
		grid[y] = []rune(line)
		for x, char := range line {
			if char == 'E' || char == 'G' {
				unit := &Unit{
					Pos:         Point{X: x, Y: y},
					Type:        char,
					HP:          200,
					AttackPower: 3,
					Alive:       true,
				}
				units = append(units, unit)
				grid[y][x] = '.' // Replace with open space
			}
		}
	}
	
	return &Game{
		Grid:   grid,
		Units:  units,
		Width:  width,
		Height: height,
	}
}

func (g *Game) countLiving() (elves, goblins int) {
	for _, unit := range g.Units {
		if unit.Alive {
			if unit.Type == 'E' {
				elves++
			} else {
				goblins++
			}
		}
	}
	return
}

func (g *Game) countInitialElves() int {
	count := 0
	for _, unit := range g.Units {
		if unit.Type == 'E' {
			count++
		}
	}
	return count
}

func (g *Game) getLivingUnitsInOrder() []*Unit {
	var living []*Unit
	for _, unit := range g.Units {
		if unit.Alive {
			living = append(living, unit)
		}
	}
	
	// Sort in reading order
	sort.Slice(living, func(i, j int) bool {
		if living[i].Pos.Y == living[j].Pos.Y {
			return living[i].Pos.X < living[j].Pos.X
		}
		return living[i].Pos.Y < living[j].Pos.Y
	})
	
	return living
}

func (g *Game) takeTurn(unit *Unit) {
	// Find targets
	targets := g.findTargets(unit)
	if len(targets) == 0 {
		return // No targets, combat ends
	}
	
	// Check if already in range to attack
	adjacentEnemies := g.findAdjacentEnemies(unit)
	if len(adjacentEnemies) == 0 {
		// Not in range, try to move
		g.moveUnit(unit, targets)
		// After moving, check for adjacent enemies again
		adjacentEnemies = g.findAdjacentEnemies(unit)
	}
	
	// Attack if possible
	if len(adjacentEnemies) > 0 {
		g.attack(unit, adjacentEnemies)
	}
}

func (g *Game) findTargets(unit *Unit) []*Unit {
	var targets []*Unit
	enemyType := 'G'
	if unit.Type == 'G' {
		enemyType = 'E'
	}
	
	for _, other := range g.Units {
		if other.Alive && other.Type == enemyType {
			targets = append(targets, other)
		}
	}
	
	return targets
}

func (g *Game) findAdjacentEnemies(unit *Unit) []*Unit {
	directions := []Point{{0, -1}, {-1, 0}, {1, 0}, {0, 1}} // Up, Left, Right, Down
	var enemies []*Unit
	
	for _, dir := range directions {
		pos := Point{unit.Pos.X + dir.X, unit.Pos.Y + dir.Y}
		if enemy := g.getUnitAt(pos); enemy != nil && enemy.Alive && enemy.Type != unit.Type {
			enemies = append(enemies, enemy)
		}
	}
	
	return enemies
}

func (g *Game) moveUnit(unit *Unit, targets []*Unit) {
	// Find all squares in range of targets
	inRange := g.findInRangeSquares(targets)
	if len(inRange) == 0 {
		return // No reachable squares
	}
	
	// Find reachable squares and their distances
	reachable := g.findReachableSquares(unit.Pos, inRange)
	if len(reachable) == 0 {
		return // No reachable squares
	}
	
	// Find minimum distance
	minDist := reachable[0].dist
	for _, r := range reachable {
		if r.dist < minDist {
			minDist = r.dist
		}
	}
	
	// Filter to squares with minimum distance
	var nearest []reachableSquare
	for _, r := range reachable {
		if r.dist == minDist {
			nearest = append(nearest, r)
		}
	}
	
	// Choose first in reading order
	sort.Slice(nearest, func(i, j int) bool {
		if nearest[i].pos.Y == nearest[j].pos.Y {
			return nearest[i].pos.X < nearest[j].pos.X
		}
		return nearest[i].pos.Y < nearest[j].pos.Y
	})
	
	target := nearest[0].pos
	
	// Find first step toward target
	if nextStep, ok := g.findFirstStep(unit.Pos, target); ok {
		unit.Pos = nextStep
	}
}

type reachableSquare struct {
	pos  Point
	dist int
}

func (g *Game) findInRangeSquares(targets []*Unit) []Point {
	var inRange []Point
	directions := []Point{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}
	
	seen := make(map[Point]bool)
	
	for _, target := range targets {
		for _, dir := range directions {
			pos := Point{target.Pos.X + dir.X, target.Pos.Y + dir.Y}
			if !seen[pos] && g.isOpen(pos) {
				inRange = append(inRange, pos)
				seen[pos] = true
			}
		}
	}
	
	return inRange
}

func (g *Game) findReachableSquares(start Point, targets []Point) []reachableSquare {
	var reachable []reachableSquare
	
	for _, target := range targets {
		if dist := g.shortestPath(start, target); dist >= 0 {
			reachable = append(reachable, reachableSquare{target, dist})
		}
	}
	
	return reachable
}

func (g *Game) shortestPath(start, end Point) int {
	if start == end {
		return 0
	}
	
	queue := []Point{start}
	distances := map[Point]int{start: 0}
	directions := []Point{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}
	
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		
		for _, dir := range directions {
			next := Point{current.X + dir.X, current.Y + dir.Y}
			
			if _, visited := distances[next]; visited {
				continue
			}
			
			if !g.isOpen(next) {
				continue
			}
			
			distances[next] = distances[current] + 1
			
			if next == end {
				return distances[next]
			}
			
			queue = append(queue, next)
		}
	}
	
	return -1 // Not reachable
}

func (g *Game) findFirstStep(start, end Point) (Point, bool) {
	if start == end {
		return Point{}, false
	}
	
	queue := []Point{start}
	distances := map[Point]int{start: 0}
	parent := map[Point]Point{start: Point{-1, -1}}
	directions := []Point{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}
	
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		
		for _, dir := range directions {
			next := Point{current.X + dir.X, current.Y + dir.Y}
			
			if _, visited := distances[next]; visited {
				continue
			}
			
			if !g.isOpen(next) {
				continue
			}
			
			distances[next] = distances[current] + 1
			parent[next] = current
			queue = append(queue, next)
			
			if next == end {
				// Trace back to find first step
				path := []Point{end}
				curr := end
				for parent[curr] != start {
					curr = parent[curr]
					path = append(path, curr)
				}
				return curr, true
			}
		}
	}
	
	return Point{}, false // Not reachable
}

func (g *Game) attack(unit *Unit, enemies []*Unit) {
	// Find enemy with lowest HP (reading order tiebreaker)
	sort.Slice(enemies, func(i, j int) bool {
		if enemies[i].HP == enemies[j].HP {
			if enemies[i].Pos.Y == enemies[j].Pos.Y {
				return enemies[i].Pos.X < enemies[j].Pos.X
			}
			return enemies[i].Pos.Y < enemies[j].Pos.Y
		}
		return enemies[i].HP < enemies[j].HP
	})
	
	target := enemies[0]
	target.HP -= unit.AttackPower
	
	if target.HP <= 0 {
		target.Alive = false
	}
}

func (g *Game) isOpen(pos Point) bool {
	if pos.X < 0 || pos.X >= g.Width || pos.Y < 0 || pos.Y >= g.Height {
		return false
	}
	
	if g.Grid[pos.Y][pos.X] != '.' {
		return false
	}
	
	// Check if occupied by a unit
	return g.getUnitAt(pos) == nil
}

func (g *Game) getUnitAt(pos Point) *Unit {
	for _, unit := range g.Units {
		if unit.Alive && unit.Pos == pos {
			return unit
		}
	}
	return nil
}