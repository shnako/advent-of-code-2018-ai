package day24

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Group struct {
	army       string
	units      int
	hp         int
	immunities []string
	weaknesses []string
	damage     int
	damageType string
	initiative int
}

func (g *Group) effectivePower() int {
	return g.units * g.damage
}

func (g *Group) damageDealtTo(defender *Group) int {
	if g.units <= 0 {
		return 0
	}
	
	damage := g.effectivePower()
	
	// Check immunities
	for _, immunity := range defender.immunities {
		if immunity == g.damageType {
			return 0
		}
	}
	
	// Check weaknesses
	for _, weakness := range defender.weaknesses {
		if weakness == g.damageType {
			return damage * 2
		}
	}
	
	return damage
}

func parseInput(input string) []*Group {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	groups := []*Group{}
	currentArmy := ""
	
	groupRegex := regexp.MustCompile(`(\d+) units each with (\d+) hit points (?:\(([^)]+)\) )?with an attack that does (\d+) (\w+) damage at initiative (\d+)`)
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "Immune System:") {
			currentArmy = "Immune System"
			continue
		}
		if strings.HasPrefix(line, "Infection:") {
			currentArmy = "Infection"
			continue
		}
		
		matches := groupRegex.FindStringSubmatch(line)
		if matches == nil {
			continue
		}
		
		units, _ := strconv.Atoi(matches[1])
		hp, _ := strconv.Atoi(matches[2])
		damage, _ := strconv.Atoi(matches[4])
		damageType := matches[5]
		initiative, _ := strconv.Atoi(matches[6])
		
		group := &Group{
			army:       currentArmy,
			units:      units,
			hp:         hp,
			damage:     damage,
			damageType: damageType,
			initiative: initiative,
			immunities: []string{},
			weaknesses: []string{},
		}
		
		// Parse immunities and weaknesses
		if matches[3] != "" {
			parts := strings.Split(matches[3], "; ")
			for _, part := range parts {
				if strings.HasPrefix(part, "immune to ") {
					immuneStr := strings.TrimPrefix(part, "immune to ")
					group.immunities = strings.Split(strings.ReplaceAll(immuneStr, " ", ""), ",")
				} else if strings.HasPrefix(part, "weak to ") {
					weakStr := strings.TrimPrefix(part, "weak to ")
					group.weaknesses = strings.Split(strings.ReplaceAll(weakStr, " ", ""), ",")
				}
			}
		}
		
		groups = append(groups, group)
	}
	
	return groups
}

func fight(groups []*Group) []*Group {
	// Make copies to avoid modifying originals
	combatGroups := make([]*Group, len(groups))
	for i, g := range groups {
		newGroup := *g
		combatGroups[i] = &newGroup
	}
	
	for {
		// Check if combat is over
		immuneCount := 0
		infectionCount := 0
		for _, g := range combatGroups {
			if g.units > 0 {
				if g.army == "Immune System" {
					immuneCount++
				} else {
					infectionCount++
				}
			}
		}
		if immuneCount == 0 || infectionCount == 0 {
			break
		}
		
		// Target selection phase
		targets := make(map[int]int) // attacker index -> defender index
		targeted := make(map[int]bool)
		
		// Sort for target selection (by effective power desc, then initiative desc)
		order := make([]int, len(combatGroups))
		for i := range order {
			order[i] = i
		}
		sort.Slice(order, func(i, j int) bool {
			gi := combatGroups[order[i]]
			gj := combatGroups[order[j]]
			if gi.effectivePower() != gj.effectivePower() {
				return gi.effectivePower() > gj.effectivePower()
			}
			return gi.initiative > gj.initiative
		})
		
		// Each group selects a target
		for _, attackerIdx := range order {
			attacker := combatGroups[attackerIdx]
			if attacker.units <= 0 {
				continue
			}
			
			bestTarget := -1
			bestDamage := 0
			
			for defenderIdx, defender := range combatGroups {
				if defender.units <= 0 || defender.army == attacker.army || targeted[defenderIdx] {
					continue
				}
				
				damage := attacker.damageDealtTo(defender)
				if damage == 0 {
					continue
				}
				
				// Choose target based on damage, then effective power, then initiative
				if damage > bestDamage ||
					(damage == bestDamage && (bestTarget == -1 || 
						defender.effectivePower() > combatGroups[bestTarget].effectivePower() ||
						(defender.effectivePower() == combatGroups[bestTarget].effectivePower() && 
							defender.initiative > combatGroups[bestTarget].initiative))) {
					bestDamage = damage
					bestTarget = defenderIdx
				}
			}
			
			if bestTarget != -1 {
				targets[attackerIdx] = bestTarget
				targeted[bestTarget] = true
			}
		}
		
		// Check for stalemate
		if len(targets) == 0 {
			break
		}
		
		// Attack phase (by initiative desc)
		attackOrder := make([]int, 0, len(targets))
		for attacker := range targets {
			attackOrder = append(attackOrder, attacker)
		}
		sort.Slice(attackOrder, func(i, j int) bool {
			return combatGroups[attackOrder[i]].initiative > combatGroups[attackOrder[j]].initiative
		})
		
		totalKilled := 0
		for _, attackerIdx := range attackOrder {
			attacker := combatGroups[attackerIdx]
			if attacker.units <= 0 {
				continue
			}
			
			defenderIdx := targets[attackerIdx]
			defender := combatGroups[defenderIdx]
			
			damage := attacker.damageDealtTo(defender)
			unitsKilled := damage / defender.hp
			if unitsKilled > defender.units {
				unitsKilled = defender.units
			}
			defender.units -= unitsKilled
			totalKilled += unitsKilled
		}
		
		// Check for stalemate (no units killed)
		if totalKilled == 0 {
			break
		}
	}
	
	return combatGroups
}

