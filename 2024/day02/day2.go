package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/kafwe/advent-of-code/fileio"
)

func main() {
	lines, err := fileio.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reports := make([][]int, 1000)

	for i, line := range lines {
		parts := strings.Split(line, " ")
		levels := make([]int, len(parts))

		for j, p := range parts {
			val, err := strconv.Atoi(p)
			if err != nil {
				fmt.Println("Error converting part:", err)
				continue
			}
			levels[j] = val
		}
		reports[i] = levels
	}

	fmt.Println(reports)
	fmt.Println("Part 1:", part1(reports))
}

func part1(reports [][]int) int {
	count := 0
	for _, report := range reports {
		var increasing bool
		safe := true
		for l, _ := range report {
			if l == len(report)-1 {
				continue
			}

			r := l + 1
			diff := report[l] - report[r]
			if l == 0 {
				increasing = isIncreasing(diff)
			}

			if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
				safe = false
			}

			if increasing != isIncreasing(diff) {
				safe = false
			}
		}

		if safe {
			count++
		}
	}
	return count
}

func isIncreasing(n int) bool {
	if n >= 1 {
		return false
	} else if n < 0 {
		return true
	}
	return false
}
