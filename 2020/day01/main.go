package main

import (
	"fmt"

	aoc "github.com/Fornax96/advent_of_code"
)

func main() {
	scanner := aoc.FileScanner("input")
	defer scanner.Close()

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, aoc.AtoiOrBust(scanner.Text()))
	}

	fmt.Println("Answer 1:", part1(numbers))
	fmt.Println("Answer 2:", part2(numbers))
}

func part1(numbers []int) int {
	for _, v1 := range numbers {
		for _, v2 := range numbers {
			if v1+v2 == 2020 {
				return v1 * v2
			}
		}
	}
	panic("number not found")
}

func part2(numbers []int) int {
	for _, v1 := range numbers {
		for _, v2 := range numbers {
			for _, v3 := range numbers {
				if v1+v2+v3 == 2020 {
					return v1 * v2 * v3
				}
			}
		}
	}
	panic("number not found")
}
