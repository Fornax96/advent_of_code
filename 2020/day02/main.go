package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Answer 1:")
	part1()
	fmt.Println("Answer 2:")
	part2()
}

func part1() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var matches = 0

	for scanner.Scan() {
		split := strings.SplitN(scanner.Text(), ": ", 2)
		conditions, password := split[0], split[1]

		split = strings.SplitN(conditions, " ", 2)
		limits, letter := split[0], split[1]

		split = strings.SplitN(limits, "-", 2)
		min, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		count := strings.Count(password, letter)
		if count >= min && count <= max {
			matches++
		}
	}

	fmt.Println(matches)
}

func part2() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var matches = 0

	for scanner.Scan() {
		split := strings.SplitN(scanner.Text(), ": ", 2)
		conditions, password := split[0], split[1]

		split = strings.SplitN(conditions, " ", 2)
		limits, letter := split[0], split[1][0]

		split = strings.SplitN(limits, "-", 2)
		pos1, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		pos2, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		if (password[pos1-1] == letter) != (password[pos2-1] == letter) {
			matches++
		}
	}

	fmt.Println(matches)
}
