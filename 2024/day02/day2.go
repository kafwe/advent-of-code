package main

import (
	"fmt"
	"log"
	"math"
	"slices"
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

	fmt.Println("Part 1:", countSafeReports(reports, false))
	fmt.Println("Part 2:", countSafeReports(reports, true))
}

func countSafeReports(reports [][]int, dampen bool) int {
	count := 0
	for _, report := range reports {
		for i, _ := range report {
			safe := isSafe(report)
			if safe {
				count++
				break
			}

			if dampen {
				dampened := slices.Delete(slices.Clone(report), i, i+1)
				safe = isSafe(dampened)
				if safe {
					count++
					break
				}
			}
		}
	}
	return count
}

func isSafe(report []int) bool {
	diff := report[1] - report[0]
	increasing := diff < 0

	for i := 1; i < len(report); i++ {
		diff = report[i] - report[i-1]
		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			return false
		}
		if diff < 0 != increasing {
			return false
		}
	}
	return true
}
