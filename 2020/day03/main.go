package main

import (
	"fmt"

	aoc "github.com/Fornax96/advent_of_code"
)

func main() {
	f := aoc.FileScanner("input")
	defer f.Close()

	var treeMap [][]bool
	for f.Scan() {
		line := f.Text()
		row := make([]bool, len(line))
		for i, v := range line {
			if v == '.' {
				row[i] = false
			} else if v == '#' {
				row[i] = true
			}
		}

		treeMap = append(treeMap, row)
	}

	fmt.Println("Answer 1:", part1(treeMap))
	fmt.Println("Answer 2:", part2(treeMap))
}

func part1(lines [][]bool) (trees int) {
	var x, y = 3, 1

	for y < len(lines) {
		if line := lines[y]; line[x%len(line)] {
			trees++
		}

		x = (x + 3) % len(lines[y])
		y++
	}
	return trees
}

func part2(lines [][]bool) (trees int) {
	type Slope struct {
		xMod int
		yMod int
	}

	trees = 1
	for _, slope := range []Slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		var x, y, slopeTrees = slope.xMod, slope.yMod, 0

		for y < len(lines) {
			if line := lines[y]; line[x%len(line)] {
				slopeTrees++
			}

			x = (x + slope.xMod) % len(lines[y])
			y += slope.yMod
		}

		trees *= slopeTrees
	}
	return trees
}
