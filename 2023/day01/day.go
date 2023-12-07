package main

import (
	"embed"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/drudge/go-advent-of-code/pkg/aoc"
)

//go:embed input.txt
var inputFile embed.FS

var digitMap = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func parseNumbers(line string, includeAlpha bool) string {
	var firstDigit, lastDigit string
	var firstIndex, lastIndex = len(line), -1

	for i := 0; i < len(line); i++ {
		if includeAlpha && unicode.IsLetter(rune(line[i])) {
			for digit, word := range digitMap {
				if strings.HasPrefix(line[i:], word) {
					index := i + len(word) - 1
					digitStr := strconv.Itoa(digit)

					if firstIndex > i {
						firstDigit = digitStr
						firstIndex = i
					}
					if lastIndex < index {
						lastDigit = digitStr
						lastIndex = index
					}
					i = index
					break
				}
			}
		} else if unicode.IsDigit(rune(line[i])) {
			digitStr := string(line[i])

			if firstIndex > i {
				firstDigit = digitStr
				firstIndex = i
			}
			if lastIndex < i {
				lastDigit = digitStr
				lastIndex = i
			}
		}
	}

	if firstIndex == len(line) {
		return ""
	}
	return firstDigit + lastDigit
}

func partOne(input []string) int {
	sum := 0

	for _, line := range input {
		sum += aoc.ToInt(parseNumbers(line, false))
	}
	return sum
}

func partTwo(input []string) int {
	sum := 0

	for i, line := range input {
		parsed := parseNumbers(line, true)
		log.Printf("[%v] line: %v, parsed: %v\n", i, line, parsed)
		if parsed != "" {
			sum += aoc.ToInt(parsed)
		}
	}
	return sum
}

func main() {
	lines := aoc.GetDayInput(inputFile)
	aoc.RunSolution(lines, partOne, partTwo)
}
