/*
 * Day 8: Memory Maneuver
 *
 * Part 1: Parse a tree structure from a flat list of numbers and sum all metadata entries.
 * Each node has a header with child count and metadata count, followed by children and metadata.
 * The tree is encoded in a specific format that needs recursive parsing.
 *
 * Part 2: Calculate the value of nodes based on specific rules involving metadata as indices.
 * Nodes with no children have value equal to sum of metadata. Nodes with children use
 * metadata as 1-indexed references to child values.
 */

package day08

import (
	"strconv"
	"strings"
)

type Node struct {
	children []Node
	metadata []int
}

type Solution struct {
	numbers []int
}

func New(input string) *Solution {
	input = strings.ReplaceAll(strings.TrimSpace(input), "\r\n", "\n")
	fields := strings.Fields(input)

	numbers := make([]int, len(fields))
	for i, field := range fields {
		num, _ := strconv.Atoi(field)
		numbers[i] = num
	}

	return &Solution{numbers: numbers}
}

func (s *Solution) Part1() (int, error) {
	root, _ := s.parseNode(0)
	return s.sumMetadata(root), nil
}

func (s *Solution) Part2() (int, error) {
	root, _ := s.parseNode(0)
	return s.nodeValue(root), nil
}

// parseNode parses a node starting at the given index and returns the node and next index
func (s *Solution) parseNode(index int) (Node, int) {
	childCount := s.numbers[index]
	metadataCount := s.numbers[index+1]
	index += 2

	var node Node
	node.children = make([]Node, childCount)

	// Parse children
	for i := 0; i < childCount; i++ {
		child, nextIndex := s.parseNode(index)
		node.children[i] = child
		index = nextIndex
	}

	// Parse metadata
	node.metadata = make([]int, metadataCount)
	for i := 0; i < metadataCount; i++ {
		node.metadata[i] = s.numbers[index]
		index++
	}

	return node, index
}

// sumMetadata recursively sums all metadata in the tree
func (s *Solution) sumMetadata(node Node) int {
	sum := 0

	// Add this node's metadata
	for _, meta := range node.metadata {
		sum += meta
	}

	// Add children's metadata
	for _, child := range node.children {
		sum += s.sumMetadata(child)
	}

	return sum
}

// nodeValue calculates the value of a node according to Part 2 rules
func (s *Solution) nodeValue(node Node) int {
	if len(node.children) == 0 {
		// No children: value is sum of metadata
		sum := 0
		for _, meta := range node.metadata {
			sum += meta
		}
		return sum
	}

	// Has children: metadata are 1-indexed references to children
	sum := 0
	for _, meta := range node.metadata {
		if meta >= 1 && meta <= len(node.children) {
			sum += s.nodeValue(node.children[meta-1])
		}
	}
	return sum
}
