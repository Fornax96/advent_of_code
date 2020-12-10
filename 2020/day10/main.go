package main

import (
	"fmt"
	"sort"
	"time"

	aoc "github.com/Fornax96/advent_of_code"
)

func main() {
	defer aoc.PrintDuration(time.Now())
	lines := append(aoc.FileInts("input"), 0) // Start with 0
	sort.Ints(lines)
	lines = append(lines, lines[len(lines)-1]+3) // End with the highest value + 3

	fmt.Println("Answer 1:", part1(lines))
	fmt.Println("Answer 2:", part2(lines))
}

func part1(ints []int) int {
	var lastJolts, j1, j3 = 0, 0, 0
	for _, v := range ints {
		if v-lastJolts == 1 {
			j1++
		} else if v-lastJolts == 3 {
			j3++
		}
		lastJolts = v
	}
	return j1 * j3
}

func part2(ints []int) int {
	var lastJolts, lastgroupend = 0, 0
	var groups [][]int
	for k, v := range ints {
		if v-lastJolts == 3 {
			groups = append(groups, ints[lastgroupend:k])
			lastgroupend = k
		}
		lastJolts = v
	}

	var variationsTot = 1
	var groupCp []int
	for _, v := range groups {
		if len(v) < 3 {
			continue
		}
		groupCp = make([]int, len(v))
		variations := 0
		for mask := 0; mask < 1<<(len(v)-2); mask++ {
			copy(groupCp, v)

			// Mask out the values we don't want
			for idx := 0; idx < len(v)-2; idx++ {
				if mask&(1<<idx) == (1 << idx) {
					groupCp[idx+1] = -1
				}
			}
			if checkGroup(groupCp) {
				variations++
			}
		}

		variationsTot *= variations
	}
	return variationsTot
}

func checkGroup(ints []int) bool {
	var last = ints[0]
	for _, v := range ints {
		if v == -1 {
			continue
		}
		if v-last > 3 {
			return false
		}
		last = v
	}
	return true
}
