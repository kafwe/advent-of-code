package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/kafwe/advent-of-code/fileio"
)

func main() {
	lines, err := fileio.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	l, r := make([]int, 1000), make([]int, 1000)

	for i, line := range lines {
		parts := strings.Split(line, "   ")

		val1, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error converting parts[0]:", err)
			continue
		}
		val2, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error converting parts[1]:", err)
			continue
		}

		l[i], r[i] = val1, val2
	}

	fmt.Println("Part 1:", part1(l, r))
	fmt.Println("Part 2:", part2(l, r))

}

func part1(l []int, r []int) int {
	sort.Ints(l)
	sort.Ints(r)

	var dist float64
	for i, _ := range l {
		dist += math.Abs(float64(l[i]) - float64(r[i]))
	}
	return int(dist)
}

func part2(l []int, r []int) int {
	left := make(map[int]int)
	right := make(map[int]int)

	for _, num := range l {
		_, checked := left[num]
		left[num]++
		// already counted num occurences in right list
		if checked {
			continue
		}

		for _, val := range r {
			if val == num {
				right[num]++
			}
		}

	}

	score := 0
	for k, v := range left {
		for i := 0; i < v; i++ {
			score += right[k] * k
		}
	}

	return score

}
