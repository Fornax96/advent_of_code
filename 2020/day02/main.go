package main

import (
	"fmt"
	"strings"

	aoc "github.com/Fornax96/advent_of_code"
)

func main() {
	fmt.Println("Answer 1:", part1())
	fmt.Println("Answer 2:", part2())
}

func part1() int {
	scanner := aoc.FileScanner("input")
	defer scanner.Close()

	var matches = 0

	for scanner.Scan() {
		split := strings.SplitN(scanner.Text(), ": ", 2)
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

func part2() int {
	scanner := aoc.FileScanner("input")
	defer scanner.Close()

	var matches = 0

	for scanner.Scan() {
		split := strings.SplitN(scanner.Text(), ": ", 2)
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
