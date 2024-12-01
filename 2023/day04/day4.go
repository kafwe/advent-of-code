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
	quantity       int
	numMatches     int
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing input file argument")
	}
	file := os.Args[1]
	cards := readCards(file)

	points := 0
	for i := range cards {
		card := &cards[i]
		countMatches(card)
		for j := 0; j < card.quantity; j++ {
			for k := 1; k <= card.numMatches; k++ {
				cards[k+i].quantity++
			}
		}
		points += calcPoints(card.numMatches)
	}

	total := 0
	for _, c := range cards {
		total += c.quantity
	}

	fmt.Println("Total points:", points)
	fmt.Println("Total instances:", total)
}

func countMatches(card *card) int {
	matches := 0
	seen := make(map[string]bool)

	for _, num := range card.winningNumbers {
		seen[num] = true
	}

	for _, num := range card.userNumbers {
		if seen[num] {
			seen[num] = false // avoid duplicates
			card.numMatches++
		}
	}

	return matches
}

func calcPoints(numMatches int) int {
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
		cards = append(cards, card{strings.Fields(nums[0]), strings.Fields(nums[1]), 1, 0})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return cards
}
