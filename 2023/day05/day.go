package main

import (
	"embed"
	"strings"

	"github.com/drudge/go-advent-of-code/pkg/aoc"
)

//go:embed input.txt
var inputFile embed.FS

type Range struct {
	DestinationStart int
	SourceStart      int
	Length           int
}

type Mapping struct {
	Category string
	Ranges   []Range
}

func (m Mapping) MapSeed(sourceNum int) int {
	for _, m := range m.Ranges {
		if sourceNum >= m.SourceStart && sourceNum < m.SourceStart+m.Length {
			return m.DestinationStart + (sourceNum - m.SourceStart)
		}
	}

	return sourceNum
}

func mapThroughCategories(seed int, maps []Mapping) int {
	currentNumber := seed
	//log.Printf("Seed: %v\n", seed)
	for _, mapping := range maps {
		currentNumber = mapping.MapSeed(currentNumber)
		//log.Printf(" - %v: %v\n", mapping.Category, currentNumber)
	}

	return currentNumber
}

func buildSeedRange(seedStrs []string) []int {
	newSeeds := []int{}
	for i := 0; i < len(seedStrs); i += 2 {
		start := aoc.ToInt(seedStrs[i])
		length := aoc.ToInt(seedStrs[i+1])

		// log.Println("Start:", start, "Length:", length)

		for j := 0; j < length; j++ {
			newSeeds = append(newSeeds, start+j)
		}
	}

	return newSeeds
}

func parseAlmanac(lines []string, seedRange bool) ([]int, []Mapping) {
	var mappings []Mapping
	seeds := []int{}

	for _, line := range lines {
		if strings.HasPrefix(line, "seeds:") {
			seedStrs := strings.Fields(strings.TrimPrefix(line, "seeds:"))
			if seedRange {
				seeds = buildSeedRange(seedStrs)
			} else {
				for _, s := range seedStrs {
					seeds = append(seeds, aoc.ToInt(s))
				}
			}
			continue
		}

		if strings.HasSuffix(line, "map:") {
			currentMap := Mapping{
				Category: strings.TrimSuffix(strings.TrimSpace(strings.Split(line, ":")[0]), " map"),
				Ranges:   []Range{},
			}
			mappings = append(mappings, currentMap)
		} else {
			numbers := strings.Fields(line)
			if len(numbers) == 3 {
				mappings[len(mappings)-1].Ranges = append(mappings[len(mappings)-1].Ranges, Range{
					DestinationStart: aoc.ToInt(numbers[0]),
					SourceStart:      aoc.ToInt(numbers[1]),
					Length:           aoc.ToInt(numbers[2]),
				})
			}
		}
	}

	return seeds, mappings
}

func lowestFunc(b bool) aoc.InputFunc {
	return func(lines []string) int {
		min := -1
		seeds, maps := parseAlmanac(lines, b)

		// log.Printf("Seeds: %v, Maps: %v", seeds, maps)

		for _, seed := range seeds {
			location := mapThroughCategories(seed, maps)
			if min == -1 || location < min {
				min = location
			}
		}

		return min
	}
}

var partOne = lowestFunc(false)

var partTwo = lowestFunc(true)

func main() {
	lines := aoc.GetDayInput(inputFile)
	aoc.RunSolution(lines, partOne, partTwo)
}
