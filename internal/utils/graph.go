package utils

import "container/heap"

// PriorityQueue implements heap.Interface for A* and Dijkstra
type Item struct {
	Value    interface{}
	Priority int
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

// BFS performs breadth-first search
func BFS(start Point, isTarget func(Point) bool, getNeighbors func(Point) []Point) ([]Point, bool) {
	queue := []Point{start}
	visited := make(map[Point]bool)
	parent := make(map[Point]Point)

	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if isTarget(current) {
			// Reconstruct path
			path := []Point{}
			for p := current; p != start; p = parent[p] {
				path = append([]Point{p}, path...)
			}
			path = append([]Point{start}, path...)
			return path, true
		}

		for _, neighbor := range getNeighbors(current) {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = current
				queue = append(queue, neighbor)
			}
		}
	}

	return nil, false
}

// DFS performs depth-first search
func DFS(start Point, isTarget func(Point) bool, getNeighbors func(Point) []Point) bool {
	visited := make(map[Point]bool)

	var dfsHelper func(Point) bool
	dfsHelper = func(current Point) bool {
		if isTarget(current) {
			return true
		}

		visited[current] = true

		for _, neighbor := range getNeighbors(current) {
			if !visited[neighbor] {
				if dfsHelper(neighbor) {
					return true
				}
			}
		}

		return false
	}

	return dfsHelper(start)
}

// Dijkstra finds shortest path using Dijkstra's algorithm
func Dijkstra(start Point, isTarget func(Point) bool, getNeighbors func(Point) []PointWithCost) (int, []Point) {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	dist := make(map[Point]int)
	parent := make(map[Point]Point)

	dist[start] = 0
	heap.Push(&pq, &Item{Value: start, Priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		current := item.Value.(Point)
		currentDist := item.Priority

		if isTarget(current) {
			// Reconstruct path
			path := []Point{}
			for p := current; p != start; p = parent[p] {
				path = append([]Point{p}, path...)
			}
			path = append([]Point{start}, path...)
			return currentDist, path
		}

		if d, ok := dist[current]; ok && currentDist > d {
			continue
		}

		for _, nc := range getNeighbors(current) {
			newDist := currentDist + nc.Cost
			if d, ok := dist[nc.Point]; !ok || newDist < d {
				dist[nc.Point] = newDist
				parent[nc.Point] = current
				heap.Push(&pq, &Item{Value: nc.Point, Priority: newDist})
			}
		}
	}

	return -1, nil
}

type PointWithCost struct {
	Point Point
	Cost  int
}

// TopologicalSort performs topological sorting on a DAG
func TopologicalSort(nodes []string, edges map[string][]string) ([]string, bool) {
	inDegree := make(map[string]int)
	for _, node := range nodes {
		inDegree[node] = 0
	}

	for _, neighbors := range edges {
		for _, neighbor := range neighbors {
			inDegree[neighbor]++
		}
	}

	queue := []string{}
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	result := []string{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		for _, neighbor := range edges[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return result, len(result) == len(nodes)
}

// FloodFill performs flood fill from a starting point
func FloodFill(grid *Grid, start Point, fillValue rune) int {
	if !grid.InBounds(start) {
		return 0
	}

	originalValue := grid.Get(start)
	if originalValue == fillValue {
		return 0
	}

	count := 0
	queue := []Point{start}
	visited := make(map[Point]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] || !grid.InBounds(current) {
			continue
		}

		if grid.Get(current) != originalValue {
			continue
		}

		visited[current] = true
		grid.Set(current, fillValue)
		count++

		for _, neighbor := range current.Neighbors4() {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}

	return count
}
