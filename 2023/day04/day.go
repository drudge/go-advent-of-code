package main

import (
	"embed"
	"fmt"
	"slices"
	"strings"

	"github.com/drudge/go-advent-of-code/pkg/aoc"
)

//go:embed input.txt
var inputFile embed.FS

type card struct {
	CardNumber     string
	WinningNumbers []string
	MyNumbers      []string
	Copies         int
}

func (c card) String() string {
	return fmt.Sprintf("Card Number: %v\nCopies: %v\nScore: %v\nNumbers: %v\nWinning Numbers: %v\n", c.CardNumber, c.Copies, c.Score(), c.WinningNumbers, c.MyNumbers)
}

func (c card) Score() int {
	return c.Copies + 1
}

func (c card) NumWinners() int {
	copies := 0
	for _, num := range c.MyNumbers {
		if slices.Contains(c.WinningNumbers, num) {
			copies += 1
		}
	}
	return copies
}

type cardSet []card

func (cs cardSet) Len() int {
	return len(cs)
}

func (cs cardSet) Score() int {
	score := 0
	for _, card := range cs {
		score += card.Score()
	}
	return score
}

func parseCardSet(file embed.FS) cardSet {
	var output cardSet

	for _, line := range aoc.GetDayInput(file) {
		output = append(output, parseCard(line))
	}

	return output
}

func parseCard(line string) card {
	colonPos := strings.Index(line, ":")
	pipePos := strings.Index(line, " |")

	return card{
		CardNumber:     line[:colonPos],
		WinningNumbers: strings.Fields(line[colonPos+2 : pipePos]),
		MyNumbers:      strings.Fields(line[pipePos+3:]),
		Copies:         0,
	}
}

func main() {
	cards := parseCardSet(inputFile)
	for i, card := range cards {
		winners := card.NumWinners()
		if winners > 0 {
			card.Copies++
			for j := 1; j <= winners; j++ {
				cards[i+j].Copies += card.Copies * 1
			}
		}
	}

	total := 0
	for _, card := range cards {
		total += card.Copies + 1
	}

	fmt.Printf("Total: %v\n", total)
}
