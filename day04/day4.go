package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type card struct {
	winningNumbers []string
	userNumbers    []string
}

func main() {
	/*
		if len(os.Args) < 2 {
			log.Fatal("Missing input file argument")
		}
		file := os.Args[1]

	*/
	cards := readCards("./day04/input.txt")

	fmt.Println("Part 1:", part1(cards))
}

func part1(cards []card) int {
	points := 0
	for _, card := range cards {
		fmt.Println(card)
		numMatches := countMatches(card.winningNumbers, card.userNumbers)
		points += calculatePoints(numMatches)
		fmt.Println("Points:", points, "Matches:", numMatches)
	}
	return points
}

func countMatches(winningNumbers, userNumbers []string) int {
	matches := 0
	seen := make(map[string]bool)

	for _, num := range winningNumbers {
		seen[num] = true
	}

	for _, num := range userNumbers {
		if seen[num] && num != "" {
			seen[num] = false
			matches++
		}
	}

	return matches
}

func calculatePoints(numMatches int) int {
	if numMatches == 0 {
		return 0
	}
	return int(math.Pow(2, float64(numMatches-1)))
}

func readCards(filePath string) []card {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cards []card
	for scanner.Scan() {
		_, parts, _ := strings.Cut(scanner.Text(), ": ")
		nums := strings.Split(parts, "|")
		winNums := strings.Split(strings.TrimSpace(nums[0]), " ")
		userNums := strings.Split(strings.TrimSpace(nums[1]), " ")
		cards = append(cards, card{winNums, userNums})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return cards
}
