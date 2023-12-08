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
	/*
		if len(os.Args) < 2 {
			log.Fatal("Missing input file argument")
		}
		file := os.Args[1]
	*/
	instructions, nodes := readNodes("./day08/input.txt")

	fmt.Println("Part 1:", part1(instructions, nodes))
}

func part1(instructions string, nodes map[string]node) int {
	steps := 0
	currNode := "AAA"
	for {
		for _, i := range instructions {
			steps++
			if i == 'L' {
				if nodes[currNode].left == "ZZZ" {
					return steps
				}
				currNode = nodes[currNode].left
			} else {
				if nodes[currNode].right == "ZZZ" {
					return steps
				}
				currNode = nodes[currNode].right
			}
		}
	}
}

func readNodes(filePath string) (string, map[string]node) {
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
	var value, left, right string
	for scanner.Scan() {
		raw := scanner.Text()
		fmt.Println(raw)
		if len(raw) > 0 {
			matches := pattern.FindStringSubmatch(raw)
			value, left, right = matches[1], matches[2], matches[3]
			nodes[value] = node{left, right}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return instructions, nodes
}
