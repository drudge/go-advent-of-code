package main

import (
	"embed"
	"strings"

	"github.com/drudge/go-advent-of-code/pkg/aoc"
)

//go:embed *.txt
var inputFile embed.FS

type History struct {
	Numbers []int
}

type Direction int

const (
	Next Direction = iota
	Previous
)

func (h History) First() int {
	if len(h.Numbers) == 0 {
		return 0
	}
	return h.Numbers[0]
}

func (h History) Last() int {
	if len(h.Numbers) == 0 {
		return 0
	}
	return h.Numbers[len(h.Numbers)-1]
}

func (h History) IsZero() bool {
	for _, v := range h.Numbers {
		if v != 0 {
			return false
		}
	}
	return true
}

func (h History) Diff() History {
	diff := History{
		Numbers: make([]int, len(h.Numbers)-1),
	}

	for i := 0; i < len(h.Numbers)-1; i++ {
		diff.Numbers[i] = h.Numbers[i+1] - h.Numbers[i]
	}

	return diff
}

func (h History) Predict(dir Direction) int {
	if h.IsZero() {
		return 0
	}

	val := h.Diff().Predict(dir)

	if dir == Next {
		return h.Last() + val
	}

	return h.First() - val
}

func parseReport(lines []string) []History {
	fullHistory := []History{}

	for _, line := range lines {
		vals := strings.Fields(line)
		history := History{
			Numbers: make([]int, len(vals)),
		}
		for j, val := range vals {
			history.Numbers[j] = aoc.ToInt(val)
		}
		fullHistory = append(fullHistory, history)
	}

	return fullHistory
}

func predict(d Direction) aoc.InputFunc {
	return func(lines []string) int {
		predicted := 0
		history := parseReport(lines)

		for _, h := range history {
			predicted += h.Predict(d)
		}

		return predicted
	}
}

var partOne = predict(Next)

var partTwo = predict(Previous)

func main() {
	lines := aoc.GetDayInput(inputFile)
	aoc.RunSolution(lines, partOne, partTwo)
}
