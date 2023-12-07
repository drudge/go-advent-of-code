package main

import (
	"embed"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/drudge/go-advent-of-code/pkg/aoc"
)

//go:embed input.txt
var inputFile embed.FS

type CubesConfig struct {
	Red   int
	Blue  int
	Green int
}

type CubeSets struct {
	Red   []int
	Blue  []int
	Green []int
}

func (c CubesConfig) String() string {
	return fmt.Sprintf("Red: %v, Blue: %v, Green: %v", c.Red, c.Blue, c.Green)
}

type Game struct {
	ID   int
	Sets CubeSets
}

func (g Game) String() string {
	return fmt.Sprintf("Game %v: %v", g.ID, g.Sets)
}

func (g Game) GetMinPower() int {
	red := slices.Max(g.Sets.Red)
	blue := slices.Max(g.Sets.Blue)
	green := slices.Max(g.Sets.Green)
	return red * blue * green
}

func (g Game) IsPossibleWithCubes(cubes CubesConfig) bool {
	for _, red := range g.Sets.Red {
		if red > cubes.Red {
			return false
		}
	}

	for _, blue := range g.Sets.Blue {
		if blue > cubes.Blue {
			return false
		}
	}

	for _, green := range g.Sets.Green {
		if green > cubes.Green {
			return false
		}
	}

	return true
}

func parseColorSets(color string, line string) []int {
	var sets []int
	rolls := strings.Split(line, ";")
	pattern := regexp.MustCompile(`(\d+) ` + color)

	for _, roll := range rolls {
		total := 0
		matches := pattern.FindAllStringSubmatch(roll, -1)
		if len(matches) > 0 {
			for _, match := range matches {
				total += aoc.ToInt(match[1])
			}
		}
		sets = append(sets, total)
	}

	return sets
}

func parseGame(line string) Game {
	g := Game{}
	if strings.HasPrefix(line, "Game ") {
		matches := regexp.MustCompile(`Game (\d+)`).FindStringSubmatch(line)
		if len(matches) > 0 {
			g.ID = aoc.ToInt(matches[1])
		}
	}

	if g.ID > 0 {
		g.Sets.Red = parseColorSets("red", line)
		g.Sets.Blue = parseColorSets("blue", line)
		g.Sets.Green = parseColorSets("green", line)
	}

	return g
}

func partOne(lines []string) int {
	cubes := CubesConfig{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	sum := 0

	for _, line := range lines {
		game := parseGame(line)
		if game.IsPossibleWithCubes(cubes) {
			sum += game.ID
		}
	}

	return sum
}

func partTwo(lines []string) int {
	sum := 0

	for _, line := range lines {
		game := parseGame(line)
		sum += game.GetMinPower()
	}

	return sum
}

func main() {
	lines := aoc.GetDayInput(inputFile)
	aoc.RunSolution(lines, partOne, partTwo)
}
