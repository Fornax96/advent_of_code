package main

import (
	"fmt"
	"time"

	aoc "github.com/Fornax96/advent_of_code"
)

func main() {
	defer aoc.PrintDuration(time.Now())
	lines := append(aoc.FileLines("input"), "")
	fmt.Println("Answer 1:", part1(lines))
	fmt.Println("Answer 2:", part2(lines))
}

func part1(lines []string) (votes int) {
	var group = make(map[rune]bool)
	var groups []map[rune]bool

	for _, line := range lines {
		if line == "" {
			groups = append(groups, group)
			group = make(map[rune]bool)
		}

		for _, vote := range line {
			if _, ok := group[vote]; !ok {
				group[vote] = true
				votes++ // If this vote is unque within the group we count it
			}
		}
	}
	return votes
}

func part2(lines []string) (votes int) {
	var group = make(map[rune]int)
	var groups []map[rune]int
	var groupSize = 0

	for _, line := range lines {
		if line == "" {
			for _, v := range group {
				if v == groupSize { // If all people in the group voted for this
					votes++
				}
			}
			groups = append(groups, group)
			group = make(map[rune]int)
			groupSize = 0
			continue
		}

		groupSize++
		for _, vote := range line {
			group[vote]++
		}
	}
	return votes
}
