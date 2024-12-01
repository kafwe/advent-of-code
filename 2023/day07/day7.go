package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

type handStrength int

const (
	Unknown handStrength = iota
	HighCard
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type hand struct {
	cards      string
	bid        int
	kind       handStrength
	jokerCards string
}

var cardStrength = map[byte]int{
	'A': 13, 'K': 12, 'Q': 11, 'T': 9, '9': 8, '8': 7,
	'7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1, 'J': 0,
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing input file argument")
	}
	file := os.Args[1]
	hands := readHands(file)

	slices.SortStableFunc(hands, compareHands)
	fmt.Println("Part 1:", calculateWinnings(hands))
	slices.SortStableFunc(hands, compareJokerHands)
	fmt.Println("Part 2:", calculateWinnings(hands))
}

func calculateWinnings(hands []hand) int {
	winnings := 0
	for rank, h := range hands {
		winnings += h.bid * (rank + 1)
	}
	return winnings
}

func compareHands(hand1, hand2 hand) int {
	hand1.kind, hand2.kind = determineHandKind(hand1.cards), determineHandKind(hand2.cards)

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

func compareJokerHands(hand1, hand2 hand) int {
	// turn 'J' card into Joker
	cardStrength['J'] = 0
	hand1.processJoker()
	hand2.processJoker()
	hand1.kind, hand2.kind = determineHandKind(hand1.jokerCards), determineHandKind(hand2.jokerCards)

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

func determineHandKind(hand string) handStrength {
	labels := make(map[rune]int)
	for _, c := range hand {
		labels[c]++
	}

	switch len(labels) {
	case 5:
		return HighCard
	case 4:
		return OnePair
	case 3:
		for _, n := range labels {
			switch n {
			case 3:
				return ThreeOfAKind
			case 2:
				return TwoPair
			}
		}
	case 2:
		for _, n := range labels {
			switch n {
			case 1:
				return FourOfAKind
			case 2:
				return FullHouse
			}
		}
	case 1:
		return FiveOfAKind
	}
	return Unknown
}

func (h *hand) processJoker() {
	labels := make(map[rune]int)
	for _, c := range h.cards {
		labels[c]++
	}

	maxCount := math.MinInt64
	var label rune
	for l, n := range labels {
		if n > maxCount && l != 'J' {
			maxCount = n
			label = l
		}
	}
	h.jokerCards = strings.ReplaceAll(h.cards, "J", string(label))
}
