package main

import (
	"fmt"
	"time"

	aoc "github.com/Fornax96/advent_of_code"
)

func main() {
	defer aoc.PrintDuration(time.Now())

	var lines = aoc.FileLines("input")
	fmt.Println("Answer 1:", part1(lines))
	fmt.Println("Answer 2:", part2(lines))
}

func part1(lines []string) int {
	var op rune
	var param int
	var direction = 1 // 0 north, 1 east, 2 south, 3 west
	var north, east = 0, 0
	for _, line := range lines {
		op = rune(line[0])
		param = aoc.AtoiOrBust(line[1:])
		switch {
		case op == 'N', op == 'F' && direction == 0:
			north += param
		case op == 'E', op == 'F' && direction == 1:
			east += param
		case op == 'S', op == 'F' && direction == 2:
			north -= param
		case op == 'W', op == 'F' && direction == 3:
			east -= param
		case op == 'L':
			direction -= (param / 90)
			if direction < 0 {
				direction += 4
			}
		case op == 'R':
			direction += (param / 90)
			if direction > 3 {
				direction -= 4
			}
		}
	}
	return abs(north) + abs(east)
}

func part2(lines []string) int {
	var op rune
	var param int
	var north, east, northWP, eastWP = 0, 0, 1, 10
	for _, line := range lines {
		op = rune(line[0])
		param = aoc.AtoiOrBust(line[1:])
		switch op {
		case 'N':
			northWP += param
		case 'E':
			eastWP += param
		case 'S':
			northWP -= param
		case 'W':
			eastWP -= param
		case 'L', 'R':
			rot := param / 90
			if op == 'L' {
				rot = -rot + 4
			}

			n, e := northWP, eastWP
			switch rot {
			case 1:
				northWP = -e
				eastWP = n
			case 2:
				northWP = -n
				eastWP = -e
			case 3:
				northWP = e
				eastWP = -n
			}
		case 'F':
			north += northWP * param
			east += eastWP * param
		}
	}
	return abs(north) + abs(east)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
