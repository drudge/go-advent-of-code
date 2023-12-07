package main

import (
	"embed"
	"strings"

	"github.com/drudge/go-advent-of-code/pkg/aoc"
)

//go:embed input.txt
var inputFile embed.FS

type Race struct {
	Duration int
	Distance int
}

func (r Race) WaysToBeat() int {
	ways := 0

	for speed := 0; speed < r.Duration; speed++ {
		distance := speed * (r.Duration - speed)

		if distance > r.Distance {
			ways++
		}
	}

	return ways
}

func parseLogbook(lines []string, concat bool) []Race {
	races := []Race{}

	if len(lines) >= 2 {
		times := strings.Fields(lines[0])
		distances := strings.Fields(lines[1])

		if len(times) != len(distances) {
			panic("Times and distances don't match")
		}

		time := ""
		distance := ""

		for i := 1; i < len(times); i++ {
			if concat {
				time += times[i]
				distance += distances[i]
				continue
			}
			races = append(races, Race{
				Duration: aoc.ToInt(times[i]),
				Distance: aoc.ToInt(distances[i]),
			})
		}

		if concat {
			races = append(races, Race{
				Duration: aoc.ToInt(time),
				Distance: aoc.ToInt(distance),
			})
		}
	}

	return races
}

func waysToBeat(concat bool) aoc.InputFunc {
	return func(lines []string) int {
		waysToBeat := 0
		races := parseLogbook(lines, concat)

		for _, race := range races {
			ways := race.WaysToBeat()
			if waysToBeat == 0 {
				waysToBeat = ways
			} else {
				waysToBeat *= ways
			}
		}

		// log.Printf("Races: %v", races)

		return waysToBeat
	}
}

var partOne = waysToBeat(false)

var partTwo = waysToBeat(true)

func main() {
	lines := aoc.GetDayInput(inputFile)
	aoc.RunSolution(lines, partOne, partTwo)
}
