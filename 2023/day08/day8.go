package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type node struct {
	left  string
	right string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing input file argument")
	}
	file := os.Args[1]
	instructions, nodes, startNodes := readNodes(file)

	fmt.Println("Part 1:", part1(instructions, nodes, "AAA"))
	fmt.Println("Part 2:", part2(instructions, nodes, startNodes))
}

func part1(instructions string, nodes map[string]node, startNode string) int {
	steps := 0
	currNode := startNode
	for {
		for _, i := range instructions {
			if currNode[2] == 'Z' {
				return steps
			}
			if i == 'L' {
				currNode = nodes[currNode].left
			} else {
				currNode = nodes[currNode].right
			}
			steps++
		}
	}
}

func part2(instructions string, nodes map[string]node, startNodes []string) int {
	var steps []int
	for _, startNode := range startNodes {
		steps = append(steps, part1(instructions, nodes, startNode))
	}
	return lcmOfSlice(steps)
}

func readNodes(filePath string) (string, map[string]node, []string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := scanner.Text()
	pattern := regexp.MustCompile(`(\w+)\s*=\s*\((\w+),\s*(\w+)\)`)

	nodes := make(map[string]node)
	var startNodes []string
	var value, left, right string
	for scanner.Scan() {
		raw := scanner.Text()
		if len(raw) > 0 {
			matches := pattern.FindStringSubmatch(raw)
			value, left, right = matches[1], matches[2], matches[3]
			if value[2] == 'A' {
				startNodes = append(startNodes, value)
			}
			nodes[value] = node{left, right}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return instructions, nodes, startNodes
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmOfSlice(nums []int) int {
	result := 1
	for _, n := range nums {
		result = lcm(result, n)
	}
	return result
}
