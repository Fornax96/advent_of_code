package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/Fornax96/advent_of_code"
)

func main() {
	part1()
	part2()
}

func part1() {
	const numbers = "1234567890"
	var total = 0

	for _, line := range aoc.FileLines("input") {
		first := strings.IndexAny(line, numbers)
		last := strings.LastIndexAny(line, numbers)

		num, _ := strconv.Atoi(string(line[first]) + string(line[last]))
		total += num
	}

	fmt.Println(total)
}

func part2() {
	const numbers = "123456789"
	var total = 0

	var wordnums = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, line := range aoc.FileLines("input") {
		firstIdx := 99
		firstNum := 0
		lastIdx := -1
		lastNum := 0

		for wordi, word := range wordnums {
			i := strings.Index(line, word)
			if i != -1 && i < firstIdx {
				firstIdx = i
				firstNum = wordi + 1
			}

			i = strings.LastIndex(line, word)
			if i != -1 && i > lastIdx {
				lastIdx = i
				lastNum = wordi + 1
			}
		}

		firstdigit := strings.IndexAny(line, numbers)
		if firstdigit != -1 && firstdigit < firstIdx {
			firstNum, _ = strconv.Atoi(string(line[firstdigit]))
		}
		lastdigit := strings.LastIndexAny(line, numbers)
		if lastdigit != -1 && lastdigit > lastIdx {
			lastNum, _ = strconv.Atoi(string(line[lastdigit]))
		}

		num, _ := strconv.Atoi(strconv.Itoa(firstNum) + strconv.Itoa(lastNum))
		total += num
	}

	fmt.Println(total)
}
