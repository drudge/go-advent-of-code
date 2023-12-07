package aoc

import (
	"bufio"
	"embed"
	"fmt"
	"io/fs"
	"strconv"
	"time"
)

type InputFunc func([]string) int

func GetDayInputFile(embedInput embed.FS, filename string) []string {
	file, err := embedInput.Open(filename)
	if err != nil {
		panic(err)
	}

	defer closeFile(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func GetDayInput(embedInput embed.FS) []string {
	return GetDayInputFile(embedInput, "input.txt")
}

func closeFile(f fs.File) {
	err := f.Close()
	if err != nil {
		panic(err)
	}
}

func Seq(x, y int) []int {
	l := make([]int, y-x+1)
	for i := x; i <= y; i++ {
		l[i-x] = i
	}
	return l
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func RunSolution(input []string, partOne InputFunc, partTwo InputFunc) {
	if partOne != nil {
		RunPart("Part One", input, partOne)
	}
	if partTwo != nil {
		RunPart("Part Two", input, partTwo)
	}
}

func RunTime(startTime time.Time) string {
	var timeStr string

	duration := time.Since(startTime)

	switch {
	case duration < time.Millisecond:
		timeStr = fmt.Sprintf("%.1fµs", float64(duration.Microseconds())+float64(duration.Nanoseconds()%1000)/1000)
	case duration < time.Second:
		timeStr = fmt.Sprintf("%.2fms", float64(duration.Milliseconds()))
	case duration < time.Minute:
		timeStr = fmt.Sprintf("%.2fs", duration.Seconds())
	}

	return timeStr
}

func RunPart(name string, input []string, part InputFunc) {
	startTime := time.Now()
	sum := part(input)
	println(fmt.Sprintf("%s: %v (%s)", name, sum, RunTime(startTime)))
}

func AssertEqual(expected, actual int) {
	if expected != actual {
		panic(fmt.Sprintf("Expected %v, got %v", expected, actual))
	}
}
