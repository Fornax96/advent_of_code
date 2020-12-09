package main

import (
	"fmt"
	"time"

	aoc "github.com/Fornax96/advent_of_code"
)

func main() {
	defer aoc.PrintDuration(time.Now())
	lines := aoc.FileInts("input")
	ans := part1(lines)
	fmt.Println("Answer 1:", ans)
	fmt.Println("Answer 2:", part2(lines, ans))
}

func part1(ints []int) int {
	for k1 := 25; k1 < len(ints); k1++ {
		var found = false
		for k2 := k1 - 25; k2 < k1 && !found; k2++ {
			for k3 := k1 - 25; k3 < k1 && !found; k3++ {
				if ints[k2]+ints[k3] == ints[k1] {
					found = true
				}
			}
		}
		if !found {
			return ints[k1]
		}
	}
	panic("number not found")
}

func part2(ints []int, ans int) int {
	for length := 2; length < len(ints)-length; length++ { // Specify a range
		for start := 0; start+length < len(ints); start++ { // Loop over the numbers in the range
			var sum, smallest, largest = 0, ints[start], ints[start]
			for k := start; k < start+length; k++ {
				sum += ints[k]
				if ints[k] < smallest {
					smallest = ints[k]
				}
				if ints[k] > largest {
					largest = ints[k]
				}
			}
			if sum == ans {
				return smallest + largest
			}
		}
	}
	panic("number not found")
}
