package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	fmt.Printf("part 1: %d\n", part1(lines))
	fmt.Printf("part 2: %d\n", part2(lines))
}

type Node struct {
	Name  string
	Left  string
	Right string
}

func part1(lines []string) int {
	instructions := lines[0]
	nodes := parseNodes(lines[2:])
	network := makeNetworkMap(nodes)

	steps := countStepsToFinalNode(network, instructions, "AAA",
		func(node string) bool {
			return node != "ZZZ"
		})
	return steps
}

func part2(lines []string) int {
	instructions := lines[0]
	nodes := parseNodes(lines[2:])
	network := makeNetworkMap(nodes)

	currentNodes := getAllNodesEndingWithA(nodes)
	nodeSteps := make([]int, 0, len(currentNodes))

	for _, n := range currentNodes {
		s := countStepsToFinalNode(network, instructions, n,
			func(node string) bool {
				return node[len(node)-1] != 'Z'
			})
		nodeSteps = append(nodeSteps, s)
	}

	uniqueSteps := uniqueIntSlice(nodeSteps)
	return leastCommonMultipleOfSlice(uniqueSteps)
}

func parseNodes(lines []string) []Node {
	nodes := make([]Node, 0, len(lines))
	re := regexp.MustCompile(`([1-9A-Z]{3}) = \(([1-9A-Z]{3}), ([1-9A-Z]{3})\)`)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		n := Node{
			Name:  matches[1],
			Left:  matches[2],
			Right: matches[3],
		}
		nodes = append(nodes, n)
	}

	return nodes
}

func makeNetworkMap(nodes []Node) map[string][]string {
	networkMap := make(map[string][]string)
	for _, n := range nodes {
		networkMap[n.Name] = []string{n.Left, n.Right}
	}
	return networkMap
}

func countStepsToFinalNode(network map[string][]string, instructions, startNode string, loopCondition func(string) bool) int {
	currentNode := startNode
	steps := 0

	for loopCondition(currentNode) {
		nodePaths, ok := network[currentNode]
		if !ok {
			fmt.Printf("Error: Node %s not found! Stopping at step %d.\n", currentNode, steps)
			break
		}
		steps++
		left, right := nodePaths[0], nodePaths[1]
		currentInstruction := instructions[(steps-1)%len(instructions)]
		if currentInstruction == 'L' {
			currentNode = left
		} else if currentInstruction == 'R' {
			currentNode = right
		} else {
			log.Fatal(fmt.Sprintf("Cannot read instruction %s at step %d\n", string(currentInstruction), steps))
		}
	}
	return steps
}

func getAllNodesEndingWithA(nodes []Node) []string {
	res := make([]string, 0, len(nodes))
	for _, n := range nodes {
		if n.Name[len(n.Name)-1] == 'A' {
			res = append(res, n.Name)
		}
	}

	return res
}

func areAllStringsEndingWithZ(strs []string) bool {
	for _, s := range strs {
		if s[len(s)-1] != 'Z' {
			return false
		}
	}

	return true
}

func uniqueIntSlice(slice []int) []int {
	seen := make(map[int]bool)
	var uniques []int

	for _, v := range slice {
		if _, found := seen[v]; !found {
			seen[v] = true
			uniques = append(uniques, v)
		}
	}

	return uniques
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}
	return a
}

func leastCommonMultiple(a, b int) int {
	return (a * b) / greatestCommonDivisor(a, b)
}

func leastCommonMultipleOfSlice(nums []int) int {
	n := nums[0]
	for i := 1; i < len(nums); i++ {
		n = leastCommonMultiple(n, nums[i])
	}

	return n
}
