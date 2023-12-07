package main

import (
	"embed"
	"log"
	"regexp"

	"github.com/drudge/go-advent-of-code/pkg/aoc"
)

//go:embed input.txt
var inputFile embed.FS

type Node struct {
	Key string
	L   string
	R   string
}

type Map struct {
	Directions []rune
	Nodes      map[string]Node
}

func parseMap(lines []string) Map {
	newMap := Map{}

	if len(lines) > 3 {
		newMap.Nodes = make(map[string]Node)
		newMap.Directions = []rune(lines[0])

		for i := 2; i < len(lines); i++ {
			node := Node{}
			re := regexp.MustCompile(`([0-9A-Z]{3}) = \(([0-9A-Z]{3}), ([0-9A-Z]{3})\)`)
			matches := re.FindAllStringSubmatch(lines[i], -1)

			if len(matches) > 0 {
				node.Key = matches[0][1]
				node.L = matches[0][2]
				node.R = matches[0][3]
			}
			newMap.Nodes[node.Key] = node
		}
	}

	return newMap
}

func partOne(lines []string) int {
	startNode := "AAA"
	endNode := "ZZZ"
	steps := 0
	parsed := parseMap(lines)
	origin, exists := parsed.Nodes[startNode]

	// log.Printf("Parsed: %+v", parsed)

	if !exists {
		panic("Start node not found")
	}

	directionIndex := 0
	for origin.Key != endNode {
		if directionIndex >= len(parsed.Directions) {
			directionIndex = 0
		}
		direction := parsed.Directions[directionIndex]
		if direction == 'R' {
			origin = parsed.Nodes[origin.R]
			// log.Printf("Moved right to %+v", origin)
		} else {
			origin = parsed.Nodes[origin.L]
			// log.Printf("Moved left to %+v", origin)
		}
		steps++
		directionIndex++
	}

	return steps
}

func partTwo(lines []string) int {
	parsed := parseMap(lines)

	// Initialize startingNodes with all nodes ending in 'A'
	startingNodes := make(map[string]int)
	for key := range parsed.Nodes {
		if len(key) == 3 && key[2] == 'A' {
			startingNodes[key] = 0 // Initialize direction index to 0
		}
	}

	steps := 0
	finished := false

	for !finished {
		finished = true
		for key, directionIndex := range startingNodes {
			if key[2] != 'Z' {
				finished = false
				direction := parsed.Directions[directionIndex%len(parsed.Directions)]
				node := parsed.Nodes[key]
				if direction == 'R' {
					key = node.R
				} else {
					key = node.L
				}
				log.Printf("Moved %c to %s", direction, key)
				startingNodes[key] = (directionIndex + 1) % len(parsed.Directions)
			}
		}
		if !finished {
			steps++
		}
	}

	return steps
}

func main() {
	lines := aoc.GetDayInput(inputFile)
	aoc.RunSolution(lines, nil, partTwo)
}

/*

Step 0: You are at 11A and 22A.
Step 1: You choose all of the left paths, leading you to 11B and 22B.
Step 2: You choose all of the right paths, leading you to 11Z and 22C.
Step 3: You choose all of the left paths, leading you to 11B and 22Z.
Step 4: You choose all of the right paths, leading you to 11Z and 22B.
Step 5: You choose all of the left paths, leading you to 11B and 22C.
Step 6: You choose all of the right paths, leading you to 11Z and 22Z.
So, in this example, you end up entirely on nodes that end in Z after 6 steps.

*/
