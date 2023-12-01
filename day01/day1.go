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
	fmt.Println("Part 1:", part1())
}

func part1() int {
	file, err := os.Open("./day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		total += getNumber(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return total
}

func getNumber(line string) int {
	digits := ""
	for _, char := range line {
		if unicode.IsDigit(char) {
			digits += string(char)
		}
	}

	number, err := strconv.Atoi(string(digits[0]) + string(digits[len(digits)-1]))
	if err != nil {
		log.Fatal(err)
	}

	return number
}
