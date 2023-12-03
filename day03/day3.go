package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing input file argument")
	}
	file := os.Args[1]
	schematic := readSchematic(file)

	fmt.Println("Part 1:", part1(schematic))
	fmt.Println("Part 2:", part2(schematic))
}

func part1(schematic []string) int {
	var partNumbers []int
	for i, line := range schematic {
		adj := false
		num := ""
		for j, char := range line {
			if unicode.IsDigit(char) {
				num += string(char)

				if !adj {
					adj = checkAdjacency(schematic, i, j)
				}
			}
			if !unicode.IsDigit(char) || j == len(line)-1 {
				if adj {
					n, err := strconv.Atoi(num)
					if err != nil {
						log.Fatal(err)
					}
					partNumbers = append(partNumbers, n)
				}
				num = ""
				adj = false
			}
		}
	}

	sum := 0
	for _, n := range partNumbers {
		sum += n
	}
	return sum
}

func part2(schematic []string) int {
	var gearRatios []int
	for r, line := range schematic {
		for c, char := range line {
			if char == '*' {
				gearRatios = append(gearRatios, findGearRatio(schematic, r, c))
			}
		}
	}

	sum := 0
	for _, n := range gearRatios {
		sum += n
	}
	return sum
}

func checkAdjacency(schematic []string, row, col int) bool {
	startRow := row - 1
	endRow := row + 1
	startCol := col - 1
	endCol := col + 1

	if row == 0 {
		startRow = 0
	}

	if row == len(schematic)-1 {
		endRow = len(schematic) - 1
	}

	if col == 0 {
		startCol = 0
	}

	if col == len(schematic[row])-1 {
		endCol = len(schematic[row]) - 1
	}

	// Iterate over the surrounding square
	for r := startRow; r <= endRow; r++ {
		for c := startCol; c <= endCol; c++ {
			if schematic[r][c] != '.' && !unicode.IsDigit(rune(schematic[r][c])) {
				return true
			}
		}
	}

	return false
}

func findGearRatio(schematic []string, row, col int) int {
	startRow := row - 4
	endRow := row + 4
	startCol := col - 4
	endCol := col + 4

	if startRow < 0 {
		startRow = 0
	}

	if endRow > len(schematic)-1 {
		endRow = len(schematic) - 1
	}

	if startCol < 0 {
		startCol = 0
	}

	if endCol > len(schematic[row])-1 {
		endCol = len(schematic[row]) - 1
	}

	partNums := []int{}
	// Iterate over the surrounding square
	for r := startRow; r <= endRow; r++ {
		var num string
		var start int
		for c := startCol; c <= endCol; c++ {
			if unicode.IsDigit(rune(schematic[r][c])) {
				if len(num) == 0 {
					start = c
				}
				num += string(schematic[r][c])
			}

			if !unicode.IsDigit(rune(schematic[r][c])) || r == len(schematic[c])-1 {
				touchesEnd := math.Abs(float64(c-1-col)) <= 1 && math.Abs(float64(r-row)) <= 1
				touchesStart := math.Abs(float64(start-col)) <= 1 && math.Abs(float64(r-row)) <= 1
				if (touchesStart || touchesEnd) && len(num) > 0 {
					n, err := strconv.Atoi(num)
					if err != nil {
						log.Fatal(err)
					}
					partNums = append(partNums, n)
				}
				num = ""
			}
		}
	}

	if len(partNums) < 2 {
		return 0
	}

	sum := 1
	for _, n := range partNums {
		sum *= n
	}

	return sum
}

func readSchematic(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	schematic := []string{}
	for scanner.Scan() {
		schematic = append(schematic, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return schematic
}
