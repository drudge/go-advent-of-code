package main

import (
	"embed"

	"github.com/drudge/go-advent-of-code/pkg/aoc"
)

//go:embed input.txt
var inputFile embed.FS

type CellType int

const (
	None CellType = iota
	Symbol
	Number
)

type Cell struct {
	Type   CellType
	Value  int
	Coords []Coord
}

type Coord struct {
	X, Y int
}

func parseGrid(lines []string) [][]Cell {
	grid := make([][]Cell, len(lines))

	for y, line := range lines {
		grid[y] = make([]Cell, len(line))

		for x, ch := range line {
			if ch == '.' {
				grid[y][x] = Cell{Type: None}
			} else if isSymbol(ch) {
				grid[y][x] = Cell{Type: Symbol, Value: int(ch)}
			} else if isDigit(ch) {
				if x > 0 && isDigit(rune(line[x-1])) {
					continue
				}

				number, coords := parseNumber(lines, x, y)

				for _, c := range coords {
					grid[c.Y][c.X] = Cell{
						Type:   Number,
						Value:  number,
						Coords: coords,
					}
				}
			}
		}
	}

	return grid
}

func parseNumber(lines []string, x, y int) (int, []Coord) {
	numberStr := ""
	coords := []Coord{}

	for checkX := x; checkX < len(lines[y]) && isDigit(rune(lines[y][checkX])); checkX++ {
		numberStr += string(lines[y][checkX])
		coords = append(coords, Coord{X: checkX, Y: y})
	}

	number := aoc.ToInt(numberStr)

	return number, coords
}

func isAdjacentToSymbol(grid [][]Cell, coords []Coord, counted map[Coord]bool) bool {
	for _, coord := range coords {
		if counted[coord] {
			continue
		}

		if checkAdjacent(grid, coord.X, coord.Y) {
			return true
		}
	}

	return false
}

var surroundingArea = []Coord{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	{1, 1}, {-1, -1}, {1, -1}, {-1, 1},
}

func getAdjacentParts(grid [][]Cell, x, y int) []int {
	uniqueParts := []int{}
	parts := make(map[int]bool)

	for _, d := range surroundingArea {
		checkX, checkY := x+d.X, y+d.Y

		if checkX >= 0 && checkY >= 0 && checkY < len(grid) && checkX < len(grid[checkY]) && grid[checkY][checkX].Type == Number {
			parts[grid[checkY][checkX].Value] = true
		}
	}

	for part := range parts {
		uniqueParts = append(uniqueParts, part)
	}

	return uniqueParts
}

func checkAdjacent(grid [][]Cell, x, y int) bool {
	for _, coord := range surroundingArea {
		checkX, checkY := x+coord.X, y+coord.Y

		if checkX >= 0 && checkY >= 0 && checkY < len(grid) && checkX < len(grid[checkY]) {
			if grid[checkY][checkX].Type == Symbol {
				return true
			}
		}
	}

	return false
}

func isSymbol(ch rune) bool {
	return !isDigit(ch) && ch != '.'
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func partOne(lines []string) int {
	sum := 0
	grid := parseGrid(lines)
	counted := make(map[Coord]bool)

	for _, row := range grid {
		for _, cell := range row {
			if cell.Type == Number && isAdjacentToSymbol(grid, cell.Coords, counted) {
				sum += cell.Value
				for _, c := range cell.Coords {
					counted[c] = true
				}
			}
		}
	}
	return sum
}

func partTwo(lines []string) int {
	sum := 0
	grid := parseGrid(lines)

	for y, row := range grid {
		for x, cell := range row {
			if cell.Type == Symbol && rune(grid[y][x].Value) == '*' {
				adjacentParts := getAdjacentParts(grid, x, y)
				if len(adjacentParts) == 2 {
					sum += adjacentParts[0] * adjacentParts[1]
				}
			}
		}
	}

	return sum
}

func main() {
	lines := aoc.GetDayInput(inputFile)
	aoc.RunSolution(lines, partOne, partTwo)
}
