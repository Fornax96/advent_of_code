package main

import (
	"fmt"
	"strings"
	"time"

	aoc "github.com/Fornax96/advent_of_code"
)

func main() {
	defer aoc.PrintDuration(time.Now())
	lines := aoc.FileLines("input")
	fmt.Println("Answer 1:", part1(lines))
	fmt.Println("Answer 2:", part2(lines))
}

func part1(lines []string) int {
	var matches = 0

	for _, line := range lines {
		split := strings.SplitN(line, ": ", 2)
		conditions, password := split[0], split[1]

		split = strings.SplitN(conditions, " ", 2)
		limits, letter := split[0], split[1]

		split = strings.SplitN(limits, "-", 2)
		min := aoc.AtoiOrBust(split[0])
		max := aoc.AtoiOrBust(split[1])

		count := strings.Count(password, letter)
		if count >= min && count <= max {
			matches++
		}
	}
	return matches
}

func part2(lines []string) int {
	var matches = 0

	for _, line := range lines {
		split := strings.SplitN(line, ": ", 2)
		conditions, password := split[0], split[1]

		split = strings.SplitN(conditions, " ", 2)
		limits, letter := split[0], split[1][0]

		split = strings.SplitN(limits, "-", 2)
		pos1 := aoc.AtoiOrBust(split[0])
		pos2 := aoc.AtoiOrBust(split[1])

		if (password[pos1-1] == letter) != (password[pos2-1] == letter) {
			matches++
		}
	}
	return matches
}
