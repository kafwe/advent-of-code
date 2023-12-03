package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	/*
		if len(os.Args) < 2 {
			log.Fatal("Missing input file argument")
		}
		file := os.Args[1]
	*/
	schematic := readSchematic("./day03/input.txt")

	fmt.Println("Part 1:", part1(schematic))
}

func part1(schematic []string) int {
	partNumbers := []int{}
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

	fmt.Println(partNumbers)

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
	for i := startRow; i <= endRow; i++ {
		for j := startCol; j <= endCol; j++ {
			// Check for non-digit and non-period characters
			if schematic[i][j] != '.' && !unicode.IsDigit(rune(schematic[i][j])) {
				return true
			}
		}
	}

	return false
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