func Part1(input string) string {
	groups := parseInput(input)
	result := fight(groups)
	
	totalUnits := 0
	for _, g := range result {
		if g.units > 0 {
			totalUnits += g.units
		}
	}
	
	return fmt.Sprintf("%d", totalUnits)
}

func fightWithBoost(groups []*Group, boost int) ([]*Group, bool) {
	// Make copies and apply boost
	combatGroups := make([]*Group, len(groups))
	for i, g := range groups {
		newGroup := *g
		if newGroup.army == "Immune System" {
			newGroup.damage += boost
		}
		combatGroups[i] = &newGroup
	}
	
	rounds := 0
	maxRounds := 10000 // Prevent infinite loops
	
	for rounds < maxRounds {
		rounds++
		
		// Check if combat is over
		immuneCount := 0
		infectionCount := 0
		for _, g := range combatGroups {
			if g.units > 0 {
				if g.army == "Immune System" {
					immuneCount++
				} else {
					infectionCount++
				}
			}
		}
		if immuneCount == 0 || infectionCount == 0 {
			return combatGroups, immuneCount > 0
		}
		
		// Target selection phase
		targets := make(map[int]int)
		targeted := make(map[int]bool)
		
		// Sort for target selection
		order := make([]int, len(combatGroups))
		for i := range order {
			order[i] = i
		}
		sort.Slice(order, func(i, j int) bool {
			gi := combatGroups[order[i]]
			gj := combatGroups[order[j]]
			if gi.effectivePower() != gj.effectivePower() {
				return gi.effectivePower() > gj.effectivePower()
			}
			return gi.initiative > gj.initiative
		})
		
		// Each group selects a target
		for _, attackerIdx := range order {
			attacker := combatGroups[attackerIdx]
			if attacker.units <= 0 {
				continue
			}
			
			bestTarget := -1
			bestDamage := 0
			
			for defenderIdx, defender := range combatGroups {
				if defender.units <= 0 || defender.army == attacker.army || targeted[defenderIdx] {
					continue
				}
				
				damage := attacker.damageDealtTo(defender)
				if damage == 0 {
					continue
				}
				
				if damage > bestDamage ||
					(damage == bestDamage && (bestTarget == -1 || 
						defender.effectivePower() > combatGroups[bestTarget].effectivePower() ||
						(defender.effectivePower() == combatGroups[bestTarget].effectivePower() && 
							defender.initiative > combatGroups[bestTarget].initiative))) {
					bestDamage = damage
					bestTarget = defenderIdx
				}
			}
			
			if bestTarget != -1 {
				targets[attackerIdx] = bestTarget
				targeted[bestTarget] = true
			}
		}
		
		// Attack phase
		attackOrder := make([]int, 0, len(targets))
		for attacker := range targets {
			attackOrder = append(attackOrder, attacker)
		}
		sort.Slice(attackOrder, func(i, j int) bool {
			return combatGroups[attackOrder[i]].initiative > combatGroups[attackOrder[j]].initiative
		})
		
		totalKilled := 0
		for _, attackerIdx := range attackOrder {
			attacker := combatGroups[attackerIdx]
			if attacker.units <= 0 {
				continue
			}
			
			defenderIdx := targets[attackerIdx]
			defender := combatGroups[defenderIdx]
			
			damage := attacker.damageDealtTo(defender)
			unitsKilled := damage / defender.hp
			if unitsKilled > defender.units {
				unitsKilled = defender.units
			}
			defender.units -= unitsKilled
			totalKilled += unitsKilled
		}
		
		// Check for stalemate
		if totalKilled == 0 {
			return combatGroups, false
		}
	}
	
	return combatGroups, false
}

func Part2(input string) string {
	groups := parseInput(input)
	
	// Binary search for minimum boost
	low, high := 0, 10000
	minBoost := high
	minUnits := 0
	
	for low <= high {
		mid := (low + high) / 2
		result, immuneWins := fightWithBoost(groups, mid)
		
		if immuneWins {
			// Count remaining immune units
			units := 0
			for _, g := range result {
				if g.army == "Immune System" && g.units > 0 {
					units += g.units
				}
			}
			
			if mid < minBoost || (mid == minBoost && units < minUnits) {
				minBoost = mid
				minUnits = units
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	
	return fmt.Sprintf("%d", minUnits)
}