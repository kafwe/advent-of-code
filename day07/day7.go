package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type handStrength int

const (
	HighCard handStrength = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type hand struct {
	cards string
	bid   int
	kind  handStrength
}

func main() {
	hands := readHands("./day07/input.txt")
	slices.SortStableFunc(hands, compareHands)

	winnings := 0
	for rank, h := range hands {
		winnings += h.bid * (rank + 1)
	}
	fmt.Println("Part 1:", winnings)
}

func compareHands(hand1, hand2 hand) int {
	cardStrength := map[byte]int{
		'A': 13, 'K': 12, 'Q': 11, 'J': 10, 'T': 9, '9': 8,
		'8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1,
	}
	hand1.kind, hand2.kind = determineHandKind(hand1), determineHandKind(hand2)

	switch {
	case hand1.kind == hand2.kind:
		for i := range hand1.cards {
			if cardStrength[hand1.cards[i]] > cardStrength[hand2.cards[i]] {
				return 1
			} else if cardStrength[hand1.cards[i]] < cardStrength[hand2.cards[i]] {
				return -1
			}
		}
	case hand1.kind > hand2.kind:
		return 1
	case hand1.kind < hand2.kind:
		return -1
	}
	return 0
}

func readHands(filePath string) []hand {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var (
		hands []hand
		hand  hand
	)
	for scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "%s %d", &hand.cards, &hand.bid)
		if err != nil {
			log.Fatal(err)
		}
		hands = append(hands, hand)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return hands
}

func determineHandKind(h hand) handStrength {
	labels := make(map[rune]int)
	for _, c := range h.cards {
		labels[c]++
	}
	var kind handStrength

	switch len(labels) {
	case 5:
		kind = HighCard
	case 4:
		kind = OnePair
	case 3:
		for _, n := range labels {
			if n == 3 {
				kind = ThreeOfAKind
			}
			if n == 2 {
				kind = TwoPair
			}
		}
	case 2:
		for _, n := range labels {
			if n == 1 {
				kind = FourOfAKind
			}
			if n == 2 {
				kind = FullHouse
			}
		}
	case 1:
		kind = FiveOfAKind
	}
	return kind
}
