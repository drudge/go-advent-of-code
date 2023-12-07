# Advent of Code - In Go

## Usage

copy the `template` folder in `2023/` and name it according to the puzzle you're solving, ie: `2023/day08/`.

copy the example input into `example.txt`, your input into `input.txt`, and the puzzle text into `puzzle.md`.

run it:

```bash
go run ./2023/dayXX
```

## aoc Package Documentation

### Overview

The aoc package is designed to assist in solving Advent of Code (AoC) puzzles. It offers utilities for managing input data, timing solution execution, and includes a client for downloading puzzle inputs.

### Package Components

#### `aoc`

- **GetDayInputFile(embedInput embed.FS, filename string) []string:** Reads a file from an embedded filesystem and returns its lines.
- **GetDayInput(embedInput embed.FS) []string:** A convenience function for reading "input.txt" from an embedded filesystem.
- **Seq(x, y int) []int:** Generates a sequence of integers from x to y.
- **ToInt(s string) int:** Converts a string to an integer, with error handling.
- **RunSolution(input []string, partOne aoc.InputFunc, partTwo aoc.InputFunc):** Executes solution functions for AoC puzzle parts, if provided.
- **RunTime(startTime time.Time) string:** Formats execution time since a provided start time.
- **RunPart(name string, input []string, part aoc.InputFunc):** Executes and times a single part of an AoC puzzle, printing results.

#### `aoc.Client`

- **NewClient(token string) aoc.Client:** Constructor function to create a new Client with a given AoC session token.
- **Client.DownloadInput(year, day int) (io.ReadCloser, error):** Method for downloading the input file for a specific year and day of AoC.
- **Client.DownloadAndSaveInput(year, day int, targetFile string) error:** Downloads the specified AoC puzzle input and saves it to a target file.

## License

__The MIT License__

Copyright (c) 2023 Nicholas Penree <nick@penree.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.