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

		l[i] = val1
		r[i] = val2
	}

	sort.Ints(l)
	sort.Ints(r)

	var dist float64

	for i, _ := range l {
		dist += math.Abs(float64(l[i]) - float64(r[i]))
	}

	fmt.Println(int(dist))

}

func part1() int {

}
