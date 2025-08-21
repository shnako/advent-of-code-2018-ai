package day23

import (
	"container/heap"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Nanobot struct {
	x, y, z int
	radius  int
}

func ParseInput(input string) []Nanobot {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	nanobots := make([]Nanobot, 0, len(lines))
	
	re := regexp.MustCompile(`pos=<(-?\d+),(-?\d+),(-?\d+)>, r=(\d+)`)
	
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) == 5 {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			z, _ := strconv.Atoi(matches[3])
			r, _ := strconv.Atoi(matches[4])
			nanobots = append(nanobots, Nanobot{x, y, z, r})
		}
	}
	
	return nanobots
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(n1, n2 Nanobot) int {
	return abs(n1.x-n2.x) + abs(n1.y-n2.y) + abs(n1.z-n2.z)
}

func Part1(input string) (string, error) {
	nanobots := ParseInput(input)
	
	if len(nanobots) == 0 {
		return "", fmt.Errorf("no nanobots found")
	}
	
	strongestIdx := 0
	maxRadius := nanobots[0].radius
	for i := 1; i < len(nanobots); i++ {
		if nanobots[i].radius > maxRadius {
			maxRadius = nanobots[i].radius
			strongestIdx = i
		}
	}
	
	strongest := nanobots[strongestIdx]
	count := 0
	for _, bot := range nanobots {
		if manhattanDistance(strongest, bot) <= strongest.radius {
			count++
		}
	}
	
	return strconv.Itoa(count), nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Box struct {
	x, y, z int
	size    int
}


func distanceToBox(bot Nanobot, box Box) int {
	dist := 0
	
	if bot.x < box.x {
		dist += box.x - bot.x
	} else if bot.x >= box.x+box.size {
		dist += bot.x - (box.x + box.size - 1)
	}
	
	if bot.y < box.y {
		dist += box.y - bot.y
	} else if bot.y >= box.y+box.size {
		dist += bot.y - (box.y + box.size - 1)
	}
	
	if bot.z < box.z {
		dist += box.z - bot.z
	} else if bot.z >= box.z+box.size {
		dist += bot.z - (box.z + box.size - 1)
	}
	
	return dist
}

func countBotsInRangeOfBox(box Box, nanobots []Nanobot) int {
	count := 0
	for _, bot := range nanobots {
		if distanceToBox(bot, box) <= bot.radius {
			count++
		}
	}
	return count
}

type Item struct {
	box      Box
	count    int
	distance int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].count != pq[j].count {
		return pq[i].count > pq[j].count
	}
	if pq[i].distance != pq[j].distance {
		return pq[i].distance < pq[j].distance
	}
	return pq[i].box.size < pq[j].box.size
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func Part2(input string) (string, error) {
	nanobots := ParseInput(input)
	
	if len(nanobots) == 0 {
		return "", fmt.Errorf("no nanobots found")
	}
	
	minCoord := nanobots[0].x
	maxCoord := nanobots[0].x
	
	for _, bot := range nanobots {
		minCoord = min(minCoord, min(bot.x, min(bot.y, bot.z)))
		maxCoord = max(maxCoord, max(bot.x, max(bot.y, bot.z)))
	}
	
	boxSize := 1
	for boxSize < maxCoord-minCoord {
		boxSize *= 2
	}
	
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	
	initialBox := Box{minCoord, minCoord, minCoord, boxSize}
	initialCount := countBotsInRangeOfBox(initialBox, nanobots)
	initialDist := abs(minCoord) + abs(minCoord) + abs(minCoord)
	heap.Push(&pq, &Item{initialBox, initialCount, initialDist, 0})
	
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		
		if item.box.size == 1 {
			return strconv.Itoa(item.distance), nil
		}
		
		newSize := item.box.size / 2
		if newSize == 0 {
			newSize = 1
		}
		
		for dx := 0; dx < 2; dx++ {
			for dy := 0; dy < 2; dy++ {
				for dz := 0; dz < 2; dz++ {
					newBox := Box{
						item.box.x + dx*newSize,
						item.box.y + dy*newSize,
						item.box.z + dz*newSize,
						newSize,
					}
					
					count := countBotsInRangeOfBox(newBox, nanobots)
					dist := abs(newBox.x) + abs(newBox.y) + abs(newBox.z)
					
					heap.Push(&pq, &Item{newBox, count, dist, 0})
				}
			}
		}
	}
	
	return "", fmt.Errorf("no solution found")
}