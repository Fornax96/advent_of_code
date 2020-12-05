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
		var row, col = 0, 0
		for _, c := range line {
			switch c {
			case 'F': // Front
				row = row << 1
			case 'B': // Back
				row = (row << 1) + 1
			case 'L': // Left
				col = col << 1
			case 'R': // Right
				col = (col << 1) + 1
			}
		}

		if (row*8)+col > idTotal {
			idTotal = (row * 8) + col
		}
	}

	return idTotal
}

func part2(lines []string) (myID int) {
	var ids = make(map[int]bool)

	for _, line := range lines {
		var row, col = 0, 0
		for _, c := range line {
			switch c {
			case 'F': // Front
				row = row << 1
			case 'B': // Back
				row = (row << 1) + 1
			case 'L': // Left
				col = col << 1
			case 'R': // Right
				col = (col << 1) + 1
			}
		}
		ids[(row*8)+col] = true
	}

	// Find the missing id
	var h1, h2, h3 = false, false, false
	var lastID = 0

	for row := 1; row < 127; row++ {
		for col := 0; col < 8; col++ {
			h3 = h2
			h2 = h1
			h1 = ids[(row*8)+col]

			if h3 && !h2 && h1 {
				return lastID
			}
			lastID = (row * 8) + col
		}
	}
	return 0
}
