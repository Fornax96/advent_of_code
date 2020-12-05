package main

import (
	"fmt"
	"time"

	aoc "github.com/Fornax96/advent_of_code"
)

func main() {
	defer aoc.PrintDuration(time.Now())
	lines := aoc.FileLines("input")
	fmt.Println("Answer 1:", part1(lines))
	fmt.Println("Answer 2:", part2(lines))
}

func part1(lines []string) (idTotal int) {
	for _, line := range lines {
		var id = 0
		for _, c := range line {
			switch c {
			case 'F', 'L': // Shift 0
				id = id << 1
			case 'B', 'R': // Shift 1
				id = (id << 1) + 1
			}
		}

		if id > idTotal {
			idTotal = id
		}
	}
	return idTotal
}

func part2(lines []string) int {
	var ids = make(map[int]bool)

	for _, line := range lines {
		var id = 0
		for _, c := range line {
			switch c {
			case 'F', 'L': // Shift 0
				id = id << 1
			case 'B', 'R': // Shift 1
				id = (id << 1) + 1
			}
		}
		ids[id] = true
	}

	// Find the missing id
	var h1, h2, h3, lastID = false, false, false, 0
	for row := 0; row < 128; row++ {
		for col := 0; col < 8; col++ {
			h3 = h2
			h2 = h1
			h1 = ids[(row<<3)+col]

			if h3 && !h2 && h1 {
				return lastID
			}
			lastID = (row << 3) + col
		}
	}
	return 0
}
