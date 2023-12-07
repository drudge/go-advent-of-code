package main

import (
	"embed"

	"github.com/drudge/go-advent-of-code/pkg/aoc"
)

//go:embed input.txt
var inputFile embed.FS

func partOne(lines []string) int {
	return 0
}

func partTwo(lines []string) int {
	return 0
}

func main() {
	lines := aoc.GetDayInput(inputFile)
	aoc.RunSolution(lines, partOne, partTwo)
}
