package main

import (
	"embed"
	"log"
	"sort"
	"strings"

	"github.com/drudge/go-advent-of-code/pkg/aoc"
)

var cardMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

type Hand struct {
	Cards     []string
	BidAmount int
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FourOfAKind
	FiveOfAKind
	FullHouse
)

func (h Hand) ClassifyHand() (HandType, []string) {
	cardCounts := map[string]int{}

	for _, card := range h.Cards {
		cardCounts[card]++
	}

	var handType HandType
	var sortedCards []string

	for card, count := range cardCounts {
		sortedCards = append(sortedCards, strings.Repeat(card, count))
	}

	sort.Slice(sortedCards, func(i, j int) bool {
		if len(sortedCards[i]) == len(sortedCards[j]) {
			return cardMap[sortedCards[i][:1]] > cardMap[sortedCards[j][:1]]
		}
		return len(sortedCards[i]) > len(sortedCards[j])
	})

	switch len(cardCounts) {
	case 1:
		handType = FiveOfAKind
	case 2:
		if len(sortedCards[0]) == 4 {
			handType = FourOfAKind
		} else {
			handType = FullHouse
		}
	case 3:
		if len(sortedCards[0]) == 3 {
			handType = ThreeOfAKind
		} else {
			handType = TwoPair
		}
	case 4:
		handType = OnePair
	default:
		handType = HighCard
	}

	return handType, sortedCards
}

//go:embed input.txt
var inputFile embed.FS

func parseHands(lines []string) []Hand {
	hands := []Hand{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		split := strings.Split(line, " ")

		hands = append(hands, Hand{
			Cards:     strings.Split(split[0], ""),
			BidAmount: aoc.ToInt(split[1]),
		})
	}

	return hands
}

func partOne(lines []string) int {
	hands := parseHands(lines)
	sum := 0

	sort.Slice(hands, func(i, j int) bool {
		handType1, sortedCards1 := hands[i].ClassifyHand()
		handType2, sortedCards2 := hands[j].ClassifyHand()

		if handType1 != handType2 {
			return handType1 > handType2
		}

		for k := range sortedCards1 {
			if sortedCards1[k] != sortedCards2[k] {
				return cardMap[sortedCards1[k][:1]] > cardMap[sortedCards2[k][:1]]
			}
		}

		return false
	})

	for i, hand := range hands {
		rank := len(hands) - i
		log.Printf("%v: %v\n", hand, rank)
		sum += hand.BidAmount * rank
	}

	return sum
}

func partTwo(lines []string) int {
	return 0
}

func main() {
	lines := aoc.GetDayInput(inputFile)
	aoc.RunSolution(lines, partOne, partTwo)
}
