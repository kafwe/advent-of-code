package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rawTime := scanner.Text()

	scanner.Scan()
	rawDist := scanner.Text()

	_, cleanTime, _ := strings.Cut(rawTime, ":")
	_, cleanDist, _ := strings.Cut(rawDist, ":")
	times := strings.Fields(cleanTime)
	dists := strings.Fields(cleanDist)

	numWays := []int{}
	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		dist, _ := strconv.Atoi(dists[i])
		numWays = append(numWays, calcNumWays(time, dist))
	}

	part1 := 1
	for _, n := range numWays {
		part1 *= n
	}
	fmt.Println("Part 1:", part1)

	time := convToOneRace(times)
	dist := convToOneRace(dists)
	fmt.Println("Part 2:", calcNumWays(time, dist))
}

func calcNumWays(time, dist int) int {
	numWays := 0
	for i := 0; i < time; i++ {
		speed := i
		remainingTime := time - i
		calcDist := speed * remainingTime
		if calcDist > dist {
			numWays++
		}
	}
	return numWays
}

func convToOneRace(vals []string) int {
	val := ""
	for _, v := range vals {
		val += v
	}
	num, _ := strconv.Atoi(val)
	return num
}
