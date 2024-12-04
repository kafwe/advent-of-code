package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/kafwe/advent-of-code/fileio"
)

func main() {
	lines, err := fileio.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`mul\((\d{1,3}),\s*(\d{1,3})\)`)
	sum := 0
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			num1, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println("Could not convert number:", err)
			}
			num2, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Println("Could not convert number:", err)
			}
			sum += num1 * num2
		}

	}
	fmt.Println("Part 1:", sum)
}
